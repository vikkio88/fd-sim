package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func teamDetailsView(ctx *AppContext) *fyne.Container {
	id := ctx.RouteParam.(string)
	team := ctx.Db.TeamR().ById(id)
	return NewFborder().
		Top(widget.NewLabel(team.String())).
		Bottom(leftAligned(backButton(ctx))).
		Get()
}
