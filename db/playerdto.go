package db

import (
	"fdsim/enums"
	"fdsim/models"
	"fdsim/utils"

	"gorm.io/gorm"
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

	TeamId *string
}

func DtoFromPlayer(player *models.Player) PlayerDto {
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

		TeamId: nil,
	}
}

func DtoFromPlayerWithTeam(player *models.Player, teamId string) PlayerDto {
	p := DtoFromPlayer(player)
	p.TeamId = &teamId

	return p
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

	return player
}

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

func (pr *PlayerRepo) ById(id string) *models.Player {
	var p PlayerDto
	pr.g.Model(&PlayerDto{}).Find(&p, "Id = ?", id)

	return p.Player()
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
