package ui

import (
	"fdsim/models"
	"fdsim/vm"
	"fdsim/widgets"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func playerDetailsView(ctx *AppContext) *fyne.Container {
	id := ctx.RouteParam.(string)
	g, isGameInit := ctx.GetGameState()
	player := ctx.Db.PlayerR().ById(id)
	canSeeDetails := false
	isManagedPlayer := false
	if player.Team != nil {
		canSeeDetails = IsFDTeam(player.Team.Id)
		// if I add scouting this can be different
		isManagedPlayer = canSeeDetails
	}

	showStats := isGameInit
	moraleInfo := valueLabel("Morale:", widgets.Icon("unknown"))
	if canSeeDetails {
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
		moraleInfo = valueLabel("Morale:", moraleIcon)
	}

	skillInfo := centered(starsFromPerc(player.Skill))
	if canSeeDetails {
		skillInfo = centered(widget.NewLabel(player.Skill.String()))
	}
	var main fyne.CanvasObject
	main = makePlayerMainDetailsView(player, moraleInfo, skillInfo, showStats, ctx, g)
	if isManagedPlayer {
		main = container.NewAppTabs(
			container.NewTabItemWithIcon("Info", theme.AccountIcon(), main),
			container.NewTabItemWithIcon("Manage", theme.DocumentIcon(), centered(widget.NewLabel("Manage"))),
		)
	}

	return NewFborder().
		Top(
			NewFborder().
				Left(backButton(ctx)).
				Get(centered(h1(player.String()))),
		).
		Get(main)
}

func makePlayerMainDetailsView(player *models.PlayerWithTeam, moraleInfo *fyne.Container, skillInfo *fyne.Container, showStats bool, ctx *AppContext, g *models.Game) *fyne.Container {
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
					moraleInfo,
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
						skillInfo,
					),
				),
			),
		))

	if showStats {
		stats := ctx.Db.LeagueR().GetStatsForPlayer(player.Id, g.LeagueId)
		score := 0.0
		if stats.Played > 0 {
			score = stats.Score / float64(stats.Played)
		}
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
					centered(widget.NewLabel(fmt.Sprintf("%.1f", score))),
				),
			))
		main.AddObject(statsWrapper)
	}
	return main
}
