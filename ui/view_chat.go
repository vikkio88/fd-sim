package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func chatView(ctx *AppContext) *fyne.Container {
	return NewFborder().
		Top(
			NewFborder().
				Left(topNavBar(ctx)).
				Get(centered(h1("Chat"))),
		).
		Get(
			container.NewCenter(
				widget.NewLabel("chat"),
			),
		)

}
