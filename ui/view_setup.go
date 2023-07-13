package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func setupView(ctx *AppContext) *fyne.Container {
	return container.NewBorder(
		centered(widget.NewLabel("Setup")),
		leftAligned(
			topNavBar(ctx),
		),
		nil,
		nil,
	)
}
