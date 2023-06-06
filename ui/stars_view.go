package ui

import (
	"fdsim/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func stars(value utils.Perc) *fyne.Container {
	s := make([]fyne.CanvasObject, 5)

	for i := range s {
		img := canvas.NewImageFromResource(theme.AccountIcon())
		s[i] = container.NewMax(img)
	}
	return container.NewAdaptiveGrid(5, s...)
}
