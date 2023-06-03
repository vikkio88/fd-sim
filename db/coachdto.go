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

	TeamId string
}

func DtoFromCoach(c *models.Coach, teamId string) CoachDto {
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

		TeamId: teamId,
	}
}

func (c CoachDto) Coach() *models.Coach {
	coach := &models.Coach{
		Module: c.Module,
	}

	coach.Id = c.Id
	coach.Name = c.Name
	coach.Surname = c.Name
	coach.Age = c.Age
	coach.Country = c.Country
	coach.Skill = utils.NewPerc(c.Skill)
	coach.Morale = utils.NewPerc(c.Morale)
	coach.Fame = utils.NewPerc(c.Fame)

	return coach
}
