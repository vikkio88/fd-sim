package ui

import (
	"fdsim/models"
	vm "fdsim/vm"
	"fdsim/widgets"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func teamMgmtView(ctx *AppContext) *fyne.Container {
	game, _ := ctx.GetGameState()

	if !game.IsEmployed() {
		return NewFborder().
			Top(leftAligned(topNavBar(ctx))).
			Get(
				container.NewCenter(
					widget.NewLabel("You have no team to manage."),
				),
			)
	}

	team, _ := ctx.Db.TeamR().ById(game.Team.Id)
	trow := ctx.Db.TeamR().TableRow(team.Id)

	return NewFborder().
		Top(
			NewFborder().
				Left(topNavBar(ctx)).
				Get(centered(h1(fmt.Sprintf("%s - Management", game.Team.Name)))),
		).
		Get(
			container.NewAppTabs(
				container.NewTabItemWithIcon("Roster", widgets.Icon("team").Resource, makeRosterManagement(team, trow, ctx.PushWithParam)),
				container.NewTabItemWithIcon("Finance", widgets.Icon("money").Resource, centered(widget.NewLabel("Finance"))),
				container.NewTabItemWithIcon("Board/Supporters", theme.AccountIcon(), centered(widget.NewLabel("Board/Supporters"))),
				container.NewTabItemWithIcon("Transfer Market", widgets.Icon("transfers").Resource, makeMarketMgMtTab(ctx)),
				container.NewTabItemWithIcon("Misc", theme.SettingsIcon(), centered(widget.NewLabel("Misc"))),
			),
		)
}

func makeMarketMgMtTab(ctx *AppContext) fyne.CanvasObject {
	game, _ := ctx.GetGameState()
	isWindowOpen, _ := game.IsTransferWindowOpen()

	offers := ctx.Db.MarketR().GetOffersByOfferingTeamId(game.GetTeamIdOrEmpty())

	trsf := "CLOSED"
	if isWindowOpen {
		trsf = "OPEN"
	}
	return NewFborder().Top(
		centered(widget.NewLabel(fmt.Sprintf("Transfer window is %s.", trsf))),
	).Get(
		makeOffersList(offers, ctx.PushWithParam),
	)
}

func makeOffersList(offers []*models.Offer, navigate NavigateWithParamFunc) fyne.CanvasObject {
	columns := widgets.NewColumnsLayout([]float32{-1, 100, 150, 100, 100, 100})
	header := widgets.NewListHeader(
		[]widgets.ListColumn{
			widgets.NewListCol("", fyne.TextAlignCenter),
			widgets.NewListCol("Team", fyne.TextAlignLeading),
			widgets.NewListCol("Stage", fyne.TextAlignLeading),
			widgets.NewListCol("Bid", fyne.TextAlignLeading),
			widgets.NewListCol("Wage", fyne.TextAlignLeading),
			widgets.NewListCol("", fyne.TextAlignLeading),
		},
		columns,
	)
	list := widget.NewList(
		func() int {
			return len(offers)
		},
		func() fyne.CanvasObject {
			return container.New(
				columns,
				hL("", func() {}),
				hL("Team", func() {}),
				widget.NewLabel("Stage"),
				widget.NewLabel("Offer"),
				widget.NewLabel("Wage"),
				widget.NewButtonWithIcon("", theme.ZoomInIcon(), func() {}),
			)
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			offer := offers[lii]
			ctr := co.(*fyne.Container)
			hlP := getCenteredHL(ctr.Objects[0])
			hlP.SetText(offer.Player.String())
			hlP.OnTapped = func() {
				navigate(
					PlayerDetails,
					offer.Player.Id,
				)
			}

			tlHl := getCenteredHL(ctr.Objects[1])
			bidlbl := ctr.Objects[3].(*widget.Label)
			tlHl.SetText("-")
			if offer.Team != nil {
				tlHl.SetText(offer.Team.Name)
				tlHl.OnTapped = func() { navigate(TeamDetails, offer.Team.Id) }
				bidlbl.SetText(offer.BidValue.StringKMB())
			}

			ofStlbl := ctr.Objects[2].(*widget.Label)
			ofStlbl.SetText(offer.Stage().String())

			wagelbl := ctr.Objects[4].(*widget.Label)
			wagelbl.SetText("-")
			if offer.WageValue != nil {
				wagelbl.SetText(offer.WageValue.StringKMB())
			}

			goBtn := ctr.Objects[5].(*widget.Button)
			goBtn.OnTapped = func() {
				navigate(PlayerDetails,
					vm.SubTabIdParam{
						Id: offer.Player.Id,
						// SUBTAB MARKET
						//TODO: maybe move this to a const
						SubtabIndex: 3,
					})
			}

		},
	)

	return NewFborder().Top(header).Get(list)
}

func makeRosterManagement(team *models.TeamDetailed, trow *models.TPHRow, navigate NavigateWithParamFunc) fyne.CanvasObject {
	teamDetailsBtn := widget.NewButton(fmt.Sprintf("%s - Team Details", team.Name), func() { navigate(TeamDetails, team.Id) })
	teamDetailsBtn.Importance = widget.LowImportance

	return NewFborder().
		Top(
			teamDetailsBtn,
		).
		Get(
			container.NewMax(
				container.NewGridWithColumns(2,
					makeLineup(team, navigate),
					makeTeamStats(trow),
				)),
		)
}

func makeLineup(team *models.TeamDetailed, navigate NavigateWithParamFunc) fyne.CanvasObject {
	lineup := team.Lineup()
	lineupList := widget.NewList(
		func() int {
			return len(lineup.FlatPlayers)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewLabel(""),
				widget.NewLabel(""),
				widget.NewLabel(""),
			)
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			playerId := lineup.FlatPlayers[lii].Id
			role := co.(*fyne.Container).Objects[0].(*widget.Label)
			name := co.(*fyne.Container).Objects[1].(*widget.Label)
			skill := co.(*fyne.Container).Objects[2].(*widget.Label)
			player, _ := team.Roster.Player(playerId)
			role.SetText(player.Role.StringShort())
			name.SetText(player.StringShort())
			skill.SetText(player.Skill.String())
		},
	)
	lineupList.OnSelected = func(id widget.ListItemID) {
		navigate(PlayerDetails, lineup.FlatPlayers[id].Id)
		lineupList.Unselect(id)
	}

	roles := models.AllPlayerRoles()
	stats := container.NewVBox()
	for _, r := range roles {
		stat := lineup.SectorStat[r]
		stats.Add(
			container.NewHBox(
				widget.NewLabel(r.StringShort()),
				widget.NewLabel(fmt.Sprintf("%.0f%%", stat.Skill)),
			),
		)

	}
	return container.NewMax(
		container.NewGridWithColumns(2,
			container.NewPadded(
				lineupList,
			),
			container.NewVBox(
				widget.NewLabel("Coach"),
				makeCoachDetails(team.Coach, navigate, true, true),
				widget.NewLabel(fmt.Sprintf("Module: %s", lineup.Module.String())),
				stats,
			),
		),
	)
}
