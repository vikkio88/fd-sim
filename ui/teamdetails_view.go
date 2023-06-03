package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func teamDetailsView(ctx *AppContext) *fyne.Container {
	id := ctx.RouteParam.(string)
	return NewFborder().
		Top(widget.NewLabel(id)).
		Bottom(leftAligned(backButton(ctx))).
		Get()
}
