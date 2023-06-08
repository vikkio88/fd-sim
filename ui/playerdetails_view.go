package ui

import (
	"fdsim/vm"
	"fdsim/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func playerDetailsView(ctx *AppContext) *fyne.Container {
	id := ctx.RouteParam.(string)
	player := ctx.Db.PlayerR().ById(id)

	morale := vm.MoraleEmojFromPerc(player.Morale)
	var moraleIcon *widget.Icon
	switch morale {
	case vm.Happy:
		moraleIcon = widgets.Icon("happy_face")
	case vm.Meh:
		moraleIcon = widgets.Icon("meh_face")
	case vm.Sad:
		moraleIcon = widgets.Icon("sad_face")
	}

	return NewFborder().
		Top(
			NewFborder().Left(backButton(ctx)).
				Get(
					centered(
						container.NewHBox(
							h1(player.String()),
							small(player.Country.String()),
						),
					),
				),
		).
		Get(
			container.NewVBox(
				container.NewGridWithColumns(2,
					widget.NewLabel("Morale"),
					moraleIcon,
				),
			),
		)
}
