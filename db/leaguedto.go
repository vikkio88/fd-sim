package db

import (
	"fdsim/enums"
	"fdsim/models"
)

type LeagueDto struct {
	Id      string `gorm:"primarykey;size:16"`
	Name    string
	Country enums.Country

	Teams     []TeamDto     `gorm:"foreignKey:league_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Rounds    []RoundDto    `gorm:"foreignKey:league_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TableRows []TableRowDto `gorm:"foreignKey:league_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	RPointer int
}

func DtoFromLeagueEmpty(l *models.League) LeagueDto {
	ldto := LeagueDto{
		Id:       l.Id,
		Name:     l.Name,
		Country:  l.Country,
		RPointer: l.RPointer,

		Teams:     []TeamDto{},
		Rounds:    []RoundDto{},
		TableRows: []TableRowDto{},
	}

	return ldto
}

func DtoFromLeagueNoTeams(l *models.League) LeagueDto {
	trs := DtoFromTableRows(l.Table.Rows(), l.Id)

	rds := make([]RoundDto, len(l.Rounds))
	for i, r := range l.Rounds {
		rds[i] = DtoFromRoundPH(r, l.Id)
	}

	ldto := LeagueDto{
		Id:        l.Id,
		Name:      l.Name,
		Country:   l.Country,
		RPointer:  l.RPointer,
		Rounds:    rds,
		TableRows: trs,

		Teams: []TeamDto{},
	}

	return ldto
}

func DtoFromLeague(l *models.League) LeagueDto {
	trs := DtoFromTableRows(l.Table.Rows(), l.Id)

	rds := make([]RoundDto, len(l.Rounds))
	for i, r := range l.Rounds {
		rds[i] = DtoFromRoundPH(r, l.Id)
	}

	ldto := LeagueDto{
		Id:       l.Id,
		Name:     l.Name,
		Country:  l.Country,
		RPointer: l.RPointer,

		Teams: DtoFromTeams(l.Teams(), l.Id),

		TableRows: trs,
		Rounds:    rds,
	}

	return ldto
}

func (l *LeagueDto) GetTeams() []*models.Team {
	ts := make([]*models.Team, len(l.Teams))

	for i, tdto := range l.Teams {
		ts[i] = tdto.Team()
	}
	return ts
}

func (l *LeagueDto) League() *models.League {
	league := models.NewLeagueWithData(l.Id, l.Name, l.GetTeams())

	league.RPointer = l.RPointer
	league.Table = TableFromTableRowsDto(l.TableRows)
	league.Rounds = RoundsPHFromDto(l.Rounds)

	return league
}
