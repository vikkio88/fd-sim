package db

import (
	"fdsim/models"
	"fdsim/utils"
	"time"
)

// IDb
type IDb interface {
	TruncateAll()
	GameR() IGameRepo
	LeagueR() ILeagueRepo
	TeamR() ITeamRepo
	PlayerR() IPlayerRepo
	CoachR() ICoachRepo

	MarketR() IMarketRepo
}

type IMarketRepo interface {
	Truncate()
	GetTransferMarketInfo() (*models.TransferMarketInfo, bool)
	AddOffer(OfferDto)
	SaveOffer(*models.Offer)
	DeleteOffer(*models.Offer)
	GetOffersByPlayerTeamId(playerId string, offeringTeamId string) (*models.Offer, bool)
	GetOffersByOfferingTeamId(string) []*models.Offer

	ApplyTransfer(*models.Offer) *models.TransferResult
}

type IGameRepo interface {
	Truncate()
	TruncateNotifications()
	All() []*models.Game
	ById(id string) *models.Game
	Create(game *models.Game)
	Update(game *models.Game)
	AddStatRow(row *models.FDStatRow)
	GetActionsDueByDate(time.Time) []*models.Idable
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
	DeleteAllNews()
	MarkNewsAsRead(id string)
	GetFDStats() *models.FDStatRow
	GetFDHistory() []*models.FDHistoryRow

	GetEvents(time.Time) []DbEventDto
	StoreEvent(DbEventDto)
	StoreEvents([]DbEventDto)
}

// ILeagueRepo ...
type ILeagueRepo interface {
	Truncate()
	PostRoundUpdate(r *models.Round, league *models.League)
	InsertOne(l *models.League)
	// Loads League with Teams (no players), Rounds (no Matches) and Table
	ById(id string) *models.League
	// Load a full League with all the info
	ByIdFull(id string) (*models.League, bool)
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
	HistoryById(id string) (*models.LeagueHistory, bool)

	// execute PostSeason actions returns new league
	PostSeason(game *models.Game) *models.League
}

// ITeamRepo ...
type ITeamRepo interface {
	InsertOne(t *models.Team)
	Insert(teams []*models.Team)
	OneByFame(utils.Perc) *models.TPH
	GetRandom() *models.TPH
	ById(id string) (*models.TeamDetailed, bool)
	GetByIds(ids []string) []*models.Team
	TableRow(teamId string) *models.TPHRow
	Truncate()
	DeleteOne(id string)
	Delete(ids []string)
	Update(teams *models.Team)
	Count() int64
	All() []*models.Team
}

// IPlayerRepo ...
type IPlayerRepo interface {
	InsertOne(p *models.Player)
	Insert(players []*models.Player)
	ById(id string) (*models.PlayerDetailed, bool)
	RetiredById(id string) (*models.RetiredPlayer, bool)
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
