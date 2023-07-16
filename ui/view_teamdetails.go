package ui

import (
	"fdsim/enums"
	"fdsim/models"
	vm "fdsim/vm"
	"fdsim/widgets"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func teamDetailsView(ctx *AppContext) *fyne.Container {
	_, isGameInit := ctx.GetGameState()
	id := ctx.RouteParam.(string)
	team, exists := ctx.Db.TeamR().ById(id)
	if !exists {
		return notFoundView(ctx, "Team")
	}

	stats := container.NewMax(centered(widget.NewLabel("No match played yet")))
	if isGameInit {
		row := ctx.Db.TeamR().TableRow(id)
		stats.RemoveAll()
		stats.Add(makeTeamStats(row))
	}
	roster := binding.NewUntypedList()
	for _, p := range team.Roster.PlayersByRole() {
		roster.Append(p)
	}

	coachSkillInfo := starsFromPerc(team.Coach.Skill)
	if IsFDTeam(id) {
		coachSkillInfo = widget.NewLabel(team.Coach.Skill.String())
	}

	coach := widget.NewCard(
		"",
		"Coach",
		container.NewVBox(
			centered(widget.NewLabel(fmt.Sprintf("%s (%d)", team.Coach.String(), team.Coach.Age))),
			centered(
				widgets.FlagIcon(team.Coach.Country),
			),
			centered(
				coachSkillInfo,
			),
			container.NewGridWithColumns(2,
				widget.NewLabel("Contract"),
				widget.NewLabel(fmt.Sprintf("%s / %d ys", team.Coach.Wage.StringKMB(), team.Coach.YContract)),
			),
			container.NewGridWithColumns(2,
				widget.NewLabel("Module:"),
				widget.NewLabel(team.Coach.Module.String()),
			),
		),
	)

	finances := widget.NewCard("",
		"Finances",
		container.NewVBox(
			container.NewGridWithColumns(2,
				widgets.Icon("money"),
				widget.NewLabel(team.Balance.StringKMB()),
			),
			container.NewGridWithColumns(2,
				widgets.Icon("contract"),
				widget.NewLabel(team.Wages().StringKMB()),
			),
			container.NewGridWithColumns(2,
				widgets.Icon("transfers"),
				widget.NewLabel(team.TransferBudget().StringKMB()),
			),
		),
	)
	playerIcon := widget.NewIcon(theme.AccountIcon())
	playerIcon.Hide()

	teamAvgSkillInfo := starsFromf64(team.Roster.AvgSkill())
	if IsFDTeam(id) {
		teamAvgSkillInfo = widget.NewLabel(
			fmt.Sprintf("%.2f", team.Roster.AvgSkill()),
		)
		playerIcon.Show()
	}

	teamDetails := container.NewVBox(
		centered(
			container.NewHBox(
				widgets.Icon("city"),
				widget.NewLabel(fmt.Sprintf("%s (%s)", team.City, team.Country)),
			),
		),
		container.NewGridWithRows(3,
			centered(
				container.NewHBox(
					widgets.Icon("team"),
					widget.NewLabel(
						fmt.Sprintf("%d", team.Roster.Len()),
					),
				),
			),
			centered(
				container.NewHBox(
					widgets.Icon("dumbell"),
					teamAvgSkillInfo,
				),
			),
			centered(
				container.NewHBox(
					widget.NewLabel("Avg Age"),
					widget.NewLabel(fmt.Sprintf("%.2f", team.Roster.AvgAge())),
				),
			),
		),
		container.NewGridWithColumns(2,
			coach,
			finances,
		),
	)

	main := container.NewAppTabs(
		container.NewTabItemWithIcon("Club Details", widgets.Icon("city").Resource, teamDetails),
		container.NewTabItemWithIcon("Roster", widgets.Icon("team").Resource, rosterUi(roster, ctx, id)),
		container.NewTabItemWithIcon("Season Stats", theme.DocumentIcon(), stats),
		container.NewTabItemWithIcon("History", theme.DocumentIcon(), makeTHistory(team.History, ctx.PushWithParam)),
	)

	return NewFborder().
		Top(
			NewFborder().Left(topNavBar(ctx)).
				Get(
					centered(
						container.NewHBox(
							playerIcon,
							h1(team.Name),
							widgets.FlagIcon(team.Country),
						),
					),
				)).
		Get(
			main,
		)
}

func makeTHistory(tHistoryRow []*models.THistoryRow, navigate NavigateWithParamFunc) fyne.CanvasObject {
	if len(tHistoryRow) < 1 {
		return centered(widget.NewLabel("No History yet"))
	}
	columns := widgets.NewColumnsLayout([]float32{-1, 100, 100, 100, 50, 50, 50, 50, 50})
	headers := widgets.NewListHeader(
		[]widgets.ListColumn{
			widgets.NewListCol("", fyne.TextAlignCenter),
			widgets.NewListCol("League", fyne.TextAlignCenter),
			widgets.NewListCol("Position", fyne.TextAlignLeading),
			widgets.NewListCol("Points", fyne.TextAlignLeading),
			widgets.NewListCol("W", fyne.TextAlignLeading),
			widgets.NewListCol("D", fyne.TextAlignLeading),
			widgets.NewListCol("L", fyne.TextAlignLeading),
			widgets.NewListCol("GS", fyne.TextAlignLeading),
			widgets.NewListCol("GC", fyne.TextAlignLeading),
		},
		columns,
	)

	historyList := widget.NewList(
		func() int { return len(tHistoryRow) },
		func() fyne.CanvasObject {
			return container.New(
				columns,
				widget.NewLabel("Year"),
				hL("League", func() {}),
				widget.NewLabel("Position"),
				widget.NewLabel("Points"),
				widget.NewLabel("W"),
				widget.NewLabel("D"),
				widget.NewLabel("L"),
				widget.NewLabel("GS"),
				widget.NewLabel("GC"),
			)
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			r := tHistoryRow[lii]
			cell := co.(*fyne.Container)

			yearLbl := cell.Objects[0].(*widget.Label)
			yearLbl.SetText(fmt.Sprintf("%d/%d", r.Year-1, r.Year))

			leagueHl := getCenteredHL(cell.Objects[1])
			leagueHl.SetText(r.LeagueName)
			leagueHl.OnTapped = func() {
				navigate(LeagueHistory, r.LeagueId)
			}

			posLbl := cell.Objects[2].(*widget.Label)
			posLbl.SetText(fmt.Sprintf("%d", r.FinalPosition))

			ptsLbl := cell.Objects[3].(*widget.Label)
			ptsLbl.SetText(fmt.Sprintf("%d", r.Points))

			wLbl := cell.Objects[4].(*widget.Label)
			wLbl.SetText(fmt.Sprintf("%d", r.Wins))

			dLbl := cell.Objects[5].(*widget.Label)
			dLbl.SetText(fmt.Sprintf("%d", r.Draws))

			lLbl := cell.Objects[6].(*widget.Label)
			lLbl.SetText(fmt.Sprintf("%d", r.Losses))

			gsLbl := cell.Objects[7].(*widget.Label)
			gsLbl.SetText(fmt.Sprintf("%d", r.GoalScored))

			gcLbl := cell.Objects[8].(*widget.Label)
			gcLbl.SetText(fmt.Sprintf("%d", r.GoalConceded))
		},
	)

	return NewFborder().Top(headers).Get(historyList)
}

func makeTeamStats(row *models.TPHRow) fyne.CanvasObject {
	return widget.NewCard(
		"", "Current Season Statistics",
		container.NewVBox(
			valueLabel("Position:",
				centered(widget.NewLabel(fmt.Sprintf("%d", row.Index))),
			),
			valueLabel("Played:",
				centered(widget.NewLabel(fmt.Sprintf("%d", row.Row.Played))),
			),
			valueLabel("Wins:",
				centered(widget.NewLabel(fmt.Sprintf("%d", row.Row.Wins))),
			),
			valueLabel("Draws:",
				centered(widget.NewLabel(fmt.Sprintf("%d", row.Row.Draws))),
			),
			valueLabel("Losses:",
				centered(widget.NewLabel(fmt.Sprintf("%d", row.Row.Losses))),
			),
			valueLabel("Goals Scored:",
				centered(widget.NewLabel(fmt.Sprintf("%d", row.Row.GoalScored))),
			),
			valueLabel("Goals Conceded:",
				centered(widget.NewLabel(fmt.Sprintf("%d", row.Row.GoalConceded))),
			),
			valueLabel("Points:",
				centered(widget.NewLabel(fmt.Sprintf("%d", row.Row.Points))),
			),
		))
}

func rosterUi(roster binding.DataList, ctx *AppContext, teamId string) fyne.CanvasObject {
	return widget.NewListWithData(
		roster,
		simpleRosterListRow,
		makeSimpleRosterRowBind(ctx, teamId),
	)
}

func simpleRosterListRow() fyne.CanvasObject {
	return NewFborder().
		Get(
			container.NewMax(
				container.NewGridWithColumns(
					4,
					centered(widget.NewHyperlink("", nil)),
					centered(widget.NewLabel("Role")),
					centered(widgets.NewFlag(enums.EN)),
					centered(
						container.NewHBox(
							starsFromf64(0),
							widget.NewLabel(""),
						),
					),
				),
			),
		)
}

func makeSimpleRosterRowBind(ctx *AppContext, teamId string) func(di binding.DataItem, co fyne.CanvasObject) {
	return func(di binding.DataItem, co fyne.CanvasObject) {
		player := vm.PlayerFromDi(di)
		c := co.(*fyne.Container)

		ctn := c.Objects[0].(*fyne.Container)
		mx := ctn.Objects[0].(*fyne.Container)
		ctr := mx.Objects[0].(*fyne.Container)
		l := ctr.Objects[0].(*widget.Hyperlink)

		l.SetText(fmt.Sprintf("%s (%d)", player.String(), player.Age))
		l.OnTapped = func() {
			ctx.PushWithParam(PlayerDetails, player.Id)
		}

		mx.Objects[1].(*fyne.Container).Objects[0].(*widget.Label).SetText(player.Role.String())
		f := mx.Objects[2].(*fyne.Container).Objects[0].(*widgets.Flag)
		f.SetCountry(player.Country)
		values := mx.Objects[3].(*fyne.Container).Objects[0].(*fyne.Container)

		star := values.Objects[0].(*widgets.StarRating)
		value := values.Objects[1].(*widget.Label)
		if IsFDTeam(teamId) {
			value.SetText(player.Skill.String())
			star.Hide()
		} else {
			star.SetValues(vm.PercToStars(player.Skill))
			value.Hide()
		}
	}
}
