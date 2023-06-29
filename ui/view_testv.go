package ui

import (
	"fdsim/enums"
	"fdsim/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func testView(ctx *AppContext) *fyne.Container {
	return NewFborder().
		Top(leftAligned(backButton(ctx))).
		Get(
			container.NewCenter(
				container.NewVBox(
					container.NewHBox(
						widgets.FlagIcon(enums.IT),
						widgets.FlagIcon(enums.EN),
						widgets.FlagIcon(enums.FR),
						widgets.FlagIcon(enums.ES),
						widgets.FlagIcon(enums.DE),
					),
					container.NewHBox(
						widgets.Icon("newspaper"),
						widgets.Icon("newspaper_read"),
					),
				),
			),
		)
}
