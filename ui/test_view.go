package ui

import (
	"fdsim/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func TestView(ctx *AppContext) *fyne.Container {
	return container.NewCenter(
		container.NewVBox(
			widgets.Icon("dumbell"),
			widgets.Icon("transfers"),
		),
	)
}
