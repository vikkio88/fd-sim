package db_test

// shared set of utils to mock db around app

import (
	"fdsim/db"
	"fdsim/generators"
	"fdsim/models"
	"fdsim/utils"
	"time"
)

type MockDb struct {
	Team   *MockTeamRepo
	Player *MockPlayerRepo
	League *MockLeagueRepo
}

func NewMockDbSeeded(seed int64) *MockDb {
	tg := generators.NewTeamGen(seed)
	ts := tg.Teams(2)

	l := models.NewLeague("Mock", ts, utils.NewDate(2023, time.July, 1))
	league := &MockLeagueRepo{League: l}
	team := &MockTeamRepo{Teams: ts}
	return &MockDb{
		Team:   team,
		League: league,
	}
}

func NewMockDbWithTeams(teams []*models.Team) *MockDb {
	return &MockDb{
		Team: &MockTeamRepo{Teams: teams},
	}
}

func NewMockDbWithPlayers(players []*models.Player) *MockDb {
	return &MockDb{
		Player: &MockPlayerRepo{Players: players},
	}
}

func (d *MockDb) TruncateAll() {
	panic("not implemented")
}
func (d *MockDb) GameR() db.IGameRepo {
	return &MockGameR{}
}
func (d *MockDb) LeagueR() db.ILeagueRepo {
	return d.League
}
func (d *MockDb) TeamR() db.ITeamRepo {
	return d.Team
}
func (d *MockDb) PlayerR() db.IPlayerRepo {
	panic("not implemented")
}
func (d *MockDb) CoachR() db.ICoachRepo {
	panic("not implemented")
}

type MockGameR struct{}

func (r *MockGameR) Truncate() {
	panic("not implemented")
}
func (r *MockGameR) All() []*models.Game {
	panic("not implemented")
}
func (r *MockGameR) ById(id string) *models.Game {
	panic("not implemented")
}
func (r *MockGameR) Create(game *models.Game) {
	panic("not implemented")
}
func (r *MockGameR) Update(game *models.Game) {
	panic("not implemented")
}

type MockTeamRepo struct {
	Teams []*models.Team
}

func (r *MockTeamRepo) InsertOne(t *models.Team) {
	panic("not implemented")
}

func (r *MockTeamRepo) Insert(teams []*models.Team) {
	panic("not implemented")
}

func (r *MockTeamRepo) ById(id string) *models.Team {
	return r.Teams[0]
}

func (r *MockTeamRepo) Truncate() {
	panic("not implemented")
}

func (r *MockTeamRepo) DeleteOne(id string) {
	panic("not implemented")
}

func (r *MockTeamRepo) Delete(ids []string) {
	panic("not implemented")
}

func (r *MockTeamRepo) Count() int64 {
	return int64(len(r.Teams))
}

func (r *MockTeamRepo) All() []*models.Team {
	return r.Teams
}

type MockPlayerRepo struct{ Players []*models.Player }

func (r *MockPlayerRepo) InsertOne(p *models.Player) {
	panic("not implemented")
}

func (r *MockPlayerRepo) Insert(players []*models.Player) {
	panic("not implemented")
}

func (r *MockPlayerRepo) ById(id string) *models.PlayerWithTeam {
	panic("not implemented")
}

func (r *MockPlayerRepo) Truncate() {
	panic("not implemented")
}

func (r *MockPlayerRepo) DeleteOne(id string) {
	panic("not implemented")
}

func (r *MockPlayerRepo) Delete(ids []string) {
	panic("not implemented")
}

func (r *MockPlayerRepo) Count() int64 {
	panic("not implemented")
}

func (r *MockPlayerRepo) FreeAgents() []*models.Player {
	panic("not implemented")
}

func (r *MockPlayerRepo) All() []*models.Player {
	return r.Players
}

type MockLeagueRepo struct {
	League      *models.League
	RoundCount  int64
	RoundResult *models.RoundResult
	Stats       *models.StatsMap
}

func (r *MockLeagueRepo) Truncate() {
	panic("not implemented")
}

func (repo *MockLeagueRepo) PostRoundUpdate(r *models.Round, league *models.League) {
	repo.League = league
}

func (r *MockLeagueRepo) InsertOne(l *models.League) {
	panic("not implemented")
}

// Loads League with Teams (no players), Rounds (no Matches) and Table
func (r *MockLeagueRepo) ById(id string) *models.League {
	panic("not implemented")
}

// Load a full League with all the info
func (r *MockLeagueRepo) ByIdFull(id string) *models.League {
	return r.League
}

func (r *MockLeagueRepo) RoundCountByDate(date time.Time) int64 {
	var c int64 = 0

	for _, r := range r.League.Rounds {
		if r.Date == date {
			c++
		}
	}

	return c
}

func (r *MockLeagueRepo) RoundByIndex(league *models.League, index int) *models.RoundResult {
	return r.RoundResult
}

// map of matchids and result placeholders
func (r *MockLeagueRepo) GetAllResults() models.ResultsPHMap {
	panic("not implemented")
}

func (r *MockLeagueRepo) GetMatchById(matchId string) *models.MatchComplete {
	panic("not implemented")
}

func (r *MockLeagueRepo) GetStatsForPlayer(playerId string, leagueId string) *models.StatRow {
	panic("not implemented")
}

func (r *MockLeagueRepo) BestScorers(leagueId string) []*models.StatRowPH {
	panic("not implemented")
}

func (r *MockLeagueRepo) GetStats(leagueId string) models.StatsMap {
	if r.Stats == nil {
		return models.StatsMap{}
	}

	return *r.Stats
}

func (r *MockLeagueRepo) UpdateStats(stats models.StatsMap) {
	r.Stats = &stats
}
