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
	YContract int

	Team    *TeamDto     `gorm:"foreignKey:team_id"`
	History *PHistoryDto `gorm:"foreignKey:player_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Offers  []OfferDto   `gorm:"foreignKey:player_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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

func (p PlayerDto) PlayerDetailedNoAwards() *models.PlayerDetailed {
	return p.PlayerDetailed([]LHistoryDto{}, []TrophyDto{})
}

func (p PlayerDto) PlayerDetailed(awardsRows []LHistoryDto, trophiesRows []TrophyDto) *models.PlayerDetailed {
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

	history := []*models.PHistoryRow{}
	if p.History != nil {
		history = p.History.HistoryRows()
	}

	awards := []models.Award{}
	if len(awardsRows) > 0 {
		for _, s := range awardsRows {
			awards = append(awards, s.Award(p.Id))
		}
	}

	trophies := []models.Trophy{}
	if len(trophiesRows) > 0 {
		for _, t := range trophiesRows {
			trophies = append(trophies, t.Trophy())
		}
	}

	return &models.PlayerDetailed{
		Player:   player,
		History:  history,
		Team:     team,
		Awards:   awards,
		Trophies: trophies,
	}
}
