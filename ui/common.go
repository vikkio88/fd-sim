package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func h1(text string) *canvas.Text {
	txt := canvas.NewText(text, theme.ForegroundColor())
	txt.TextSize = 20
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