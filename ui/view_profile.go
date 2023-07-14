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

func profileView(ctx *AppContext) *fyne.Container {
	game, _ := ctx.GetGameState()
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Info", theme.AccountIcon(), centered(h1("Account"))),
	)
	if game.IsEmployed() {
		stats := ctx.Db.GameR().GetFDStats()
		tabs.Append(
			container.NewTabItemWithIcon("Stats", theme.DocumentIcon(), makeFDStats(stats)),
		)
	}

	history := ctx.Db.GameR().GetFDHistory()
	tabs.Append(container.NewTabItemWithIcon("History", theme.DocumentIcon(), makeFDHistory(history, ctx.PushWithParam)))
	return NewFborder().
		Top(
			NewFborder().
				Left(topNavBar(ctx)).
				Get(centered(h1("Your Profile"))),
		).
		Get(
			tabs,
		)
}

func makeFDStats(stats *models.FDStatRow) fyne.CanvasObject {
	tHL := widget.NewHyperlink(stats.TeamName, nil)
	tHL.OnTapped = func() {
		fmt.Println(stats.TeamId)
	}
	return widget.NewCard(
		"", "Current Season Statistics",
		container.NewVBox(
			valueLabel("Team:",
				centered(tHL),
			),
			valueLabel("Players Signed:",
				centered(widget.NewLabel(fmt.Sprintf("%d", stats.PlayersSigned))),
			),
			valueLabel("Players Sold:",
				centered(widget.NewLabel(fmt.Sprintf("%d", stats.PlayersSold))),
			),
			valueLabel("Coaches Signed:",
				centered(widget.NewLabel(fmt.Sprintf("%d", stats.CoachesSigned))),
			),
			valueLabel("Coaches Sacked:",
				centered(widget.NewLabel(fmt.Sprintf("%d", stats.CoachesSacked))),
			),
			valueLabel("Max Spent on a Single Transfer:",
				centered(widget.NewLabel(stats.MaxSpent.StringKMB())),
			),
			valueLabel("Total Spent:",
				centered(widget.NewLabel(stats.TotalSpent.StringKMB())),
			),
			valueLabel("Max Cashed on a Single Transfer:",
				centered(widget.NewLabel(stats.MaxCashed.StringKMB())),
			),
			valueLabel("Total Cashed:",
				centered(widget.NewLabel(stats.TotalCashed.StringKMB())),
			),
		))
}

func makeFDHistory(history []*models.FDHistoryRow, navigate NavigateWithParamFunc) fyne.CanvasObject {
	if len(history) < 1 {
		return centered(widget.NewLabel("No History yet"))
	}
	columns := widgets.NewColumnsLayout([]float32{-1, 100, 100, 100, 100, 100, 100, 100})
	headers := widgets.NewListHeader(
		[]widgets.ListColumn{
			widgets.NewListCol("", fyne.TextAlignCenter),
			widgets.NewListCol("League", fyne.TextAlignCenter),
			widgets.NewListCol("Team", fyne.TextAlignLeading),
			widgets.NewListCol("Wage", fyne.TextAlignLeading),
			widgets.NewListCol("Players In", fyne.TextAlignLeading),
			widgets.NewListCol("Players Out", fyne.TextAlignLeading),
			widgets.NewListCol("Spent", fyne.TextAlignLeading),
			widgets.NewListCol("Cashed", fyne.TextAlignLeading),
		},
		columns,
	)

	historyList := widget.NewList(
		func() int { return len(history) },
		func() fyne.CanvasObject {
			return container.New(
				columns,
				widget.NewLabel("Year"),
				widget.NewLabel("LeagueName"),
				centered(widget.NewHyperlink("Team", nil)),
				widget.NewLabel("Wage"),
				widget.NewLabel("Signed"),
				widget.NewLabel("Sold"),
				widget.NewLabel("TSpent"),
				widget.NewLabel("TCashed"),
			)
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			r := history[lii]
			cell := co.(*fyne.Container)

			yearS := r.StartDate.Year()
			monthS := r.StartDate.Month()
			yearE := r.EndDate.Year()
			monthE := r.EndDate.Month()

			yearLbl := cell.Objects[0].(*widget.Label)
			yearLbl.SetText(fmt.Sprintf("%d/%d - %d/%d", yearS, monthS, yearE, monthE))

			leagueLbl := cell.Objects[1].(*widget.Label)
			leagueLbl.SetText(r.LeagueName)

			teamHl := getCenteredHL(cell.Objects[2])
			teamHl.SetText(r.TeamName)
			teamHl.OnTapped = func() {
				navigate(TeamDetails, r.TeamId)
			}

			wageLbl := cell.Objects[3].(*widget.Label)
			wageLbl.SetText(r.Wage.StringKMB())

			pSLbl := cell.Objects[4].(*widget.Label)
			pSLbl.SetText(fmt.Sprintf("%d", r.PlayersSigned))

			pSoldLbl := cell.Objects[5].(*widget.Label)
			pSoldLbl.SetText(fmt.Sprintf("%d", r.PlayersSold))

			spentLbl := cell.Objects[6].(*widget.Label)
			spentLbl.SetText(r.TotalSpent.StringKMB())

			cashedLbl := cell.Objects[7].(*widget.Label)
			cashedLbl.SetText(r.TotalCashed.StringKMB())

			//TODO: add other stats Maxes and Coaches
		},
	)

	return NewFborder().Top(headers).Get(historyList)
}
