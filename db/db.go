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
)

type Db struct {
	g     *gorm.DB
	cache map[string]interface{}
}

func NewDb(fileName string) *Db {
	g, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s?_foreign_keys=on", fileName)), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	g.AutoMigrate(
		&LeagueDto{}, &MatchDto{}, &ResultDto{},
		&TableRowDto{}, &RoundDto{}, &TeamDto{},
		//For some reason this kills everything
		&PlayerDto{}, &CoachDto{}, &GameDto{},
	)
	cache := map[string]interface{}{}
	return &Db{g, cache}
}

func (db *Db) TruncateAll() {
	db.GameR().Truncate()
	db.LeagueR().Truncate()
	db.TeamR().Truncate()
	db.PlayerR().Truncate()
	db.CoachR().Truncate()
}

func (db *Db) GameR() *GameRepo {
	if repo, ok := db.cache[gameRepoCacheKey]; ok {
		return repo.(*GameRepo)
	}
	repo := NewGameRepo(db.g)
	db.cache[gameRepoCacheKey] = repo

	return repo
}

func (db *Db) LeagueR() *LeagueRepo {
	if repo, ok := db.cache[leagueRepoCacheKey]; ok {
		return repo.(*LeagueRepo)
	}
	repo := NewLeagueRepo(db.g)
	db.cache[leagueRepoCacheKey] = repo

	return repo
}

func (db *Db) TeamR() *TeamsRepo {
	if repo, ok := db.cache[teamRepoCacheKey]; ok {
		return repo.(*TeamsRepo)
	}
	repo := NewTeamsRepo(db.g)
	db.cache[teamRepoCacheKey] = repo

	return repo
}

func (db *Db) PlayerR() *PlayerRepo {
	if repo, ok := db.cache[playerRepoCacheKey]; ok {
		return repo.(*PlayerRepo)
	}
	repo := NewPlayerRepo(db.g)
	db.cache[playerRepoCacheKey] = repo

	return repo
}

func (db *Db) CoachR() *CoachRepo {
	if repo, ok := db.cache[coachRepoCacheKey]; ok {
		return repo.(*CoachRepo)
	}
	repo := NewCoachRepo(db.g)
	db.cache[coachRepoCacheKey] = repo

	return repo
}
