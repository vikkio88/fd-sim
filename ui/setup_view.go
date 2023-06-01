package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SetupView(ctx *AppContext) *fyne.Container {
	return container.NewBorder(
		centered(widget.NewLabel("Setup")),
		leftAligned(
			backButton(ctx),
		),
		nil,
		nil,
	)
}
