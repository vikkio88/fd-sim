package ui

import (
	"fdsim/enums"
	"fdsim/models"
	vm "fdsim/vm"
	"fdsim/widgets"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func rosterUi(team *models.TeamDetailed, ctx *AppContext, isGameInit bool) fyne.CanvasObject {
	if team.Id != fdTeamId {
		return simpleRoster(team, ctx.PushWithParam)
	}

	return container.NewAppTabs(
		container.NewTabItem("Roles", simpleRoster(team, ctx.PushWithParam)),
		container.NewTabItem("Contracts", contractRoster(team, ctx.PushWithParam)),
		container.NewTabItem("Performance", performanceRoster(team, ctx.PushWithParam)),
	)
}

func simpleRoster(team *models.TeamDetailed, navigate NavigateWithParamFunc) fyne.CanvasObject {
	players := team.Roster.PlayersByRole()
	return widget.NewList(
		func() int { return len(players) },
		func() fyne.CanvasObject {
			return container.NewMax(
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
			)
		},
		func(index widget.ListItemID, co fyne.CanvasObject) {
			player := players[index]
			cnt := co.(*fyne.Container)
			max := cnt.Objects[0].(*fyne.Container)
			grid := max.Objects[0].(*fyne.Container)
			l := grid.Objects[0].(*widget.Hyperlink)

			l.SetText(fmt.Sprintf("%s (%d)", player.String(), player.Age))
			l.OnTapped = func() {
				navigate(PlayerDetails, player.Id)
			}

			max.Objects[1].(*fyne.Container).Objects[0].(*widget.Label).SetText(player.Role.String())
			f := max.Objects[2].(*fyne.Container).Objects[0].(*widgets.Flag)
			f.SetCountry(player.Country)
			values := max.Objects[3].(*fyne.Container).Objects[0].(*fyne.Container)

			star := values.Objects[0].(*widgets.StarRating)
			value := values.Objects[1].(*widget.Label)
			if IsFDTeam(team.Id) {
				value.SetText(player.Skill.String())
				star.Hide()
			} else {
				star.SetValues(vm.PercToStars(player.Skill))
				value.Hide()
			}
		},
	)
}

func contractRoster(team *models.TeamDetailed, navigate NavigateWithParamFunc) fyne.CanvasObject {
	players := team.Roster.PlayersByContract()

	columns := widgets.NewColumnsLayout([]float32{-1, 100, 100, 100, 100})
	headers := widgets.NewListHeader(
		[]widgets.ListColumn{
			widgets.NewListCol("", fyne.TextAlignCenter),
			widgets.NewListCol("Skill", fyne.TextAlignLeading),
			widgets.NewListCol("Wage/Yr", fyne.TextAlignLeading),
			widgets.NewListCol("C Yrs Left", fyne.TextAlignLeading),
			widgets.NewListCol("Value", fyne.TextAlignLeading),
		},
		columns,
	)

	rosterByContract := widget.NewList(
		func() int { return len(players) },
		func() fyne.CanvasObject {
			return container.New(
				columns,
				hL("Player", func() {}),
				widget.NewLabel("Skill"),
				widget.NewLabel("Wage"),
				widget.NewLabel("Contract"),
				widget.NewLabel("Value"),
			)
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			p := players[lii]
			row := co.(*fyne.Container)

			pHl := getCenteredHL(row.Objects[0])
			pHl.SetText(fmt.Sprintf("%s (%d)", p.String(), p.Age))
			pHl.OnTapped = func() {
				navigate(PlayerDetails, p.Id)
			}

			skillL := row.Objects[1].(*widget.Label)
			//TODO: remember to change this if you want to show this for other teams too
			skillL.SetText(p.Skill.String())

			wageL := row.Objects[2].(*widget.Label)
			wageL.SetText(p.Wage.StringKMB())

			contL := row.Objects[3].(*widget.Label)
			contL.SetText(fmt.Sprintf("%d", p.YContract))

			vL := row.Objects[4].(*widget.Label)
			vL.SetText(p.Value.StringKMB())
		})

	return NewFborder().Top(headers).Get(rosterByContract)
}

func performanceRoster(team *models.TeamDetailed, navigate NavigateWithParamFunc) fyne.CanvasObject {
	// TODO: need to get stats for players of this team and bind them to the row
	players := team.Roster.PlayersByValueSkill()

	columns := widgets.NewColumnsLayout([]float32{-1, 100, 100, 100})
	headers := widgets.NewListHeader(
		[]widgets.ListColumn{
			widgets.NewListCol("", fyne.TextAlignCenter),
			widgets.NewListCol("Skill", fyne.TextAlignLeading),
			widgets.NewListCol("Value", fyne.TextAlignLeading),
			widgets.NewListCol("Stats", fyne.TextAlignLeading),
		},
		columns,
	)

	rosterByPerf := widget.NewList(
		func() int { return len(players) },
		func() fyne.CanvasObject {
			return container.New(
				columns,
				hL("Player", func() {}),
				widget.NewLabel("Skill"),
				widget.NewLabel("Wage"),
				widget.NewLabel("Value"),
				widget.NewLabel("Stats"),
			)
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			p := players[lii]
			row := co.(*fyne.Container)

			pHl := getCenteredHL(row.Objects[0])
			pHl.SetText(fmt.Sprintf("%s (%d)", p.String(), p.Age))
			pHl.OnTapped = func() {
				navigate(PlayerDetails, p.Id)
			}

			skillL := row.Objects[1].(*widget.Label)
			//TODO: remember to change this if you want to show this for other teams too
			skillL.SetText(p.Skill.String())

			vL := row.Objects[2].(*widget.Label)
			vL.SetText(p.Value.StringKMB())

			// TODO: bind Stats here
			contL := row.Objects[3].(*widget.Label)
			contL.SetText("-")
		})

	return NewFborder().Top(headers).Get(rosterByPerf)
}
