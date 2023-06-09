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
				widgets.Flag(enums.IT),
				widgets.Flag(enums.EN),
				widgets.Flag(enums.FR),
				widgets.Flag(enums.ES),
				widgets.Flag(enums.DE),
			),
			widgets.Icon("sdsadsa"),
		),
	)
}
