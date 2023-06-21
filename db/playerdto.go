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

	Value     int64
	IdealWage int64

	TeamId    *string
	Wage      int64
	YContract uint8

	Team *TeamDto `gorm:"foreignKey:team_id"`
}

func DtoFromPlayer(player *models.Player) PlayerDto {
	return PlayerDto{
		Id:      player.Id,
		Name:    player.Name,
		Surname: player.Surname,
		Age:     player.Age,
		Country: player.Country,
		Role:    player.Role,

		Skill:     player.Skill.Val(),
		Morale:    player.Morale.Val(),
		Fame:      player.Fame.Val(),
		IdealWage: player.IdealWage.Val,
		Value:     player.Value.Val,

		TeamId:    nil,
		Wage:      player.Wage.Val,
		YContract: player.YContract,
	}
}

func DtoFromPlayerWithTeam(player *models.Player, teamId string) PlayerDto {
	p := DtoFromPlayer(player)
	p.TeamId = &teamId

	return p
}

func (p PlayerDto) PlayerPH() *models.PNPH {
	return &models.PNPH{
		Id:      p.Id,
		Name:    p.Name,
		Surname: p.Surname,
	}
}

func (p PlayerDto) Player() *models.Player {
	player := &models.Player{
		Role: p.Role,
	}

	player.Id = p.Id
	player.Name = p.Name
	player.Surname = p.Surname
	player.Age = p.Age
	player.Country = p.Country
	player.Skill = utils.NewPerc(p.Skill)
	player.Morale = utils.NewPerc(p.Morale)
	player.Fame = utils.NewPerc(p.Fame)
	player.Value = toMoney(p.Value)
	player.IdealWage = toMoney(p.IdealWage)
	player.Wage = toMoney(p.Wage)
	player.YContract = p.YContract

	return player
}

func (p PlayerDto) PlayerWithTeam() *models.PlayerWithTeam {
	player := models.Player{
		Role: p.Role,
	}

	player.Id = p.Id
	player.Name = p.Name
	player.Surname = p.Surname
	player.Age = p.Age
	player.Country = p.Country
	player.Skill = utils.NewPerc(p.Skill)
	player.Morale = utils.NewPerc(p.Morale)
	player.Fame = utils.NewPerc(p.Fame)
	player.Value = toMoney(p.Value)
	player.IdealWage = toMoney(p.IdealWage)
	player.Wage = toMoney(p.Wage)
	player.YContract = p.YContract

	var team *models.TPH = nil
	if p.Team != nil {
		team = p.Team.TeamPH()
	}

	return &models.PlayerWithTeam{
		Player: player,
		Team:   team,
	}
}
