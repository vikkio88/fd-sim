package db

import (
	"fdsim/models"

	"gorm.io/gorm"
)

type PlayerRepo struct {
	g *gorm.DB
}

func NewPlayerRepo(g *gorm.DB) *PlayerRepo {
	return &PlayerRepo{
		g,
	}
}

func (pr *PlayerRepo) InsertOne(p *models.Player) {
	pdto := DtoFromPlayer(p)

	pr.g.Create(&pdto)
}

func (pr *PlayerRepo) Insert(players []*models.Player) {
	pdtos := make([]PlayerDto, len(players))
	for i, p := range players {
		pdtos[i] = DtoFromPlayer(p)
	}

	pr.g.Create(&pdtos)
}

func (pr *PlayerRepo) ById(id string) *models.PlayerDetailed {
	var p PlayerDto
	pr.g.Model(&PlayerDto{}).Preload(teamRel).Preload("History").Find(&p, "Id = ?", id)

	return p.PlayerDetailed()
}

func (pr *PlayerRepo) Truncate() {
	pr.g.Where("1 = 1").Delete(&PlayerDto{})
}

func (pr *PlayerRepo) DeleteOne(id string) {
	pr.g.Delete(&PlayerDto{}, id)
}

func (pr *PlayerRepo) Delete(ids []string) {
	pr.g.Delete(&PlayerDto{}, ids)
}

func (pr *PlayerRepo) Count() int64 {
	var c int64
	pr.g.Model(&PlayerDto{}).Count(&c)

	return c
}

func (pr *PlayerRepo) FreeAgents() []*models.Player {
	var pdtos []PlayerDto
	pr.g.Model(&PlayerDto{}).Where("team_id", nil).Find(&pdtos)

	ps := make([]*models.Player, len(pdtos))
	for i, t := range pdtos {
		ps[i] = t.Player()
	}
	return ps
}

func (pr *PlayerRepo) All() []*models.Player {
	var pdtos []PlayerDto
	pr.g.Model(&PlayerDto{}).Find(&pdtos)

	ps := make([]*models.Player, len(pdtos))
	for i, t := range pdtos {
		ps[i] = t.Player()
	}
	return ps
}
