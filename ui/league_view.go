package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func leagueView(ctx *AppContext) *fyne.Container {

	return NewFborder().
		Top(NewFborder().Left(backButton(ctx)).Get()).
		Get(
			widget.NewLabel("League"),
		)
}
