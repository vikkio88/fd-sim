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
	roundsView := makeRounds(rounds, navigate)
	statsView := container.NewCenter(widget.NewLabel("Stats"))
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

func makeRounds(rounds []*models.RPHTPH, navigate func(AppRoute, any)) fyne.CanvasObject {
	c := container.NewGridWrap(fyne.NewSize(350, 250))

	for _, r := range rounds {
		c.AddObject(makeRound(r))
	}

	return container.NewVScroll(c)
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

	header.DisableSorting = true
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
