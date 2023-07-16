package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func notFoundView(ctx *AppContext, objectName string) *fyne.Container {
	return NewFborder().
		Top(leftAligned(topNavBar(ctx))).
		Get(
			centered(
				h1(fmt.Sprintf("%s Not Found.", objectName)),
			),
		)
}
