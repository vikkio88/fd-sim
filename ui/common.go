package ui

import (
	"fdsim/utils"
	vm "fdsim/vm"
	"fdsim/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func h1(text string) *canvas.Text {
	txt := canvas.NewText(text, theme.ForegroundColor())
	txt.TextSize = 20
	txt.Alignment = fyne.TextAlignCenter
	return txt
}

func h2(text string) *canvas.Text {
	txt := canvas.NewText(text, theme.ForegroundColor())
	txt.TextSize = 18
	return txt
}

func small(text string) *canvas.Text {
	txt := canvas.NewText(text, theme.ForegroundColor())
	txt.TextSize = 10
	return txt
}

func centered(object fyne.CanvasObject) *fyne.Container {
	return container.NewCenter(object)
}
func rightAligned(object fyne.CanvasObject) *fyne.Container {
	return container.NewBorder(nil, nil, nil, object)
}

func leftAligned(object fyne.CanvasObject) *fyne.Container {
	return container.NewBorder(nil, nil, object, nil)
}

func topAligned(object fyne.CanvasObject) *fyne.Container {
	return container.NewBorder(object, nil, nil, nil)
}

func bottomAligned(object fyne.CanvasObject) *fyne.Container {
	return container.NewBorder(nil, object, nil, nil)
}

func backButton(ctx *AppContext) *widget.Button {
	return widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		ctx.Pop()
	})
}

func starsFromPerc(perc utils.Perc) fyne.CanvasObject {
	return widgets.NewStarRatingFromFloat(vm.PercToStars(perc))
}

func starsFromf64(value float64) fyne.CanvasObject {
	return widgets.NewStarRatingFromFloat(vm.PercFToStars(value))
}

func valueLabel(label string, value fyne.CanvasObject) *fyne.Container {
	labelLbl := boldLabel(label)
	return container.NewGridWithColumns(2,
		centered(labelLbl),
		value,
	)
}
func boldLabel(label string) *widget.Label {
	labelLbl := widget.NewLabel(label)
	labelLbl.TextStyle = fyne.TextStyle{Bold: true}

	return labelLbl
}
