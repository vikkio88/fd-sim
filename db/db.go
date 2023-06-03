package db

import (
	"fdsim/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Db struct {
	g *gorm.DB
}

func NewDb(fileName string) Db {
	g, err := gorm.Open(sqlite.Open(fileName), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	g.AutoMigrate(&TeamDto{}, &PlayerDto{}, &CoachDto{})

	return Db{g}
}

func (d *Db) TeamsCount() int64 {
	var c int64
	d.g.Model(&TeamDto{}).Count(&c)

	return c
}

func (d *Db) InsertTeam(t *models.Team) {
	tdto := DtoFromTeam(t)

	d.g.Create(&tdto)
}

func (d *Db) InsertManyTeams(ts []*models.Team) {
	tdtos := make([]TeamDto, len(ts))
	for i, t := range ts {
		tdtos[i] = DtoFromTeam(t)
	}

	d.g.Create(&tdtos)
}

func (d *Db) TeamById(id string) *models.Team {
	var t TeamDto
	d.g.Model(&TeamDto{}).Preload("Players").Preload("Coach").Find(&t, "Id = ?", id)

	return t.Team()
}

func (d *Db) AllTeams() []TeamDto {
	var t []TeamDto
	d.g.Model(&TeamDto{}).Preload("Players").Find(&t)

	return t
}
