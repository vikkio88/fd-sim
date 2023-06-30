package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func pressView(ctx *AppContext) *fyne.Container {
	return NewFborder().
		Top(
			NewFborder().
				Left(backButton(ctx)).
				Get(centered(h1("Press"))),
		).
		Get(
			container.NewCenter(
				widget.NewLabel("Press"),
			),
		)
}
