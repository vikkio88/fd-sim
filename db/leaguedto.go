package db

import "fdsim/models"

type LeagueDto struct {
	Id   string
	Name string

	Teams     []TeamDto     `gorm:"foreignKey:league_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Rounds    []RoundDto    `gorm:"foreignKey:league_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TableRows []TableRowDto `gorm:"foreignKey:league_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	RPointer int
}

func DtoFromLeague(l *models.League) LeagueDto {

	trs := make([]TableRowDto, len(l.Table.Rows()))
	for i, tr := range l.Table.Rows() {
		trs[i] = DtoFromTableRow(tr, l.Id)
	}

	rds := make([]RoundDto, len(l.Rounds))
	for i, r := range l.Rounds {
		rds[i] = DtoFromRoundPH(r, l.Id)
	}

	ldto := LeagueDto{
		Id:       l.Id,
		Name:     l.Name,
		RPointer: l.RPointer,
		// TEAMS WILL BE UPDATED/ADDED BY THE REPO

		TableRows: trs,
		Rounds:    rds,
	}

	return ldto
}
