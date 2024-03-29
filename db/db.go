package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	teamRepoCacheKey   = "db_TR"
	playerRepoCacheKey = "db_PR"
	coachRepoCacheKey  = "db_CR"

	gameRepoCacheKey   = "db_GR"
	leagueRepoCacheKey = "db_LR"
	marketRepoCacheKey = "db_MKR"
)

type Db struct {
	g     *gorm.DB
	cache map[string]interface{}
}

func NewDb(fileName string) IDb {
	g, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s?_foreign_keys=on", fileName)), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	migrate(g)
	cache := map[string]interface{}{}
	return &Db{g, cache}
}

func (db *Db) TruncateAll() {
	db.GameR().Truncate()
	db.GameR().TruncateNotifications()
	db.LeagueR().Truncate()
	db.TeamR().Truncate()
	db.PlayerR().Truncate()
	db.CoachR().Truncate()
	db.MarketR().Truncate()
}

func (db *Db) GameR() IGameRepo {
	if repo, ok := db.cache[gameRepoCacheKey]; ok {
		return repo.(*GameRepo)
	}
	repo := NewGameRepo(db.g)
	db.cache[gameRepoCacheKey] = repo

	return repo
}

func (db *Db) LeagueR() ILeagueRepo {
	if repo, ok := db.cache[leagueRepoCacheKey]; ok {
		return repo.(*LeagueRepo)
	}
	repo := NewLeagueRepo(db.g)
	db.cache[leagueRepoCacheKey] = repo

	return repo
}

func (db *Db) TeamR() ITeamRepo {
	if repo, ok := db.cache[teamRepoCacheKey]; ok {
		return repo.(*TeamRepo)
	}
	repo := NewTeamsRepo(db.g)
	db.cache[teamRepoCacheKey] = repo

	return repo
}

func (db *Db) PlayerR() IPlayerRepo {
	if repo, ok := db.cache[playerRepoCacheKey]; ok {
		return repo.(*PlayerRepo)
	}
	repo := NewPlayerRepo(db.g)
	db.cache[playerRepoCacheKey] = repo

	return repo
}

func (db *Db) CoachR() ICoachRepo {
	if repo, ok := db.cache[coachRepoCacheKey]; ok {
		return repo.(*CoachRepo)
	}
	repo := NewCoachRepo(db.g)
	db.cache[coachRepoCacheKey] = repo

	return repo
}

func (db *Db) MarketR() IMarketRepo {
	if repo, ok := db.cache[marketRepoCacheKey]; ok {
		return repo.(*MarketRepo)
	}
	repo := NewMarketRepo(db.g)
	db.cache[marketRepoCacheKey] = repo

	return repo
}
