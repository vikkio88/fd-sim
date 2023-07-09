package db

import (
	"fdsim/data"
	"fdsim/libs"
	"fdsim/models"
	"fmt"
	"time"

	"golang.org/x/exp/maps"
)

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

func (lr *LeagueRepo) playersEndOfSeason(gameDate time.Time) {

	// Check contracts and if 0 put them on the free market
	// Check contracts and if 0 put them on the free market

}

func (lr *LeagueRepo) retirePlayers(indexedP map[string]PHistoryDto, leagueId, leagueName string, gameDate time.Time) {
	// Age players/Coach
	lr.g.Raw("update player_dtos set age = age + 1 where 1=1; update coach_dtos set age = age+1 where 1=1;")
	//TODO: dont forget to udpate FD Age

	// TODO: maybe inject this
	rng := libs.NewRng(time.Now().Unix())
	var playersCount int64
	lr.g.Model(&PlayerDto{}).Count(&playersCount)
	var playersToRetire []PlayerDto
	lr.g.Raw("select * from player_dtos where age > 35 order by RANDOM() LIMIT ?", int(playersCount)/rng.UInt(2, 10)).Preload(teamRel).Find(&playersToRetire)
	if len(playersToRetire) == 0 {
		return
	}
	//TODO: maybe ad add a way to replace players
	pIds := make([]string, len(playersToRetire))
	retiring := make([]RetiredPlayer, len(playersToRetire))
	for i, p := range playersToRetire {
		retiring[i] = NewRetiredPlayerFromDto(p, indexedP, gameDate.Year(), leagueId, leagueName)
		pIds[i] = p.Id
	}
	lr.g.Create(&retiring)

	lr.g.Delete(&PHistoryDto{}, pIds)
	trx := lr.g.Delete(&PlayerDto{}, pIds)
	if trx.RowsAffected == 0 {
		panic("AAAARGH")
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
