package ui

import (
	"fdsim/conf"
	"fdsim/models"
	"fdsim/utils"
	"fdsim/vm"
	"fdsim/widgets"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func playerDetailsView(ctx *AppContext) *fyne.Container {
	id := ""
	subTab := -1

	switch v := ctx.RouteParam.(type) {
	case string:
		id = v
	case vm.SubTabIdParam:
		id = v.Id
		subTab = v.SubtabIndex
	}

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
			container.NewTabItemWithIcon("Manage", theme.AccountIcon(), centered(widget.NewLabel("Manage"))),
		)
	}

	main.Append(container.NewTabItemWithIcon("History",
		theme.DocumentIcon(),
		makePHistory(player.History, ctx.PushWithParam),
	))

	main.Append(container.NewTabItemWithIcon("Awards",
		theme.DocumentIcon(),
		makePAwards(player.Awards, player.Trophies, ctx.PushWithParam),
	))

	if g.IsEmployed() && !isManagedPlayer {
		main.Append(container.NewTabItemWithIcon("Transfer",
			widgets.Icon("transfers").Resource,
			makePTransferTab(ctx, player, canSeeDetails),
		))
	}
	handleSubtabs(subTab, main)

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
			container.NewMax(
				NewFborder().Top(
					centered(
						valueLabel("Retired in:", widget.NewLabel(fmt.Sprintf("%d (%d years old)", retired.YearRetired, retired.Age))),
					),
				).Get(
					container.NewGridWithRows(2,
						container.NewPadded(
							makePHistory(retired.History, ctx.PushWithParam),
						),
						container.NewVBox(
							container.NewPadded(
								makePAwards(retired.Awards, retired.Trophies, ctx.PushWithParam),
							),
						),
					),
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

func makePAwards(awards []models.Award, trophies []models.Trophy, navigate NavigateWithParamFunc) fyne.CanvasObject {
	columns := widgets.NewColumnsLayout([]float32{100, 100, 100, 100})
	headersAwards := widgets.NewListHeader(
		[]widgets.ListColumn{
			widgets.NewListCol("League", fyne.TextAlignCenter),
			widgets.NewListCol("Team", fyne.TextAlignCenter),
			widgets.NewListCol("Awards", fyne.TextAlignCenter),
			widgets.NewListCol("Stats", fyne.TextAlignCenter),
		},
		columns,
	)
	headersTrophies := widgets.NewListHeader(
		[]widgets.ListColumn{
			widgets.NewListCol("Year", fyne.TextAlignCenter),
			widgets.NewListCol("League", fyne.TextAlignCenter),
			widgets.NewListCol("Team", fyne.TextAlignCenter),
			widgets.NewListCol("Trophy", fyne.TextAlignCenter),
		},
		columns,
	)

	var awardList fyne.CanvasObject
	awardList = container.NewVBox(centered(h2("Nothing here")))
	if len(awards) > 0 {
		awardList = widget.NewList(
			func() int { return len(awards) },
			func() fyne.CanvasObject {
				return container.New(
					columns,
					centered(widget.NewHyperlink("League", nil)),
					centered(widget.NewHyperlink("Team", nil)),
					widget.NewLabel("Awards"),
					widget.NewLabel("Stats"),
				)
			},
			func(lii widget.ListItemID, co fyne.CanvasObject) {
				ar := awards[lii]
				cells := co.(*fyne.Container)

				leagueHl := getCenteredHL(cells.Objects[0])
				leagueHl.SetText(ar.LeagueName)
				leagueHl.OnTapped = func() { navigate(LeagueHistory, ar.LeagueId) }

				teamHl := getCenteredHL(cells.Objects[1])
				teamHl.SetText(ar.Team.Name)
				teamHl.OnTapped = func() {
					navigate(TeamDetails, ar.Team.Id)
				}

				awLbl := cells.Objects[2].(*widget.Label)
				awLbl.SetText(ar.String())

				statsLbl := cells.Objects[3].(*widget.Label)
				statsLbl.SetText(ar.StatString())

			},
		)
	}

	var trophiesList fyne.CanvasObject
	trophiesList = container.NewVBox(centered(h2("Nothing here")))
	if len(trophies) > 0 {
		trophiesList = widget.NewList(
			func() int { return len(trophies) },
			func() fyne.CanvasObject {
				return container.New(
					columns,
					widget.NewLabel("Year"),
					centered(widget.NewHyperlink("League", nil)),
					centered(widget.NewHyperlink("Team", nil)),
					widget.NewLabel("Trophy"),
				)
			},
			func(lii widget.ListItemID, co fyne.CanvasObject) {
				tf := trophies[lii]
				cells := co.(*fyne.Container)
				yearLbl := cells.Objects[0].(*widget.Label)
				yearLbl.SetText(fmt.Sprintf("%d", tf.Year))

				leagueHl := getCenteredHL(cells.Objects[1])
				leagueHl.SetText(tf.LeagueName)
				leagueHl.OnTapped = func() { navigate(LeagueHistory, tf.LeagueId) }

				teamHl := getCenteredHL(cells.Objects[2])
				teamHl.SetText(tf.Team.Name)
				teamHl.OnTapped = func() {
					navigate(TeamDetails, tf.Team.Id)
				}

				tfLbl := cells.Objects[3].(*widget.Label)
				tfLbl.SetText("League Winner")

			},
		)
	}
	return container.NewGridWithColumns(2,
		NewFborder().
			Top(
				container.NewVBox(
					h2("Individual Awards"),
					headersAwards,
				),
			).
			Get(awardList),
		NewFborder().
			Top(
				container.NewVBox(
					h2("Trophies"),
					headersTrophies,
				),
			).
			Get(trophiesList),
	)
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
				if r.TeamId != "" {
					navigate(TeamDetails, r.TeamId)
				}
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
	playerValue := "?"
	playerWage := "?"
	if canSeeDetails {
		skillInfo = centered(widget.NewLabel(player.Skill.String()))
		playerValue = player.Value.StringKMB()
		playerWage = player.Wage.StringKMB()
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
					centered(widget.NewLabel(playerValue)),
				),
				valueLabel("Contract:",
					widget.NewLabel(fmt.Sprintf("%s / %d years", playerWage, player.YContract)),
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
		main.Add(statsWrapper)
	}
	return main
}

func makePTransferTab(ctx *AppContext, player *models.PlayerDetailed, canSeeDetails bool) fyne.CanvasObject {
	g, _ := ctx.GetGameState()

	if offer, ok := player.GetOfferFromTeamId(g.Team.Id); ok {
		switch offer.Stage() {

		case models.OfstOffered:
			return centered(h2(
				fmt.Sprintf("Your already made an offer for this player on %s (%s). Waiting for response.", offer.OfferDate.Format(conf.DateFormatShort), offer.BidValue.StringKMB()),
			))
		case models.OfstTeamAccepted:
			return centered(h2("Team accepted the offer. Now you can discuss contract."))
		case models.OfstReadyTP:
			return centered(h2("Player and Team accepted your offer."))
		}
	}

	tInfo, ok := ctx.Db.MarketR().GetTransferMarketInfo()

	if !ok {
		// this should not happen as it wont appear if you have no team
		panic("you should not see this if you are hired")
	}

	iWage := getApproxMoney(player.IdealWage)
	wage := getApproxMoney(player.Wage)
	value := getApproxMoney(player.Value)

	lowerV, higherV := utils.GetApproxRangeF(player.Value.Value())
	lowerW, higherW := utils.GetApproxRangeF(player.IdealWage.Value())

	isFreeAgent := player.Team == nil
	actionBtn := widget.NewButton("Offer Contract", func() {
		contractY := 1
		ctx.PushWithParam(Chat, vm.ChatParams{
			IsPlayerOffer: true,
			Player:        player,
			ValueF:        lowerW,
			ValueF1:       higherW,
			ValueI:        &contractY,
		})
	})
	if !isFreeAgent {
		actionBtn = widget.NewButton("Make an Offer", func() {
			ctx.PushWithParam(Chat, vm.ChatParams{
				IsPlayerOffer: true,
				Player:        player,
				Team:          player.Team,
				ValueF:        lowerV,
				ValueF1:       higherV,
			})
		})
	}

	contractInfo := valueLabel("Contract", widget.NewLabel("-"))
	if !isFreeAgent {
		contractInfo = valueLabel("Contract", widget.NewLabel(fmt.Sprintf("%s / %d yrs", wage, player.YContract)))
	}

	return NewFborder().Top(
		centered(h2(fmt.Sprintf("Transfer Budget: %s", tInfo.TransferBudget.StringKMB()))),
	).
		Bottom(rightAligned(actionBtn)).
		Get(container.NewVBox(
			valueLabel("Value: ", widget.NewLabel(value)),
			valueLabel("Ideal Wage: ", widget.NewLabel(iWage)),
			contractInfo,
		),
		)

}
