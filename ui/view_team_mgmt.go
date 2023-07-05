package ui

import (
	"fdsim/models"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func teamMgmtView(ctx *AppContext) *fyne.Container {
	game, _ := ctx.GetGameState()

	if !game.IsEmployed() {
		return NewFborder().
			Top(leftAligned(backButton(ctx))).
			Get(
				container.NewCenter(
					widget.NewLabel("You have no team to manage."),
				),
			)
	}

	team := ctx.Db.TeamR().ById(game.Team.Id)

	return NewFborder().
		Top(
			NewFborder().
				Left(backButton(ctx)).
				Get(centered(h1(fmt.Sprintf("%s - Management", game.Team.Name)))),
		).
		Get(
			container.NewAppTabs(
				container.NewTabItem("Roster", makeRosterManagement(team)),
				container.NewTabItem("Finance", centered(widget.NewLabel("Finance"))),
				container.NewTabItem("Board/Supporters", centered(widget.NewLabel("Board/Supporters"))),
				container.NewTabItem("Transfer Market", centered(widget.NewLabel("Transfer"))),
				container.NewTabItem("Misc", centered(widget.NewLabel("Misc"))),
			),
		)
}

func makeRosterManagement(team *models.Team) fyne.CanvasObject {
	return container.NewMax(
		container.NewGridWithColumns(2,
			makeLineup(team),
			centered(widget.NewLabel("Performance")),
		))
}

func makeLineup(team *models.Team) fyne.CanvasObject {
	lineup := team.Lineup()
	lineupList := widget.NewList(func() int {
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
	return container.NewMax(
		container.NewGridWithColumns(2,
			container.NewPadded(
				lineupList,
			),
			container.NewVBox(
				widget.NewLabel(fmt.Sprintf("Module: %s", lineup.Module.String())),
				//TODO: Add Lineup Role Stats
				//lineup.SectorStat
			),
		),
	)
}
