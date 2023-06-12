package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func dashboardView(ctx *AppContext) *fyne.Container {
	// gamestate
	// store name,surname current LeagueId
	return container.NewCenter(
		container.NewVBox(
			widget.NewLabel("Dashboard"),
			widget.NewButton("League", func() {
				// ctx.PushWithParam(League, )
				ctx.Push(League)
			}),
		),
	)
}
