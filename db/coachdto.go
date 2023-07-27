package db

import (
	"fdsim/enums"
	"fdsim/models"
	"fdsim/utils"
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

	IdealWage int64

	TeamId    *string
	Wage      int64
	YContract int

	RngSeed int64
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

		IdealWage: c.IdealWage.Val,

		TeamId:    nil,
		Wage:      c.Wage.Val,
		YContract: c.YContract,

		RngSeed: c.RngSeed,
	}
}
func DtoFromCoachWithTeam(c *models.Coach, teamId string) *CoachDto {
	coach := DtoFromCoach(c)
	coach.TeamId = &teamId
	return &coach
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
	coach.IdealWage = toMoney(c.IdealWage)
	coach.Wage = toMoney(c.Wage)
	coach.YContract = c.YContract
	coach.RngSeed = c.RngSeed

	return coach
}
