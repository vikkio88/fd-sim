package db

import (
	"fdsim/conf"
	"fdsim/data"
	"fdsim/libs"
	"fdsim/models"
	"fmt"
	"time"

	"golang.org/x/exp/maps"
	"gorm.io/gorm"
)

type LeagueRepo struct {
	g *gorm.DB
}

func NewLeagueRepo(g *gorm.DB) *LeagueRepo {
	return &LeagueRepo{
		g,
	}
}

func (lr *LeagueRepo) Truncate() {
	lr.g.Where("1 = 1").Delete(&LeagueDto{})
	lr.g.Where("1 = 1").Delete(&TableRowDto{})
	lr.g.Where("1 = 1").Delete(&ResultDto{})
	lr.g.Where("1 = 1").Delete(&MatchDto{})
	lr.g.Where("1 = 1").Delete(&RoundDto{})
}

func (lr *LeagueRepo) PostRoundUpdate(r *models.Round, league *models.League) {
	table := DtoFromTableRows(league.Table.Rows(), league.Id)
	lr.g.Save(table)
	rdto := DtoFromRound(r, league.Id)
	lr.g.Save(rdto)

	lr.g.Model(&LeagueDto{}).Where("Id = ?", league.Id).Update("RPointer", league.RPointer)
}

func (lr *LeagueRepo) InsertOne(l *models.League) {
	ldto := DtoFromLeague(l)
	lr.g.Create(&ldto)
}

// Loads League with Teams (no players), Rounds (no Matches) and Table
func (lr *LeagueRepo) ById(id string) *models.League {
	var ldto LeagueDto
	lr.g.Model(&LeagueDto{}).
		Preload(teamsRel).
		Preload(roundsRel).
		Preload(tableRowsRel).
		Find(&ldto, "Id = ?", id)

	return ldto.League()
}

// Load a full League with all the info
func (lr *LeagueRepo) ByIdFull(id string) *models.League {
	var ldto LeagueDto
	lr.g.Model(&LeagueDto{}).
		Preload(teamsRel).
		Preload(teamsAndPlayersRel).
		Preload(teamsAndCoachRel).
		Preload(roundsAndMatchesRel).
		Preload(tableRowsRel).
		Find(&ldto, "Id = ?", id)
	return ldto.League()
}

func (lr *LeagueRepo) RoundWithResults(roundId string) *models.RPHTPH {
	var dto RoundDto
	lr.g.Model(&RoundDto{}).
		Preload("Matches").
		Preload("Matches.HomeTeam").
		Preload("Matches.AwayTeam").
		Preload("Matches.Result").
		Where("id = ?", roundId).Find(&dto)
	return dto.RoundPHTPH()
}

func (lr *LeagueRepo) RoundByIndex(league *models.League, index int) *models.RoundResult {
	var rdto RoundDto
	lr.g.Model(&RoundDto{}).Preload("Matches.Result").Where("`index` = ? AND league_id = ?", index, league.Id).Find(&rdto)
	return rdto.Round(league.TeamMap)
}

func (lr *LeagueRepo) RoundCountByDate(date time.Time) int64 {
	var count int64
	lr.g.Model(&RoundDto{}).Where("date = ?", date).Count(&count)
	return count
}

// map of matchids and result placeholders
func (lr *LeagueRepo) GetAllResults() models.ResultsPHMap {
	var dtos []ResultDto
	lr.g.Model(&ResultDto{}).Find(&dtos)

	return ResultsMapPHFromDtos(dtos)
}

func (lr *LeagueRepo) GetMatchById(matchId string) *models.MatchComplete {
	var m MatchDto
	lr.g.Model(&MatchDto{}).
		Preload("Round").
		Preload("HomeTeam.Players").
		Preload("AwayTeam.Players").
		Preload("Result").
		Find(&m, "id = ?", matchId)

	return m.MatchComplete()
}

func (lr *LeagueRepo) GetStatsForPlayer(playerId, leagueId string) *models.StatRow {
	var stat StatRowDto
	lr.g.Model(&StatRowDto{}).Where("player_id = ? and league_id = ?", playerId, leagueId).Find(&stat)
	return stat.StatRow()
}

func (lr *LeagueRepo) BestScorers(leagueId string) []*models.StatRowPH {
	var sdtos []StatRowDto
	lr.g.Model(&StatRowDto{}).
		Preload(playerRel).
		Preload(teamRel).
		Where("league_id = ? and goals > 0", leagueId).
		Order("goals desc, played asc, team_id asc").
		Limit(conf.StatsRowsLimit).
		Find(&sdtos)

	return StatRowsPhFromDtos(sdtos)
}
func (lr *LeagueRepo) GetStats(leagueId string) models.StatsMap {
	var sdtos []StatRowDto
	lr.g.Model(&StatRowDto{}).Where("league_id = ?", leagueId).Find(&sdtos)
	return StatsMapFromDtos(sdtos)
}

func (lr *LeagueRepo) UpdateStats(stats models.StatsMap) {
	sdtos := DtosFromStatsMap(stats)
	lr.g.Save(sdtos)
}

func (lr *LeagueRepo) PostSeason(game *models.Game, leagueName string) *models.League {
	// maybe store player retired
	gameDate := game.Date
	indexedP, _ := lr.convertStatsToHistory(leagueName, gameDate, game.LeagueId)
	lr.retirePlayers(indexedP, game.LeagueId, leagueName, gameDate)

	// add young players

	//TODO: implement this
	lr.playersEndOfSeason(gameDate)

	return lr.createNewLeague(game)
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
