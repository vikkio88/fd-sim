package db

import "fdsim/models"

type TableRowIndexDto struct {
	TeamId       string
	Played       int
	Wins         int
	Draws        int
	Losses       int
	Points       int
	GoalScored   int
	GoalConceded int
	// Index is a reserved word
	Position int
}

func (tr *TableRowIndexDto) TPHRow() *models.TPHRow {
	return &models.TPHRow{
		Team:  models.TPH{Id: tr.TeamId},
		Index: tr.Position,
		Row: &models.Row{
			Team:         tr.TeamId,
			Played:       tr.Played,
			Wins:         tr.Wins,
			Draws:        tr.Draws,
			Losses:       tr.Losses,
			Points:       tr.Points,
			GoalScored:   tr.GoalScored,
			GoalConceded: tr.GoalConceded,
		},
	}
}

type TableRowDto struct {
	TeamId       string `gorm:"primarykey;size:16"`
	Played       int
	Wins         int
	Draws        int
	Losses       int
	Points       int
	GoalScored   int
	GoalConceded int

	LeagueId string `gorm:"primarykey;size:16"`
}

func DtoFromTableRows(rs []*models.Row, leagueId string) []TableRowDto {
	result := make([]TableRowDto, len(rs))
	for i, tr := range rs {
		result[i] = DtoFromTableRow(tr, leagueId)
	}

	return result
}
func DtoFromTableRow(tr *models.Row, leagueId string) TableRowDto {
	return TableRowDto{
		TeamId:       tr.Team,
		Played:       tr.Played,
		Wins:         tr.Wins,
		Draws:        tr.Draws,
		Losses:       tr.Losses,
		Points:       tr.Points,
		GoalScored:   tr.GoalScored,
		GoalConceded: tr.GoalConceded,

		LeagueId: leagueId,
	}
}

func TableFromTableRowsDto(trs []TableRowDto) *models.Table {
	rs := make([]*models.Row, len(trs))
	for i, tr := range trs {
		rs[i] = &models.Row{
			Team:         tr.TeamId,
			Played:       tr.Played,
			Wins:         tr.Wins,
			Draws:        tr.Draws,
			Losses:       tr.Losses,
			Points:       tr.Points,
			GoalScored:   tr.GoalScored,
			GoalConceded: tr.GoalConceded,
		}
	}

	return models.NewTableFromRows(rs)
}
