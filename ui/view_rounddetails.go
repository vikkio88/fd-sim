package ui

import (
	"fdsim/conf"
	"fdsim/models"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func roundDetailsView(ctx *AppContext) *fyne.Container {
	id := ctx.RouteParam.(string)
	round := ctx.Db.LeagueR().RoundWithResults(id)
	return NewFborder().
		Top(
			NewFborder().
				Left(topNavBar(ctx)).
				Get(centered(container.NewVBox(
					h1(fmt.Sprintf("Round %d", round.Index+1)),
					h2(fmt.Sprintf("%s", round.Date.Format(conf.DateFormatGame))),
				))),
		).
		Get(
			roundMatchList(round, ctx.PushWithParam),
		)
}

func roundMatchList(
	round *models.RPHTPH,
	navigate NavigateWithParamFunc,
) fyne.CanvasObject {
	return widget.NewList(
		func() int {
			return len(round.Matches)
		},
		func() fyne.CanvasObject {
			return container.NewGridWithColumns(3,
				widget.NewLabel(""),
				centered(widget.NewHyperlink("vs", nil)),
				widget.NewLabel(""),
			)

		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			m := round.Matches[lii]
			homeLbl := co.(*fyne.Container).Objects[0].(*widget.Label)
			homeLbl.SetText(m.Home.Name)
			if IsFDTeam(m.Home.Id) {
				signalFdTeam(homeLbl)
			}

			resHL := co.(*fyne.Container).Objects[1].(*fyne.Container).Objects[0].(*widget.Hyperlink)
			resHL.SetText(m.Result.String())
			resHL.OnTapped = func() {
				navigate(MatchDetails, m.Id)
			}
			awayLbl := co.(*fyne.Container).Objects[2].(*widget.Label)
			awayLbl.SetText(m.Away.Name)
			if IsFDTeam(m.Away.Id) {
				signalFdTeam(awayLbl)
			}
		})
}
