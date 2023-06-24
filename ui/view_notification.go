package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func notificationView(ctx *AppContext, route AppRoute) *fyne.Container {
	id := ctx.RouteParam.(string)
	return NewFborder().
		Top(NewFborder().Left(backButton(ctx)).Get(centered(h1(route.String())))).
		Get(
			container.NewCenter(
				widget.NewLabel(
					fmt.Sprintf("Notification %s", id),
				),
			),
		)
}
