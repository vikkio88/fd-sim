package ui

import (
	"fdsim/enums"
	"fdsim/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func TestView(ctx *AppContext) *fyne.Container {
	return container.NewCenter(
		container.NewVBox(
			container.NewHBox(
				widgets.FlagIcon(enums.IT),
				widgets.FlagIcon(enums.EN),
				widgets.FlagIcon(enums.FR),
				widgets.FlagIcon(enums.ES),
				widgets.FlagIcon(enums.DE),
			),
			widgets.Icon("contract"),
		),
	)
}
