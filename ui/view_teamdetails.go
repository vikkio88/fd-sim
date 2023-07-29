package ui

import (
	"fdsim/models"
	"fdsim/widgets"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
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

	coach := makeCoachCard(team.Coach, IsFDTeam(id), false)

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
		container.NewTabItemWithIcon("Roster", widgets.Icon("team").Resource, rosterUi(team, ctx)),
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

func makeCoachCard(coach *models.Coach, showSkillInfo bool, interactive bool) fyne.CanvasObject {
	coachSkillInfo := starsFromPerc(coach.Skill)
	if showSkillInfo {
		coachSkillInfo = widget.NewLabel(coach.Skill.String())
	}

	details := container.NewVBox(
		centered(widget.NewLabel(fmt.Sprintf("%s (%d)", coach.String(), coach.Age))),
		centered(
			widgets.FlagIcon(coach.Country),
		),
		centered(
			coachSkillInfo,
		),
		container.NewGridWithColumns(2,
			widget.NewLabel("Contract"),
			widget.NewLabel(fmt.Sprintf("%s / %d ys", coach.Wage.StringKMB(), coach.YContract)),
		),
	)

	if !interactive {
		details.Add(
			container.NewGridWithColumns(2,
				widget.NewLabel("Module:"),
				widget.NewLabel(coach.Module.String()),
			),
		)
	} else {
		details.Add(
			widget.NewButton("Chat", func() { fmt.Println("Chat") }),
		)
	}

	return widget.NewCard(
		"",
		"Coach",
		details,
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
