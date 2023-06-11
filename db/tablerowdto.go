package db

import "fdsim/models"

type TableRowDto struct {
	TeamId       string
	Played       int
	Wins         int
	Draws        int
	Losses       int
	Points       int
	GoalScored   int
	GoalConceded int

	LeagueId string
}

func DtoFromTableRow(tr *models.Row, leagueId string) TableRowDto {
	return TableRowDto{
		TeamId:       tr.Team.Id,
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
