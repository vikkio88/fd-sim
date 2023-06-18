package ui

import (
	"fdsim/conf"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func loadGameView(ctx *AppContext) *fyne.Container {
	games := ctx.Db.GameR().All()
	pHolder := widget.NewLabel("No Save Games...")
	gamesList := widget.NewList(
		func() int {
			return len(games)
		},
		func() fyne.CanvasObject {
			return widget.NewCard("", "", NewFborder().
				Right(widget.NewButtonWithIcon("Load", theme.LoginIcon(), func() {})).
				Left(widget.NewLabel("Date")).
				Get(centered(widget.NewLabel("SaveGame"))))
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			g := games[lii]
			card := co.(*widget.Card)
			container := card.Content.(*fyne.Container)
			container.Objects[0].(*fyne.Container).Objects[0].(*widget.Label).SetText(fmt.Sprintf("Player: %s", g.SaveName))
			container.Objects[1].(*widget.Label).SetText(g.Date.Format(conf.GameDateFormat))
			container.Objects[2].(*widget.Button).OnTapped = func() {
				ctx.NavigateToWithParam(Dashboard, g.Id)
			}
		},
	)
	gamesList.Hide()

	if len(games) > 0 {
		pHolder.Hide()
		gamesList.Show()
	}
	return NewFborder().
		Top(centered(h1("Load Game"))).
		Bottom(leftAligned(backButton(ctx))).
		Get(
			container.NewCenter(pHolder),
			gamesList,
		)
}
