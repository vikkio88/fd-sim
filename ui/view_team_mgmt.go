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
			Top(leftAligned(topNavBar(ctx))).
			Get(
				container.NewCenter(
					widget.NewLabel("You have no team to manage."),
				),
			)
	}

	team, _ := ctx.Db.TeamR().ById(game.Team.Id)
	trow := ctx.Db.TeamR().TableRow(team.Id)

	trsf := "closed"
	if game.IsTransferWindowOpen() {
		trsf = "open"
	}

	return NewFborder().
		Top(
			NewFborder().
				Left(topNavBar(ctx)).
				Get(centered(h1(fmt.Sprintf("%s - Management", game.Team.Name)))),
		).
		Get(
			container.NewAppTabs(
				container.NewTabItem("Roster", makeRosterManagement(team, trow, ctx.PushWithParam)),
				container.NewTabItem("Finance", centered(widget.NewLabel("Finance"))),
				container.NewTabItem("Board/Supporters", centered(widget.NewLabel("Board/Supporters"))),
				container.NewTabItem("Transfer Market", centered(widget.NewLabel(fmt.Sprintf("Transfer window is %s.", trsf)))),
				container.NewTabItem("Misc", centered(widget.NewLabel("Misc"))),
			),
		)
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
				widget.NewLabel(fmt.Sprintf("Module: %s", lineup.Module.String())),
				stats,
			),
		),
	)
}
