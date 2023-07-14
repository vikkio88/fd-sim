package ui

import (
	"fdsim/models"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func leaguehistoryView(ctx *AppContext) *fyne.Container {
	leagueId := ctx.RouteParam.(string)

	league := ctx.Db.LeagueR().HistoryById(leagueId)
	mvp := league.Mvp

	navigate := ctx.PushWithParam

	return NewFborder().
		Top(NewFborder().Left(topNavBar(ctx)).Get(centered(h1(fmt.Sprintf("%s - Historical Stats", league.Name))))).
		Get(
			container.NewGridWithRows(3,
				container.NewCenter(
					container.NewVBox(
						centered(h2("Teams")),
						makePodiumRow(league.Podium[0], "1st: ", navigate),
						makePodiumRow(league.Podium[1], "2nd: ", navigate),
						makePodiumRow(league.Podium[2], "3rd: ", navigate),
					),
				),
				container.NewCenter(
					container.NewVBox(
						centered(h2("MVP")),
						makeMvpRow(mvp, navigate),
					),
				),
				container.NewCenter(
					container.NewVBox(
						centered(h2("Scorers")),
						makeScorerRow(league.BestScorers[0], "1st", navigate),
						makeScorerRow(league.BestScorers[1], "2nd", navigate),
						makeScorerRow(league.BestScorers[2], "3rd", navigate),
					),
				),
			),
		)
}

func makeMvpRow(mvp *models.PlayerHistorical, navigate NavigateWithParamFunc) fyne.CanvasObject {
	return container.NewGridWithColumns(5,
		hL(fmt.Sprintf("%s %s", mvp.Name, mvp.Surname), func() {}),
		hL(mvp.Team.Name, func() { navigate(PlayerDetails, mvp.Id) }),
		valueLabel("Played:", widget.NewLabel(fmt.Sprintf("%d", mvp.Played))),
		valueLabel("Goals:", widget.NewLabel(fmt.Sprintf("%d", mvp.Goals))),
		valueLabel("Score:", widget.NewLabel(fmt.Sprintf("%.2f", mvp.Score/float64(mvp.Played)))),
	)
}

func makeScorerRow(scorer *models.PlayerHistorical, position string, navigate NavigateWithParamFunc) fyne.CanvasObject {
	return container.NewGridWithColumns(5,
		boldLabel(position),
		hL(fmt.Sprintf("%s %s", scorer.Name, scorer.Surname), func() {
			navigate(PlayerDetails, scorer.Id)
		}),
		hL(scorer.Team.Name, func() {}),
		valueLabel("Played:", widget.NewLabel(fmt.Sprintf("%d", scorer.Played))),
		valueLabel("Goals:", widget.NewLabel(fmt.Sprintf("%d", scorer.Goals))),
	)
}

func makePodiumRow(entry *models.TPHRow, position string, navigate NavigateWithParamFunc) fyne.CanvasObject {
	row := entry.Row
	return container.NewGridWithColumns(4,
		valueLabel(position, hL(entry.Team.Name, func() { navigate(TeamDetails, entry.Team.Id) })),
		valueLabel("Points", widget.NewLabel(fmt.Sprintf("%d", row.Points))),
		valueLabel("W/D/L", widget.NewLabel(fmt.Sprintf("%d/%d/%d", row.Wins, row.Draws, row.Losses))),
		valueLabel("GS/GC", widget.NewLabel(fmt.Sprintf("%d/%d", row.GoalScored, row.GoalConceded))),
	)
}
