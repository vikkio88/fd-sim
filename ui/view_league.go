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

func leagueView(ctx *AppContext) *fyne.Container {
	leagueId := ctx.RouteParam.(string)
	league := ctx.Db.LeagueR().ByIdFull(leagueId)
	navigate := ctx.PushWithParam
	rows := league.TableRows()
	leagueTable := container.NewMax(
		makeTableView(rows, navigate),
	)
	rounds := league.RoundsPH()
	results := ctx.Db.LeagueR().GetAllResults()
	roundsView := makeRounds(rounds, results, navigate, league.RPointer)
	statsView := makeStats(ctx, leagueId)
	main := container.NewAppTabs(
		container.NewTabItemWithIcon("Table", theme.ListIcon(), leagueTable),
		container.NewTabItemWithIcon("Rounds", theme.GridIcon(), roundsView),
		container.NewTabItemWithIcon("Stats", theme.StorageIcon(), statsView),
	)
	return NewFborder().
		Top(NewFborder().Left(backButton(ctx)).Get(centered(h1("League")))).
		Get(
			container.NewMax(main),
		)
}

func makeRounds(rounds []*models.RPHTPH, results models.ResultsPHMap, navigate func(AppRoute, any), roundPointer int) fyne.CanvasObject {
	cToPlay := container.NewGridWrap(fyne.NewSize(350, 250))
	cPlayed := container.NewCenter(widget.NewLabel("Nothing yet..."))

	playedRounds := rounds[:roundPointer]
	if len(playedRounds) > 0 {

		cPlayed = container.NewGridWrap(fyne.NewSize(350, 250))
	}
	toPlayRounds := rounds[roundPointer:]

	for _, r := range toPlayRounds {
		cToPlay.AddObject(makeRound(r))
	}

	for i := len(playedRounds) - 1; i > 0; i-- {
		cPlayed.AddObject(makeRoundWithResults(playedRounds[i], results, navigate))
	}

	return container.NewAppTabs(
		container.NewTabItem("Upcoming Matches", container.NewVScroll(cToPlay)),
		container.NewTabItem("Played", container.NewVScroll(cPlayed)),
	)

}

func makeRound(round *models.RPHTPH) fyne.CanvasObject {
	matchList := widget.NewList(
		func() int {
			return len(round.Matches)
		},
		func() fyne.CanvasObject {
			return container.NewGridWithColumns(3,
				widget.NewLabel(""),
				centered(widget.NewLabel("vs")),
				widget.NewLabel(""),
			)

		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			m := round.Matches[lii]
			co.(*fyne.Container).Objects[0].(*widget.Label).SetText(m.Home.Name)
			co.(*fyne.Container).Objects[2].(*widget.Label).SetText(m.Away.Name)
		})
	c := widget.NewCard("", fmt.Sprintf("Round %d", round.Index+1), matchList)
	return container.NewPadded(c)
}

func makeRoundWithResults(round *models.RPHTPH, results models.ResultsPHMap, navigate func(AppRoute, any)) fyne.CanvasObject {
	matchList := widget.NewList(
		func() int {
			return len(round.Matches)
		},
		func() fyne.CanvasObject {
			return container.NewGridWithColumns(3,
				widget.NewLabel(""),
				centered(widget.NewHyperlink("vs", nil)),
				widget.NewLabel(""),
			)

		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			m := round.Matches[lii]
			co.(*fyne.Container).Objects[0].(*widget.Label).SetText(m.Home.Name)
			if result, ok := results[m.Id]; ok {
				resHL := co.(*fyne.Container).Objects[1].(*fyne.Container).Objects[0].(*widget.Hyperlink)
				resHL.SetText(result.String())
				resHL.OnTapped = func() {
					navigate(MatchDetails, m.Id)
				}
			}
			co.(*fyne.Container).Objects[2].(*widget.Label).SetText(m.Away.Name)
		})
	c := widget.NewCard("", fmt.Sprintf("Round %d", round.Index+1), matchList)
	return container.NewPadded(c)
}

func makeTableView(table []*models.TPHRow, navigate func(AppRoute, any)) *fyne.Container {
	columns := widgets.NewColumnsLayout([]float32{-1, 350, 50, 50, 50, 50, 50, 50, 50})
	header := widgets.NewListHeader(
		[]widgets.ListColumn{
			widgets.NewListCol("", fyne.TextAlignCenter),
			widgets.NewListCol("", fyne.TextAlignCenter),
			widgets.NewListCol("P", fyne.TextAlignLeading),
			widgets.NewListCol("W", fyne.TextAlignLeading),
			widgets.NewListCol("D", fyne.TextAlignLeading),
			widgets.NewListCol("L", fyne.TextAlignLeading),
			widgets.NewListCol("GS", fyne.TextAlignLeading),
			widgets.NewListCol("GC", fyne.TextAlignLeading),
			widgets.NewListCol("Points", fyne.TextAlignLeading),
		},
		columns,
	)

	return NewFborder().
		Top(header).
		Get(
			widget.NewList(
				func() int {
					return len(table)
				},
				teamTableRow(columns),
				func(lii widget.ListItemID, co fyne.CanvasObject) {
					teamRow := table[lii]
					c := co.(*fyne.Container)
					c.Objects[0].(*widget.Label).SetText(fmt.Sprintf("%d . ", teamRow.Index+1))
					c.Objects[1].(*fyne.Container).Objects[0].(*widget.Hyperlink).SetText(teamRow.Team.Name)
					c.Objects[1].(*fyne.Container).Objects[0].(*widget.Hyperlink).OnTapped = func() {
						navigate(TeamDetails, teamRow.Team.Id)
					}
					c.Objects[2].(*widget.Label).SetText(fmt.Sprintf("%d", teamRow.Row.Played))
					c.Objects[4].(*widget.Label).SetText(fmt.Sprintf("%d", teamRow.Row.Draws))
					c.Objects[3].(*widget.Label).SetText(fmt.Sprintf("%d", teamRow.Row.Wins))
					c.Objects[5].(*widget.Label).SetText(fmt.Sprintf("%d", teamRow.Row.Losses))
					c.Objects[6].(*widget.Label).SetText(fmt.Sprintf("%d", teamRow.Row.GoalScored))
					c.Objects[7].(*widget.Label).SetText(fmt.Sprintf("%d", teamRow.Row.GoalConceded))
					c.Objects[8].(*widget.Label).SetText(fmt.Sprintf("%d", teamRow.Row.Points))
				},
			),
		)
}

func teamTableRow(layout *widgets.ColumnsLayout) func() fyne.CanvasObject {
	return func() fyne.CanvasObject {
		return container.New(layout,
			widget.NewLabel("#"),
			centered(widget.NewHyperlink("", nil)),
			widget.NewLabel("P"),
			widget.NewLabel("W"),
			widget.NewLabel("D"),
			widget.NewLabel("L"),
			widget.NewLabel("GS"),
			widget.NewLabel("GC"),
			widget.NewLabel("Points"),
		)
	}
}

func makeStats(ctx *AppContext, leagueId string) fyne.CanvasObject {
	stats := ctx.Db.LeagueR().BestScorers(leagueId)
	columns := widgets.NewColumnsLayout([]float32{-1, 350, 250, 100, 100})
	header := statsHeader(columns)

	return NewFborder().
		Top(header).
		Get(widget.NewList(
			func() int {
				return len(stats)
			},
			scorerTableRow(columns),
			func(lii widget.ListItemID, co fyne.CanvasObject) {
				statRow := stats[lii]
				c := co.(*fyne.Container)
				c.Objects[0].(*widget.Label).SetText(fmt.Sprintf("%d.", statRow.Index+1))
				playerHL := c.Objects[1].(*fyne.Container).Objects[0].(*widget.Hyperlink)
				playerHL.SetText(statRow.Player.String())
				playerHL.OnTapped = func() {
					ctx.PushWithParam(PlayerDetails, statRow.Player.Id)
				}
				teamHL := c.Objects[2].(*fyne.Container).Objects[0].(*widget.Hyperlink)
				teamHL.SetText(statRow.Team.Name)
				teamHL.OnTapped = func() {
					ctx.PushWithParam(TeamDetails, statRow.Team.Id)
				}
				c.Objects[3].(*widget.Label).SetText(fmt.Sprintf("%d", statRow.Played))
				c.Objects[4].(*widget.Label).SetText(fmt.Sprintf("%d", statRow.Goals))
			}),
		)
}

func statsHeader(columns *widgets.ColumnsLayout) *widgets.ListHeader {
	return widgets.NewListHeader(
		[]widgets.ListColumn{
			widgets.NewListCol("", fyne.TextAlignCenter),
			widgets.NewListCol("Player", fyne.TextAlignCenter),
			widgets.NewListCol("Team", fyne.TextAlignCenter),
			widgets.NewListCol("Played", fyne.TextAlignLeading),
			widgets.NewListCol("Goals", fyne.TextAlignLeading),
		},
		columns,
	)
}

func scorerTableRow(layout *widgets.ColumnsLayout) func() fyne.CanvasObject {
	return func() fyne.CanvasObject {
		return container.New(layout,
			widget.NewLabel("#"),
			centered(widget.NewHyperlink("", nil)),
			centered(widget.NewHyperlink("", nil)),
			widget.NewLabel("Played"),
			widget.NewLabel("Goals"),
		)
	}
}