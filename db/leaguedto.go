package db

import (
	"fdsim/conf"
	"fdsim/models"

	"gorm.io/gorm"
)

type LeagueDto struct {
	Id   string `gorm:"primarykey;size:16"`
	Name string

	Teams     []TeamDto     `gorm:"foreignKey:league_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Rounds    []RoundDto    `gorm:"foreignKey:league_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TableRows []TableRowDto `gorm:"foreignKey:league_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	RPointer int
}

func DtoFromLeague(l *models.League) LeagueDto {
	trs := DtoFromTableRows(l.Table.Rows(), l.Id)

	rds := make([]RoundDto, len(l.Rounds))
	for i, r := range l.Rounds {
		rds[i] = DtoFromRoundPH(r, l.Id)
	}

	ldto := LeagueDto{
		Id:       l.Id,
		Name:     l.Name,
		RPointer: l.RPointer,

		Teams: DtoFromTeams(l.Teams(), l.Id),

		TableRows: trs,
		Rounds:    rds,
	}

	return ldto
}

func (l *LeagueDto) GetTeams() []*models.Team {
	ts := make([]*models.Team, len(l.Teams))

	for i, tdto := range l.Teams {
		ts[i] = tdto.Team()
	}
	return ts
}

func (l *LeagueDto) League() *models.League {
	league := models.NewLeagueWithData(l.Id, l.Name, l.GetTeams())

	league.RPointer = l.RPointer
	league.Table = TableFromTableRowsDto(l.TableRows)
	league.Rounds = RoundsPHFromDto(l.Rounds)

	return league
}

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

func (lr *LeagueRepo) RoundByIndex(league *models.League, index int) *models.RoundResult {
	var rdto RoundDto
	lr.g.Model(&RoundDto{}).Preload("Matches.Result").Where("`index` = ? AND league_id = ?", index, league.Id).Find(&rdto)
	return rdto.Round(league.TeamMap)
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
	lr.g.Model(&StatRowDto{}).Where("player_id = ? and league_id = ?", playerId, leagueId).First(&stat)
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
