package db_test

// shared set of utils to mock db around app

import (
	"fdsim/db"
	"fdsim/models"
)

type MockDb struct {
	teams   *MockTeamRepo
	players *MockPlayerRepo
}

func NewMockDbWithTeams(teams []*models.Team) *MockDb {
	return &MockDb{
		teams: &MockTeamRepo{Teams: teams},
	}
}

func NewMockDbWithPlayers(players []*models.Player) *MockDb {
	return &MockDb{
		players: &MockPlayerRepo{Players: players},
	}
}

func (d *MockDb) TruncateAll() {
	panic("not implemented")
}
func (d *MockDb) GameR() db.IGameRepo {
	return &MockGameR{}
}
func (d *MockDb) LeagueR() db.ILeagueRepo {
	panic("not implemented")
}
func (d *MockDb) TeamR() db.ITeamRepo {
	return d.teams
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
