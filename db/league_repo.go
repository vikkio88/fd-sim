package db

import (
	"fdsim/conf"
	"fdsim/models"
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

func (lr *LeagueRepo) PostSeasonStats(leagueId string, gameDate time.Time) {
	var playersStats []StatRowDto
	lr.g.Model(&StatRowDto{}).Preload(teamRel).Find(&playersStats)

	var pHistoryRows []PHistoryDto
	lr.g.Model(&PHistoryDto{}).Find(&pHistoryRows)
	indexedPHRows := indexPHistoryRows(pHistoryRows)

	for _, s := range playersStats {
		if existingRow, ok := indexedPHRows[s.PlayerId]; ok {
			//TODO: maybe move to pointer
			existingRow.Update(s, gameDate)
			indexedPHRows[s.PlayerId] = existingRow

		} else {
			indexedPHRows[s.PlayerId] = DtoFromPHistoryRow(
				models.NewPHistoryRow(
					s.StatRow(),
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
			existingRow.Update(s, leagueId, gameDate)
			indexedTHRows[s.TeamId] = existingRow

		} else {
			indexedTHRows[s.TeamId] = DtoFromTHistoryRow(
				models.NewTHistoryRow(
					s.TPHRow(),
					leagueId,
					gameDate,
				),
			)
		}
	}
	throws := maps.Values(indexedTHRows)

	lr.cleanStats()
	lr.g.Save(phrows)
	lr.g.Save(throws)
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
