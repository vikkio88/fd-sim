package db

import (
	"fdsim/models"
	"time"
)

// IDb ...
type IDb interface {
	TruncateAll()
	GameR() IGameRepo
	LeagueR() ILeagueRepo
	TeamR() ITeamRepo
	PlayerR() IPlayerRepo
	CoachR() ICoachRepo
}

// IGameRepo ...
type IGameRepo interface {
	Truncate()
	All() []*models.Game
	ById(id string) *models.Game
	Create(game *models.Game)
	Update(game *models.Game)
	AddEmails([]*models.Email)
	GetEmails() []*models.Email
	GetEmailById(id string) *models.Email
	DeleteEmail(id string)
	MarkEmailAsRead(id string)
	UpdateEmail(*models.Email)
	AddNews([]*models.News)
	GetNews() []*models.News
	GetNewsById(id string) *models.News
	DeleteNews(id string)
	MarkNewsAsRead(id string)
}

// ILeagueRepo ...
type ILeagueRepo interface {
	Truncate()
	PostRoundUpdate(r *models.Round, league *models.League)
	InsertOne(l *models.League)
	// Loads League with Teams (no players), Rounds (no Matches) and Table
	ById(id string) *models.League
	// Load a full League with all the info
	ByIdFull(id string) *models.League
	RoundCountByDate(date time.Time) int64
	RoundByIndex(league *models.League, index int) *models.RoundResult
	// get Round with all the results
	RoundWithResults(roundId string) *models.RPHTPH
	// map of matchids and result placeholders
	GetAllResults() models.ResultsPHMap
	GetMatchById(matchId string) *models.MatchComplete
	GetStatsForPlayer(playerId, leagueId string) *models.StatRow
	BestScorers(leagueId string) []*models.StatRowPH
	GetStats(leagueId string) models.StatsMap
	UpdateStats(stats models.StatsMap)
}

// ITeamRepo ...
type ITeamRepo interface {
	InsertOne(t *models.Team)
	Insert(teams []*models.Team)
	ById(id string) *models.Team
	Truncate()
	DeleteOne(id string)
	Delete(ids []string)
	Count() int64
	All() []*models.Team
}

// IPlayerRepo ...
type IPlayerRepo interface {
	InsertOne(p *models.Player)
	Insert(players []*models.Player)
	ById(id string) *models.PlayerWithTeam
	Truncate()
	DeleteOne(id string)
	Delete(ids []string)
	Count() int64
	FreeAgents() []*models.Player
	All() []*models.Player
}

// ICoachRepo ...
type ICoachRepo interface {
	InsertOne(c *models.Coach)
	Insert(coaches []*models.Coach)
	ById(id string) *models.Coach
	Truncate()
	DeleteOne(id string)
	Delete(ids []string)
	Count() int64
	FreeAgents() []*models.Coach
	All() []*models.Coach
}
