package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func dashboardView(ctx *AppContext) *fyne.Container {
	// gamestate
	gameId := ctx.RouteParam.(string)
	game := ctx.Db.GameR().ById(gameId)
	fd := game.FootDirector()
	// store name,surname current LeagueId
	return container.NewCenter(
		container.NewVBox(
			widget.NewLabel("Dashboard"),
			widget.NewLabel(fmt.Sprintf("%s %s (%d)", fd.Fame, fd.Surname, fd.Age)),
			starsFromPerc(fd.Fame),
			widget.NewButton("League", func() {
				// ctx.PushWithParam(League, )
				ctx.Push(League)
			}),
		),
	)
}
