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

	g.AutoMigrate(&TeamDto{}, &PlayerDto{}, &CoachDto{})
	cache := map[string]interface{}{}
	return &Db{g, cache}
}

func (db *Db) TeamR() *TeamsRepo {
	if tr, ok := db.cache[teamRepoCacheKey]; ok {
		return tr.(*TeamsRepo)
	}
	tr := NewTeamsRepo(db.g)
	db.cache[teamRepoCacheKey] = tr

	return tr
}

func (db *Db) PlayerR() *PlayerRepo {
	if pr, ok := db.cache[playerRepoCacheKey]; ok {
		return pr.(*PlayerRepo)
	}
	pr := NewPlayerRepo(db.g)
	db.cache[playerRepoCacheKey] = pr

	return pr
}

func (db *Db) CoachR() *CoachRepo {
	if cr, ok := db.cache[coachRepoCacheKey]; ok {
		return cr.(*CoachRepo)
	}
	cr := NewCoachRepo(db.g)
	db.cache[coachRepoCacheKey] = cr

	return cr
}

func (db *Db) TruncateAll() {
	db.TeamR().Truncate()
	db.PlayerR().Truncate()
	db.CoachR().Truncate()
}
