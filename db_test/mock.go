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

	l := models.NewLeague(ts, utils.NewDate(2023, time.July, 1))
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
func (d *MockDb) MarketR() db.IMarketRepo {
	panic("not implemented")
}

type MockGameR struct{}

// StoreEvent implements db.IGameRepo.
func (*MockGameR) StoreEvent(db.DbEventDto) {
	panic("unimplemented")
}

// GetTransferMarketInfo implements db.IGameRepo.
func (*MockGameR) GetTransferMarketInfo() (*models.TransferMarketInfo, bool) {
	panic("unimplemented")
}

// GetEvents implements db.IGameRepo.
func (*MockGameR) GetEvents(time.Time) []db.DbEventDto {
	return []db.DbEventDto{}
}

// StoreEvents implements db.IGameRepo.
func (*MockGameR) StoreEvents([]db.DbEventDto) {
	panic("unimplemented")
}

// GetFDHistory implements db.IGameRepo.
func (*MockGameR) GetFDHistory() []*models.FDHistoryRow {
	panic("unimplemented")
}

// GetFDStats implements db.IGameRepo.
func (*MockGameR) GetFDStats() *models.FDStatRow {
	panic("unimplemented")
}

// AddStatRow implements db.IGameRepo.
func (*MockGameR) AddStatRow(*models.FDStatRow) {
	panic("unimplemented")
}

// GetActionsDueToday implements db.IGameRepo.
func (*MockGameR) GetActionsDueByDate(time.Time) []*models.Idable {
	return []*models.Idable{}
}

// TruncateNotifications implements db.IGameRepo.
func (*MockGameR) TruncateNotifications() {
	panic("unimplemented")
}

// DeleteAllNews implements db.IGameRepo.
func (*MockGameR) DeleteAllNews() {
	panic("unimplemented")
}

// UpdateEmail implements db.IGameRepo.
func (*MockGameR) UpdateEmail(*models.Email) {
}

// AddEmails implements db.IGameRepo.
func (*MockGameR) AddEmails([]*models.Email) {
	panic("unimplemented")
}

// AddNews implements db.IGameRepo.
func (*MockGameR) AddNews([]*models.News) {
	panic("unimplemented")
}

// DeleteEmail implements db.IGameRepo.
func (*MockGameR) DeleteEmail(id string) {
	panic("unimplemented")
}

// DeleteNews implements db.IGameRepo.
func (*MockGameR) DeleteNews(id string) {
	panic("unimplemented")
}

// GetEmailById implements db.IGameRepo.
func (*MockGameR) GetEmailById(id string) *models.Email {
	panic("unimplemented")
}

// GetEmails implements db.IGameRepo.
func (*MockGameR) GetEmails() []*models.Email {
	panic("unimplemented")
}

// GetNews implements db.IGameRepo.
func (*MockGameR) GetNews() []*models.News {
	panic("unimplemented")
}

// GetNewsById implements db.IGameRepo.
func (*MockGameR) GetNewsById(id string) *models.News {
	panic("unimplemented")
}

// MarkEmailAsRead implements db.IGameRepo.
func (*MockGameR) MarkEmailAsRead(id string) {
	panic("unimplemented")
}

// MarkNewsAsRead implements db.IGameRepo.
func (*MockGameR) MarkNewsAsRead(id string) {
	panic("unimplemented")
}

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

}

type MockTeamRepo struct {
	Teams []*models.Team
}

func (m *MockTeamRepo) Update(team *models.Team) {
	m.Teams = append(m.Teams, team)
}

// GetByIds implements db.ITeamRepo.
func (m *MockTeamRepo) GetByIds(ids []string) []*models.Team {
	ts := m.Teams

	m.Teams = []*models.Team{}

	return ts
}

// TableRow implements db.ITeamRepo.
func (*MockTeamRepo) TableRow(id string) *models.TPHRow {
	panic("unimplemented")
}

// GetRandom implements db.ITeamRepo.
func (*MockTeamRepo) GetRandom() *models.TPH {
	panic("unimplemented")
}

// OneByFame implements db.ITeamRepo.
func (r *MockTeamRepo) OneByFame(utils.Perc) *models.TPH {
	t := r.Teams[0].PH()

	return &t
}

func (r *MockTeamRepo) InsertOne(t *models.Team) {
	panic("not implemented")
}

func (r *MockTeamRepo) Insert(teams []*models.Team) {
	panic("not implemented")
}

func (r *MockTeamRepo) ById(id string) (*models.TeamDetailed, bool) {
	return &models.TeamDetailed{Team: *r.Teams[0]}, true
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

func (r *MockPlayerRepo) ById(id string) (*models.PlayerDetailed, bool) {
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

// HistoryById implements db.ILeagueRepo.
func (*MockLeagueRepo) HistoryById(id string) (*models.LeagueHistory, bool) {
	panic("unimplemented")
}

// PostSeason implements db.ILeagueRepo.
func (*MockLeagueRepo) PostSeason(*models.Game) *models.League {
	panic("implement")
}

// RoundWithResults implements db.ILeagueRepo.
func (*MockLeagueRepo) RoundWithResults(roundId string) *models.RPHTPH {
	panic("unimplemented")
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
func (r *MockLeagueRepo) ByIdFull(id string) (*models.League, bool) {
	return r.League, true
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
