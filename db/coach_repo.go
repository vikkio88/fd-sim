package db

import (
	"fdsim/models"

	"gorm.io/gorm"
)

type CoachRepo struct {
	g *gorm.DB
}

func NewCoachRepo(g *gorm.DB) *CoachRepo {
	return &CoachRepo{
		g,
	}
}

func (pr *CoachRepo) InsertOne(c *models.Coach) {
	cdto := DtoFromCoach(c)

	pr.g.Create(&cdto)
}

func (cr *CoachRepo) Insert(coaches []*models.Coach) {
	cdtos := make([]CoachDto, len(coaches))
	for i, c := range coaches {
		cdtos[i] = DtoFromCoach(c)
	}

	cr.g.Create(&cdtos)
}

func (cr *CoachRepo) ById(id string) *models.Coach {
	var c CoachDto
	cr.g.Model(&CoachDto{}).Find(&c, "Id = ?", id)

	return c.Coach()
}

func (cr *CoachRepo) Truncate() {
	cr.g.Where("1 = 1").Delete(&CoachDto{})
}

func (cr *CoachRepo) DeleteOne(id string) {
	cr.g.Delete(&CoachDto{}, id)
}

func (cr *CoachRepo) Delete(ids []string) {
	cr.g.Delete(&CoachDto{}, ids)
}

func (cr *CoachRepo) Count() int64 {
	var c int64
	cr.g.Model(&CoachDto{}).Count(&c)

	return c
}

func (cr *CoachRepo) FreeAgents() []*models.Coach {
	var cdtos []CoachDto
	cr.g.Model(&CoachDto{}).Where("team_id", nil).Find(&cdtos)

	cs := make([]*models.Coach, len(cdtos))
	for i, c := range cdtos {
		cs[i] = c.Coach()
	}
	return cs
}

func (cr *CoachRepo) All() []*models.Coach {
	var cdtos []CoachDto
	cr.g.Model(&CoachDto{}).Find(&cdtos)

	cs := make([]*models.Coach, len(cdtos))
	for i, c := range cdtos {
		cs[i] = c.Coach()
	}
	return cs
}
