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
	g, isGameInit := ctx.GetGameState()
	player := ctx.Db.PlayerR().ById(id)

	showStats := isGameInit

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

	main := container.NewVBox(
		container.NewGridWithColumns(3,
			centered(widgets.FlagIcon(player.Country)),
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
	)

	if showStats {
		stats := ctx.Db.LeagueR().GetStatsForPlayer(player.Id, g.LeagueId)
		statsWrapper := widget.NewCard(
			"", "Season stats",
			container.NewVBox(
				container.NewGridWithColumns(2,
					centered(widget.NewLabel("Played")),
					centered(widget.NewLabel(fmt.Sprintf("%d", stats.Played))),
				),
				container.NewGridWithColumns(2,
					centered(widget.NewLabel("Goals")),
					centered(widget.NewLabel(fmt.Sprintf("%d", stats.Goals))),
				),
				container.NewGridWithColumns(2,
					centered(widget.NewLabel("Score")),
					centered(widget.NewLabel(fmt.Sprintf("%.2f", stats.Score))),
				),
			))

		main.AddObject(widget.NewSeparator())
		main.AddObject(statsWrapper)
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
		Get(main)
}
