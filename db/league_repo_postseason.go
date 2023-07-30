package db

import (
	"encoding/json"
	"fdsim/data"
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/libs"
	"fdsim/models"
	"fdsim/utils"
	"fmt"
	"strings"
	"time"

	"golang.org/x/exp/maps"
)

type TeamLosingPlayersMap map[string]map[models.Role]int

func (lr *LeagueRepo) PostSeason(game *models.Game) *models.League {
	// TODO: maybe inject this
	rng := libs.NewRng(time.Now().Unix())
	dbEvents := []DbEventDto{}

	league := lr.ById(game.LeagueId)
	leagueName := league.Name
	winnerTeamId := league.TableRow(0).Team.Id
	teamsC := league.Table.Len()
	losersIds := []string{
		league.TableRow(teamsC - 1).Team.Id,
		league.TableRow(teamsC - 2).Team.Id,
		league.TableRow(teamsC - 3).Team.Id,
	}

	var playersStats []StatRowDto
	lr.g.Model(&StatRowDto{}).
		Preload(playerRel).
		Preload(teamRel).
		Order("goals desc, played asc, team_id asc").
		Find(&playersStats)

	var mvp StatRowDto
	lr.g.Model(&StatRowDto{}).
		Preload(playerRel).
		Preload(teamRel).
		Order("score desc, played desc").
		First(&mvp)

	resultEvents := lr.createLeagueHistory(league, mvp, playersStats[:3])
	dbEvents = append(dbEvents, resultEvents...)

	resultEvents, indexedP := lr.convertStatsToHistory(game, leagueName, playersStats)
	dbEvents = append(dbEvents, resultEvents...)

	teamLostPlayers := TeamLosingPlayersMap{}

	resultEvents = lr.playersEndOfSeason(game, teamLostPlayers, winnerTeamId, leagueName, losersIds, rng)
	dbEvents = append(dbEvents, resultEvents...)

	resultEvents = lr.retirePlayers(indexedP, game, leagueName, teamLostPlayers, rng)
	dbEvents = append(dbEvents, resultEvents...)

	// add young players/replace retired
	resultEvents = lr.generateYoungReplacements(game, teamLostPlayers, rng)
	dbEvents = append(dbEvents, resultEvents...)

	//TODO: store fd info/stats
	resultEvents = lr.updateFDInfo(game, leagueName)
	dbEvents = append(dbEvents, resultEvents...)

	if len(dbEvents) > 0 {
		lr.g.Model(&DbEventDto{}).Create(&dbEvents)
	}

	return lr.createNewLeague(game)
}

func (lr *LeagueRepo) createLeagueHistory(league *models.League, mvp StatRowDto, scorers []StatRowDto) []DbEventDto {
	lh := NewLHistoryDtoFromLeague(league, mvp, scorers)
	lr.g.Create(lh)

	return []DbEventDto{}
}

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
	lr.g.Where("id = ?", oldLeague.Id).Delete(&LeagueDto{})

	nl, _ := lr.ByIdFull(newLeague.Id)

	return nl
}
func (lr *LeagueRepo) generateYoungReplacements(game *models.Game, tm TeamLosingPlayersMap, rng *libs.Rng) []DbEventDto {
	pg := generators.NewPeopleGenSeeded(rng)
	fdTeamId := game.GetTeamIdOrEmpty()

	pToAdd := []PlayerDto{}
	for teamId, rc := range tm {
		for role, count := range rc {
			if count == 0 {
				continue
			}
			effectiveCount := rng.UInt(2, count)
			ps := pg.YoungPlayersWithRole(effectiveCount, role)
			for _, p := range ps {
				pdto := DtoFromPlayer(p)
				nteamId := teamId
				pdto.TeamId = &nteamId
				pToAdd = append(pToAdd, pdto)

				if nteamId == fdTeamId {
					// add to list of new players to report to FD
				}
			}
		}
	}

	lr.g.Create(&pToAdd)

	return []DbEventDto{}
}

// Checks contracts and if 0 put them on the free market
// grows players skills and fame
func (lr *LeagueRepo) playersEndOfSeason(
	game *models.Game, tm TeamLosingPlayersMap,
	winningTeamId, leagueName string, losersIds []string,
	rng *libs.Rng,
) []DbEventDto {
	losersStr := strings.Join(losersIds, ", ")
	fdTeamId := game.GetTeamIdOrEmpty()
	pg := generators.NewPeopleGenSeeded(rng)

	lr.g.Exec("update player_dtos set y_contract = y_contract - 1 where team_id is not null")

	//TODO: this is the slowest of them all
	// TODO: Add events for players team
	// ITERATE OVER EVERY PLAYER
	var allPlayers []*PlayerDto
	lr.g.Model(&PlayerDto{}).Preload(teamRel).Find(&allPlayers)

	trophies := []TrophyDto{}
	for _, p := range allPlayers {
		// check if player is FD Managed
		isFdTeamPlayer := false
		if p.TeamId != nil && *p.TeamId == fdTeamId {
			isFdTeamPlayer = true
		}

		// Increase Players Skill if Young
		if p.Age < 22 {
			nSkill := utils.NewPerc(p.Skill + rng.UInt(2, 8))
			p.Skill = nSkill.Val()
		}

		// Decrease Players Skill if Older
		if p.Age > 29 {
			nSkill := utils.NewPerc(p.Skill - rng.UInt(1, 5))
			p.Skill = nSkill.Val()
		}

		// If it has been the whole year with no team
		if p.TeamId == nil {
			nSkill := utils.NewPerc(p.Skill + rng.PlusMinusVal(3, 50))
			p.Skill = nSkill.Val()
		}

		// Calculate new Value and new Ideal Wage
		p.IdealWage = pg.GetWage(p.Skill, p.Age, false).Val
		p.Value = pg.GetValue(p.Skill, p.Age).Val

		// If player won, increase fame
		if p.TeamId != nil && *p.TeamId == winningTeamId {
			nFame := utils.NewPerc(p.Fame + rng.UInt(5, 10))
			p.Fame = nFame.Val()
			trophies = append(trophies, NewTrophyDto(p.Id, game.LeagueId, leagueName, *p.TeamId, p.Team.Name, game.Date.Year()))

			// if withing last 3 teams decrease fame
		} else if p.TeamId != nil && strings.Contains(losersStr, *p.TeamId) {
			nFame := utils.NewPerc(p.Fame - rng.UInt(2, 5))
			p.Fame = nFame.Val()

			// otherwise randomly increase or decrease
		} else {
			nFame := utils.NewPerc(p.Fame + rng.PlusMinusVal(2, 50))
			p.Fame = nFame.Val()
		}

		// contract expiry
		if p.YContract < 1 && p.TeamId != nil {
			// 50% chance autorenew if not managed by FD
			if !isFdTeamPlayer && rng.ChanceI(50) {
				p.YContract = 1
			} else {
				countPlayerLoss(p, tm)
				p.TeamId = nil
				p.Team = nil
				p.Wage = 0
			}
		}

	}
	lr.g.Save(allPlayers)
	lr.g.Save(trophies)

	return []DbEventDto{}
}

func (lr *LeagueRepo) retirePlayers(indexedP map[string]PHistoryDto, game *models.Game, leagueName string, tm TeamLosingPlayersMap, rng *libs.Rng) []DbEventDto {
	leagueId := game.LeagueId
	gameDate := game.Date
	fdTeamId := game.GetTeamIdOrEmpty()

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
		return []DbEventDto{}
	}
	//TODO: maybe ad add a way to replace players
	pIds := make([]string, len(playersToRetire))
	retiring := make([]RetiredPlayerDto, len(playersToRetire))
	playerTeamRetired := []*models.PNPH{}
	for i, p := range playersToRetire {
		retiring[i] = NewRetiredPlayerFromDto(p, indexedP, gameDate.Year(), leagueId, leagueName)
		pIds[i] = p.Id
		if p.TeamId != nil && *p.TeamId == fdTeamId {
			playerTeamRetired = append(playerTeamRetired, p.PlayerPH())
		}

		countPlayerLoss(&p, tm)
	}
	lr.g.Create(&retiring)

	lr.g.Delete(&PHistoryDto{}, pIds)
	lr.g.Delete(&PlayerDto{}, pIds)

	events := []DbEventDto{}
	if len(playerTeamRetired) > 0 {
		data, _ := json.Marshal(&playerTeamRetired)
		events = append(
			events,
			NewDbEventDto(
				DbEvPlRetiredFdTeam, game.BaseCountry, string(data),
				models.EventParams{TeamName: game.Team.Name},
				gameDate.Add(enums.A_day),
			),
		)
	}

	return events
}

func countPlayerLoss(p *PlayerDto, tm TeamLosingPlayersMap) {
	if p.TeamId == nil {
		return
	}

	teamId := *p.TeamId

	if pc, ok := tm[teamId]; ok {
		pc[p.Role]++
	} else {
		tm[teamId] = models.NewEmptyRoleCounter()
		tm[teamId][p.Role]++
	}
}

func (lr *LeagueRepo) convertStatsToHistory(game *models.Game, leagueName string, playersStats []StatRowDto) (
	[]DbEventDto,
	map[string]PHistoryDto,
) {
	leagueId := game.LeagueId
	gameDate := game.Date

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
	phrows = append(phrows, lr.backFillEmptyStatsHistory(leagueId, leagueName, gameDate)...)

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

	return []DbEventDto{}, indexedPHRows
}

func (lr *LeagueRepo) backFillEmptyStatsHistory(leagueId, leagueName string, gameDate time.Time) []PHistoryDto {
	var ps []PlayerDto
	phrows := []PHistoryDto{}
	// Players with no stats
	lr.g.Raw("select pd.* from player_dtos pd left join stat_row_dtos srd on pd.id  = srd.player_id where srd.player_id is null").Preload(teamRel).Find(&ps)
	for _, p := range ps {
		if p.Team != nil {
			phrows = append(phrows, *NewEmptyHistoryRow(p.Id, leagueId, leagueName, *p.TeamId, p.Team.Name, gameDate.Year()))
		} else {
			phrows = append(phrows, *NewEmptyHistoryRow(p.Id, leagueId, leagueName, "", "FREE AGENT", gameDate.Year()))
		}
	}

	return phrows
}

func (lr *LeagueRepo) updateFDInfo(game *models.Game, leagueName string) []DbEventDto {
	game.Age++

	if !game.IsEmployed() {
		return []DbEventDto{}
	}

	var stat FDStatRowDto
	lr.g.Model(&FDStatRowDto{}).Order("hired_date desc").First(&stat)
	h := NewFDHistoryDto(stat)
	h.UpdateEndOfSeason(game.LeagueId, leagueName, game.Wage, game.Date)
	lr.g.Save(&h)

	game.YContract -= 1
	if game.YContract == 0 {
		game.UnsetTeamContract()
	} else {
		newStat := models.NewFDStatRow(game.Date, stat.TeamId, stat.TeamName)
		lr.g.Model(&FDStatRowDto{}).Create(DtoFromFDStatRow(newStat))
	}

	return []DbEventDto{}
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
