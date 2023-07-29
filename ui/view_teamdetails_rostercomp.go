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
	"fyne.io/fyne/v2/widget"
)

func rosterUi(team *models.TeamDetailed, ctx *AppContext, isGameInit bool) fyne.CanvasObject {
	if team.Id != fdTeamId {
		return simpleRoster(team, ctx.PushWithParam)
	}

	return container.NewAppTabs(
		container.NewTabItem("Roles", simpleRoster(team, ctx.PushWithParam)),
		container.NewTabItem("Contracts", simpleRoster(team, ctx.PushWithParam)),
		container.NewTabItem("Performance", simpleRoster(team, ctx.PushWithParam)),
	)
}

func simpleRoster(team *models.TeamDetailed, navigate NavigateWithParamFunc) fyne.CanvasObject {
	roster := binding.NewUntypedList()
	for _, p := range team.Roster.PlayersByRole() {
		roster.Append(p)
	}

	return widget.NewListWithData(
		roster,
		simpleRosterListRow,
		makeSimpleRosterRowBind(navigate, team.Id),
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

func makeSimpleRosterRowBind(navigate NavigateWithParamFunc, teamId string) func(di binding.DataItem, co fyne.CanvasObject) {
	return func(di binding.DataItem, co fyne.CanvasObject) {
		player := vm.PlayerFromDi(di)
		c := co.(*fyne.Container)

		ctn := c.Objects[0].(*fyne.Container)
		mx := ctn.Objects[0].(*fyne.Container)
		ctr := mx.Objects[0].(*fyne.Container)
		l := ctr.Objects[0].(*widget.Hyperlink)

		l.SetText(fmt.Sprintf("%s (%d)", player.String(), player.Age))
		l.OnTapped = func() {
			navigate(PlayerDetails, player.Id)
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
