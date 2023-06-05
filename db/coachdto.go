package db

import (
	"fdsim/enums"
	"fdsim/models"
	"fdsim/utils"

	"gorm.io/gorm"
)

type CoachDto struct {
	Id      string `gorm:"primarykey;size:16"`
	Name    string
	Surname string
	Country enums.Country
	Age     int
	Module  models.Module

	Skill  int
	Morale int
	Fame   int

	TeamId *string
}

func DtoFromCoach(c *models.Coach) CoachDto {
	return CoachDto{
		Id:      c.Id,
		Name:    c.Name,
		Surname: c.Surname,
		Age:     c.Age,
		Country: c.Country,

		Skill:  c.Skill.Val(),
		Morale: c.Morale.Val(),
		Fame:   c.Fame.Val(),

		Module: c.Module,

		TeamId: nil,
	}
}
func DtoFromCoachWithTeam(c *models.Coach, teamId string) CoachDto {
	coach := DtoFromCoach(c)
	coach.TeamId = &teamId
	return coach
}

func (c CoachDto) Coach() *models.Coach {
	coach := &models.Coach{
		Module: c.Module,
	}

	coach.Id = c.Id
	coach.Name = c.Name
	coach.Surname = c.Surname
	coach.Age = c.Age
	coach.Country = c.Country
	coach.Skill = utils.NewPerc(c.Skill)
	coach.Morale = utils.NewPerc(c.Morale)
	coach.Fame = utils.NewPerc(c.Fame)

	return coach
}

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
