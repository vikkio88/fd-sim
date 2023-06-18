package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func roundView(ctx *AppContext) *fyne.Container {
	return container.NewCenter(
		widget.NewLabel("Round"),
	)
}
