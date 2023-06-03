package db

import (
	"fdsim/enums"
	"fdsim/models"
	"fdsim/utils"
)

type PlayerDto struct {
	Id      string `gorm:"primarykey;size:16"`
	Name    string
	Surname string
	Country enums.Country
	Age     int
	Role    models.Role

	Skill  int
	Morale int
	Fame   int

	TeamId string
}

func DtoFromPlayer(player *models.Player, teamId string) PlayerDto {
	return PlayerDto{
		Id:      player.Id,
		Name:    player.Name,
		Surname: player.Surname,
		Age:     player.Age,
		Country: player.Country,
		Role:    player.Role,

		Skill:  player.Skill.Val(),
		Morale: player.Morale.Val(),
		Fame:   player.Fame.Val(),

		TeamId: teamId,
	}
}

func (p PlayerDto) Player() *models.Player {
	player := &models.Player{
		Role: p.Role,
	}

	player.Id = p.Id
	player.Name = p.Name
	player.Surname = p.Name
	player.Age = p.Age
	player.Country = p.Country
	player.Skill = utils.NewPerc(p.Skill)
	player.Morale = utils.NewPerc(p.Morale)
	player.Fame = utils.NewPerc(p.Fame)

	return player
}
