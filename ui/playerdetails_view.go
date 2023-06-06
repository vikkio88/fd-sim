package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func playerDetailsView(ctx *AppContext) *fyne.Container {
	id := ctx.RouteParam.(string)
	player := ctx.Db.PlayerR().ById(id)

	return NewFborder().
		Top(
			NewFborder().Left(backButton(ctx)).
				Get(
					centered(
						container.NewHBox(
							widget.NewLabel(player.String()),
							small(player.Country.String()),
						),
					),
				)).Get()
}
