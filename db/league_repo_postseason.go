package db

import (
	"fdsim/data"
	"fdsim/generators"
	"fdsim/libs"
	"fdsim/models"
	"fmt"
	"time"

	"golang.org/x/exp/maps"
)

type TeamLosingPlayersMap map[string]map[models.Role]int

func (lr *LeagueRepo) createNewLeague(game *models.Game) *models.League {
	oldLeague := lr.ById(game.LeagueId)

	lr.g.Raw("update team_dtos set league_id = null where 1=1")
	newLeague := models.NewLeague(oldLeague.Teams(), game.Date)
	leagueName := data.GetLeagueName(oldLeague.Country)
	name := fmt.Sprintf("%s %d/%d", leagueName, game.Date.Year(), game.Date.Year()+1)
	newLeague.UpdateLocales(name, oldLeague.Country)
	game.LeagueId = newLeague.Id
	newLeagueDto := DtoFromLeagueEmpty(newLeague)
	lr.g.Create(&newLeagueDto)

	lr.g.Raw("update team_dtos set league_id = ? where 1=1", newLeague.Id)
	newLeagueDto = DtoFromLeague(newLeague)

	lr.g.Save(&newLeagueDto)
	gameDto := DtoFromGame(game)
	lr.g.Save(&gameDto)

	return lr.ByIdFull(newLeague.Id)
}
func (lr *LeagueRepo) generateYoungReplacements(game *models.Game, tm TeamLosingPlayersMap, rng *libs.Rng) {
	pg := generators.NewPeopleGenSeeded(rng)
	pToAdd := []PlayerDto{}
	for teamId, rc := range tm {
		for role, count := range rc {
			// TODO: maybe do this as 0, could avoid replacing it at all in case
			effectiveCount := rng.UInt(1, count)
			ps := pg.YoungPlayersWithRole(effectiveCount, role)
			for _, p := range ps {
				pdto := DtoFromPlayer(p)
				pdto.TeamId = &teamId
				pToAdd = append(pToAdd, pdto)
			}
		}
	}

	lr.g.Create(&pToAdd)
}

// Check contracts and if 0 put them on the free market
func (lr *LeagueRepo) playersEndOfSeason(game *models.Game, tm TeamLosingPlayersMap, rng *libs.Rng) {
	fdTeamId := ""
	if game.IsEmployed() {
		fdTeamId = game.Team.Id
	}

	var expiredContractsP []PlayerDto
	lr.g.Exec("update player_dtos set y_contract = y_contract - 1 where team_id is not null")

	lr.g.Model(&PlayerDto{}).Where("y_contract < ?", 1).Find(&expiredContractsP)
	for i, p := range expiredContractsP {
		// 50% chance of renewing contract for 1 year
		// if not in FDTeam
		if p.TeamId != &fdTeamId && rng.ChanceI(50) {
			expiredContractsP[i].YContract = 1
		} else {
			expiredContractsP[i].TeamId = nil
			countPlayerLoss(p, tm)
		}
	}
	if len(expiredContractsP) > 0 {
		lr.g.Save(&expiredContractsP)
	}
}

func (lr *LeagueRepo) retirePlayers(indexedP map[string]PHistoryDto, leagueId, leagueName string, gameDate time.Time, tm TeamLosingPlayersMap, rng *libs.Rng) {
	// Age players/Coach
	trx := lr.g.Exec("update player_dtos set age = age + 1 where 1=1; update coach_dtos set age = age + 1 where 1=1;")
	if trx.RowsAffected < 1 {
		panic("something wrong")
	}
	var playersCount int64
	lr.g.Model(&PlayerDto{}).Count(&playersCount)
	var playersToRetire []PlayerDto
	lr.g.Raw("select * from player_dtos where age > 35 order by RANDOM() LIMIT ?", int(playersCount)/rng.UInt(2, 10)).Preload(teamRel).Find(&playersToRetire)
	if len(playersToRetire) == 0 {
		return
	}
	//TODO: maybe ad add a way to replace players
	pIds := make([]string, len(playersToRetire))
	retiring := make([]RetiredPlayerDto, len(playersToRetire))
	for i, p := range playersToRetire {
		retiring[i] = NewRetiredPlayerFromDto(p, indexedP, gameDate.Year(), leagueId, leagueName)
		pIds[i] = p.Id

		countPlayerLoss(p, tm)
	}
	lr.g.Create(&retiring)

	lr.g.Delete(&PHistoryDto{}, pIds)
	lr.g.Delete(&PlayerDto{}, pIds)
}

func countPlayerLoss(p PlayerDto, tm TeamLosingPlayersMap) {
	if p.TeamId == nil {
		return
	}

	if pc, ok := tm[*p.TeamId]; ok {
		pc[p.Role]++
	} else {
		tm[*p.TeamId] = models.NewEmptyRoleCounter()
		tm[*p.TeamId][p.Role]++
	}
}

func (lr *LeagueRepo) convertStatsToHistory(leagueName string, gameDate time.Time, leagueId string) (map[string]PHistoryDto, map[string]THistoryDto) {
	var playersStats []StatRowDto
	lr.g.Model(&StatRowDto{}).Preload(teamRel).Find(&playersStats)

	var pHistoryRows []PHistoryDto
	lr.g.Model(&PHistoryDto{}).Find(&pHistoryRows)
	indexedPHRows := indexPHistoryRows(pHistoryRows)

	for _, s := range playersStats {
		if existingRow, ok := indexedPHRows[s.PlayerId]; ok {
			//TODO: maybe move to pointer

			existingRow.Update(s, leagueName, gameDate)
			indexedPHRows[s.PlayerId] = existingRow

		} else {
			indexedPHRows[s.PlayerId] = DtoFromPHistoryRow(
				models.NewPHistoryRow(
					s.StatRow(),
					leagueName,
					gameDate,
				),
			)
		}
	}
	phrows := maps.Values(indexedPHRows)

	var teamsStats []TableRowIndexDto
	lr.g.Raw(`SELECT team_id, played, wins, draws, losses, points, goal_scored, goal_conceded,
			   ROW_NUMBER() OVER (ORDER BY points DESC, goal_scored DESC, goal_conceded ASC) AS position
		FROM table_row_dtos`).Find(&teamsStats)

	var tHistoryRows []THistoryDto
	lr.g.Model(&THistoryDto{}).Find(&tHistoryRows)
	indexedTHRows := indexTHistoryRows(tHistoryRows)
	for _, s := range teamsStats {
		if existingRow, ok := indexedTHRows[s.TeamId]; ok {
			//TODO: maybe move to pointer
			existingRow.Update(s, leagueId, leagueName, gameDate)
			indexedTHRows[s.TeamId] = existingRow

		} else {
			indexedTHRows[s.TeamId] = DtoFromTHistoryRow(
				models.NewTHistoryRow(
					s.TPHRow(),
					leagueId,
					leagueName,
					gameDate,
				),
			)
		}
	}
	throws := maps.Values(indexedTHRows)
	lr.cleanStats()
	lr.g.Save(phrows)
	lr.g.Save(throws)

	return indexedPHRows, indexedTHRows
}

func (lr *LeagueRepo) updateFDInfo(game *models.Game, leagueName string) {
	game.Age++

	if !game.IsEmployed() {
		return
	}

	var stat FDStatRowDto
	lr.g.Model(&FDStatRowDto{}).Order("hired_date desc").First(&stat)
	h := NewFDHistoryDto(stat)
	h.UpdateEndOfSeason(game.LeagueId, leagueName, game.Wage)
	lr.g.Save(&h)
	// lr.g.Delete(&stat).Where("1=1")

	game.YContract -= 1
	if game.YContract == 0 {
		game.UnsetTeamContract()
	} else {
		newStat := models.NewFDStatRow(game.Date, stat.TeamId, stat.TeamName)
		lr.g.Model(&FDStatRowDto{}).Create(&newStat)
	}
}

func (lr *LeagueRepo) cleanStats() {
	lr.g.Where("1 = 1").Delete(&TableRowDto{})
	lr.g.Where("1 = 1").Delete(&ResultDto{})
	lr.g.Where("1 = 1").Delete(&MatchDto{})
	lr.g.Where("1 = 1").Delete(&RoundDto{})
	lr.g.Where("1 = 1").Delete(&StatRowDto{})

}

func indexPHistoryRows(historyRows []PHistoryDto) map[string]PHistoryDto {
	result := map[string]PHistoryDto{}
	for _, v := range historyRows {
		result[v.PlayerId] = v
	}

	return result
}

func indexTHistoryRows(historyRows []THistoryDto) map[string]THistoryDto {
	result := map[string]THistoryDto{}
	for _, v := range historyRows {
		result[v.TeamId] = v
	}

	return result
}
