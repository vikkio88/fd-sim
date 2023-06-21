package db

import (
	"fdsim/enums"
	"fdsim/models"
	"fdsim/utils"
)

type TeamDto struct {
	Id            string `gorm:"primarykey;size:16"`
	Name          string
	City          string
	Country       enums.Country
	Balance       float64
	TransferRatio float64

	LeagueId *string
	CoachId  *string
	Coach    CoachDto    `gorm:"foreignKey:team_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Players  []PlayerDto `gorm:"foreignKey:team_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func DtoFromTeams(teams []*models.Team, leagueId string) []TeamDto {
	tdtos := make([]TeamDto, len(teams))
	for i, t := range teams {
		tdto := DtoFromTeam(t)
		tdto.LeagueId = &leagueId
		tdtos[i] = tdto
	}

	return tdtos
}

func DtoFromTeam(team *models.Team) TeamDto {
	ps := team.Roster.Players()
	pdtos := make([]PlayerDto, len(ps))
	for i, p := range ps {
		pdtos[i] = DtoFromPlayerWithTeam(p, team.Id)
	}
	return TeamDto{
		Id:            team.Id,
		Name:          team.Name,
		City:          team.City,
		Country:       team.Country,
		Balance:       team.Balance.Value(),
		TransferRatio: team.TransferRatio,
		CoachId:       &team.Coach.Id,

		Coach: DtoFromCoachWithTeam(team.Coach, team.Id),

		Players: pdtos,
	}
}

func (t TeamDto) TeamPH() *models.TPH {
	return &models.TPH{
		Id:   t.Id,
		Name: t.Name,
	}
}

func (t TeamDto) Team() *models.Team {
	ts := &models.Team{
		Name:          t.Name,
		City:          t.City,
		Balance:       utils.NewEurosFromF(t.Balance),
		TransferRatio: t.TransferRatio,
		Country:       t.Country,
		Roster:        models.NewRoster(),
	}
	ts.Id = t.Id
	ts.Coach = t.Coach.Coach()

	for _, p := range t.Players {
		ts.Roster.AddPlayer(p.Player())
	}
	return ts
}
