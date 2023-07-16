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
	player, exists := ctx.Db.PlayerR().ById(id)

	if !exists {
		retired, exists := ctx.Db.PlayerR().RetiredById(id)
		if !exists {
			return notFoundView(ctx, "Player")
		}

		return makeRetiredPlayerView(retired, ctx)
	}

	canSeeDetails := false
	isManagedPlayer := false
	if player.Team != nil {
		canSeeDetails = IsFDTeam(player.Team.Id)
		// if I add scouting this can be different
		isManagedPlayer = canSeeDetails
	}

	showStats := isGameInit
	main := container.NewAppTabs(
		container.NewTabItemWithIcon("Info", theme.AccountIcon(), makePlayerMainDetailsView(player, canSeeDetails, showStats, ctx, g)),
	)
	if isManagedPlayer {
		main.Append(
			container.NewTabItemWithIcon("Manage", theme.DocumentIcon(), centered(widget.NewLabel("Manage"))),
		)
	}

	main.Append(container.NewTabItemWithIcon("History",
		theme.DocumentIcon(),
		makePHistory(player.History, ctx.PushWithParam),
	))

	return NewFborder().
		Top(
			NewFborder().
				Left(topNavBar(ctx)).
				Get(makePlayerHeader(player)),
		).
		Get(main)
}

func makeRetiredPlayerView(retired *models.RetiredPlayer, ctx *AppContext) *fyne.Container {

	return NewFborder().
		Top(
			NewFborder().
				Left(topNavBar(ctx)).
				Get(
					centered(
						container.NewHBox(
							centered(widgets.FlagIcon(retired.Country)),
							h1(fmt.Sprintf("%s", retired.String())),
							h2(
								fmt.Sprintf(" - (%s)",
									retired.Role.StringShort(),
								),
							),
						),
					),
				),
		).
		Get(
			container.NewVBox(
				centered(
					valueLabel("Retired in:", widget.NewLabel(fmt.Sprintf("%d (%d years old)", retired.YearRetired, retired.Age))),
				),
				container.NewPadded(
					makePHistory(retired.History, ctx.PushWithParam),
				),
			),
		)
}

func makePlayerHeader(player *models.PlayerDetailed) fyne.CanvasObject {
	return centered(container.NewHBox(
		widgets.FlagIcon(player.Country),
		h1L(player.String()),
		widget.NewLabel(fmt.Sprintf("(%d) - %s", player.Age, player.Role.StringShort())),
	))
}

func makePHistory(pHistoryRow []*models.PHistoryRow, navigate NavigateWithParamFunc) fyne.CanvasObject {
	if len(pHistoryRow) < 1 {
		return centered(widget.NewLabel("No History yet"))
	}
	columns := widgets.NewColumnsLayout([]float32{-1, 100, 100, 80, 80, 80, 80})
	headers := widgets.NewListHeader(
		[]widgets.ListColumn{
			widgets.NewListCol("", fyne.TextAlignCenter),
			widgets.NewListCol("League", fyne.TextAlignLeading),
			widgets.NewListCol("Team", fyne.TextAlignCenter),
			widgets.NewListCol("Played", fyne.TextAlignLeading),
			widgets.NewListCol("Goals", fyne.TextAlignLeading),
			widgets.NewListCol("Score", fyne.TextAlignLeading),
			widgets.NewListCol("Cost", fyne.TextAlignLeading),
		},
		columns,
	)

	historyList := widget.NewList(
		func() int { return len(pHistoryRow) },
		func() fyne.CanvasObject {
			return container.New(
				columns,
				widget.NewLabel("Y"),
				centered(widget.NewHyperlink("League", nil)),
				centered(widget.NewHyperlink("Team", nil)),
				widget.NewLabel("Played"),
				widget.NewLabel("Goals"),
				widget.NewLabel("Score"),
				widget.NewLabel("Cost"),
			)
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			r := pHistoryRow[lii]
			cells := co.(*fyne.Container)

			yearLbl := cells.Objects[0].(*widget.Label)
			// TODO: if halfseson bool is true he got moved so need
			// to showcase it in here
			yearLbl.SetText(fmt.Sprintf("%d/%d", r.StartYear-1, r.StartYear))

			leagueHl := getCenteredHL(cells.Objects[1])
			leagueHl.SetText(r.LeagueName)
			leagueHl.OnTapped = func() { navigate(LeagueHistory, r.LeagueId) }

			teamHl := getCenteredHL(cells.Objects[2])
			teamHl.SetText(r.TeamName)
			teamHl.OnTapped = func() {
				navigate(TeamDetails, r.TeamId)
			}

			pLbl := cells.Objects[3].(*widget.Label)
			pLbl.SetText(fmt.Sprintf("%d", r.Played))

			gLbl := cells.Objects[4].(*widget.Label)
			gLbl.SetText(fmt.Sprintf("%d", r.Goals))

			sLbl := cells.Objects[5].(*widget.Label)
			score := "-"
			if r.Played > 0 {
				score = fmt.Sprintf("%.2f", r.Score/float64(r.Played))
			}
			sLbl.SetText(score)

			costLbl := cells.Objects[6].(*widget.Label)
			cost := "-"
			if r.TransferCost != nil {
				cost = *r.TransferCost
			}
			costLbl.SetText(cost)
		},
	)

	return NewFborder().Top(headers).Get(historyList)
}

func makePlayerMainDetailsView(player *models.PlayerDetailed, canSeeDetails bool, showStats bool, ctx *AppContext, g *models.Game) *fyne.Container {
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

	teamInfo := widget.NewCard("", "Team Info",
		centered(widget.NewLabel("Free agent")),
	)
	if player.Team != nil {
		teamInfo.SetContent(
			container.NewVBox(
				hL(player.Team.Name, func() { ctx.PushWithParam(TeamDetails, player.Team.Id) }),
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
			))
	}

	main := container.NewGridWithRows(2,
		container.NewGridWithColumns(2,
			teamInfo,
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
