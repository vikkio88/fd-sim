package db

import (
	"fdsim/enums"
	"fdsim/models"

	"gorm.io/gorm"
)

type TeamDto struct {
	Id      string `gorm:"primarykey;size:16"`
	Name    string
	City    string
	Country enums.Country

	//TODO: test ONDelete constraint
	Coach   CoachDto    `gorm:"foreignKey:team_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Players []PlayerDto `gorm:"foreignKey:team_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func DtoFromTeam(team *models.Team) TeamDto {
	ps := team.Roster.Players()
	pdtos := make([]PlayerDto, len(ps))
	for i, p := range ps {
		pdtos[i] = DtoFromPlayerWithTeam(p, team.Id)
	}
	return TeamDto{
		Id:      team.Id,
		Name:    team.Name,
		City:    team.City,
		Country: team.Country,

		Coach: DtoFromCoachWithTeam(team.Coach, team.Id),

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

type TeamsRepo struct {
	g *gorm.DB
}

func NewTeamsRepo(g *gorm.DB) *TeamsRepo {
	return &TeamsRepo{
		g,
	}
}

func (tr *TeamsRepo) InsertOne(t *models.Team) {
	tdto := DtoFromTeam(t)

	tr.g.Create(&tdto)
}

func (tr *TeamsRepo) Insert(teams []*models.Team) {
	tdtos := make([]TeamDto, len(teams))
	for i, t := range teams {
		tdtos[i] = DtoFromTeam(t)
	}

	tr.g.Create(&tdtos)
}

func (tr *TeamsRepo) ById(id string) *models.Team {
	var t TeamDto
	tr.g.Model(&TeamDto{}).Preload("Players").Preload("Coach").Find(&t, "Id = ?", id)

	return t.Team()
}

func (tr *TeamsRepo) DeleteOne(id string) {
	tr.g.Where("Id = ?", id).Delete(&TeamDto{})
}

func (tr *TeamsRepo) Delete(ids []string) {
	tr.g.Delete(&TeamDto{}, ids)
}

func (tr *TeamsRepo) Count() int64 {
	var c int64
	tr.g.Model(&TeamDto{}).Count(&c)

	return c
}

func (tr *TeamsRepo) All() []*models.Team {
	var tdtos []TeamDto
	tr.g.Model(&TeamDto{}).Preload("Players").Find(&tdtos)

	ts := make([]*models.Team, len(tdtos))
	for i, t := range tdtos {
		ts[i] = t.Team()
	}
	return ts
}
