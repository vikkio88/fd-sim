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

	main := container.NewGridWithRows(2,
		container.NewGridWithColumns(2,
			widget.NewCard("", "Team Info",
				container.NewVBox(
					centered(widget.NewLabel(player.Team.Name)),
					valueLabel("Fame:",
						centered(starsFromPerc(player.Fame)),
					),
					valueLabel("Value:",
						centered(widget.NewLabel(player.Value.StringKMB())),
					),
					valueLabel("Contract:",
						widget.NewLabel(fmt.Sprintf("%s / %d years", player.Wage.StringKMB(), player.YContract)),
					),
					valueLabel("Morale:",
						moraleIcon,
					),
				),
			),
			widget.NewCard("", "Personal Info",
				container.NewVBox(
					valueLabel("Age:",
						centered(widget.NewLabel(fmt.Sprintf("%d", player.Age))),
					),
					valueLabel("Role:",
						centered(widget.NewLabel(player.Role.String())),
					),
					container.NewGridWithColumns(3,
						centered(boldLabel("Nationality:")),
						centered(widgets.FlagIcon(player.Country)),
						centered(widget.NewLabel(fmt.Sprintf("(%s)", player.Country.Nationality()))),
					),
					valueLabel("Skill:",
						centered(starsFromPerc(player.Skill)),
					),
				),
			),
		))

	if showStats {
		stats := ctx.Db.LeagueR().GetStatsForPlayer(player.Id, g.LeagueId)
		statsWrapper := widget.NewCard(
			"", "Season stats",
			container.NewVBox(
				valueLabel("Played:",
					centered(widget.NewLabel(fmt.Sprintf("%d", stats.Played))),
				),
				valueLabel("Goals:",
					centered(widget.NewLabel(fmt.Sprintf("%d", stats.Goals))),
				),
				valueLabel("Score:",
					centered(widget.NewLabel(fmt.Sprintf("%.2f", stats.Score))),
				),
			))
		main.AddObject(statsWrapper)
	}

	return NewFborder().
		Top(
			NewFborder().
				Left(backButton(ctx)).
				Get(centered(h1(player.String()))),
		).
		Get(main)
}
