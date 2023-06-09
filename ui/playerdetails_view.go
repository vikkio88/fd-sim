package ui

import (
	"fdsim/vm"
	"fdsim/widgets"
	"fmt"

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
							small(fmt.Sprintf("%d", player.Age)),
						),
					),
				),
		).
		Get(
			container.NewVBox(
				container.NewGridWithColumns(3,
					centered(widget.NewLabel(player.Country.Nationality())),
					centered(widget.NewLabel(player.Role.String())),
					centered(starsFromPerc(player.Skill)),
				),
				container.NewGridWithColumns(2,
					centered(widget.NewLabel("Fame")),
					centered(starsFromPerc(player.Fame)),
				),
				container.NewGridWithColumns(2,
					centered(widget.NewLabel("Value")),
					centered(widget.NewLabel(player.Value.StringKMB())),
				),
				container.NewGridWithColumns(2,
					centered(widget.NewLabel("Contract")),
					widget.NewLabel(fmt.Sprintf("%s / %d years", player.Wage.StringKMB(), player.YContract)),
				),
				container.NewGridWithColumns(2,
					centered(widget.NewLabel("Morale")),
					moraleIcon,
				),
			),
		)
}
