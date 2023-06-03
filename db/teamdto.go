package db

import (
	"fdsim/enums"
	"fdsim/models"
)

type TeamDto struct {
	Id      string `gorm:"primarykey;size:16"`
	Name    string
	City    string
	Country enums.Country

	Coach   CoachDto    `gorm:"foreignKey:TeamId"`
	Players []PlayerDto `gorm:"foreignKey:TeamId"`
}

func DtoFromTeam(team *models.Team) TeamDto {
	ps := team.Roster.Players()
	pdtos := make([]PlayerDto, len(ps))
	for i, p := range ps {
		pdtos[i] = DtoFromPlayer(p, team.Id)
	}
	return TeamDto{
		Id:      team.Id,
		Name:    team.Name,
		City:    team.City,
		Country: team.Country,

		Coach: DtoFromCoach(team.Coach, team.Id),

		Players: pdtos,
	}
}

func (t TeamDto) Team() *models.Team {
	ts := &models.Team{
		Name:    t.Name,
		City:    t.City,
		Country: t.Country,
		Roster:  models.NewRoster(),
	}
	ts.Id = t.Id
	ts.Coach = t.Coach.Coach()

	for _, p := range t.Players {
		ts.Roster.AddPlayer(p.Player())
	}
	return ts
}
