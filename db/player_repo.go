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

func (pr *PlayerRepo) RetiredById(id string) (*models.RetiredPlayer, bool) {
	var p RetiredPlayerDto
	trx := pr.g.Model(&RetiredPlayerDto{}).Find(&p, "Id = ?", id)
	if trx.RowsAffected != 1 {
		return nil, false
	}

	return p.RetiredPlayer(), true
}

func (pr *PlayerRepo) ById(id string) (*models.PlayerDetailed, bool) {
	var p PlayerDto
	trx := pr.g.Model(&PlayerDto{}).Preload(teamRel).Preload("History").Find(&p, "Id = ?", id)
	if trx.RowsAffected != 1 {
		return nil, false
	}

	return p.PlayerDetailed(), true
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
