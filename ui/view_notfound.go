package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func notFoundView(ctx *AppContext, objectName string) *fyne.Container {
	return NewFborder().
		Top(leftAligned(topNavBar(ctx))).
		Get(
			centered(
				container.NewVBox(
					widget.NewIcon(theme.ErrorIcon()),
					h1(fmt.Sprintf("%s Not Found.", objectName)),
				),
			),
		)
}
