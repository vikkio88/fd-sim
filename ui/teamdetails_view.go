package ui

import (
	vm "fdsim/vm"
	"fdsim/widgets"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func teamDetailsView(ctx *AppContext) *fyne.Container {
	id := ctx.RouteParam.(string)
	team := ctx.Db.TeamR().ById(id)
	roster := binding.NewUntypedList()
	for _, p := range team.Roster.Players() {
		roster.Append(p)
	}

	coach := widget.NewCard(
		"",
		"Coach",
		container.NewVBox(
			centered(widget.NewLabel(team.Coach.String())),
			container.NewGridWithColumns(2,
				widget.NewLabel(fmt.Sprintf("%d", team.Coach.Age)),
				widget.NewLabel(team.Coach.Country.Nationality()),
			),
			centered(
				starsFromPerc(team.Coach.Skill),
			),
			container.NewGridWithColumns(2,
				widget.NewLabel("Module:"),
				widget.NewLabel(team.Coach.Module.String()),
			),
		),
	)

	finances := widget.NewCard("",
		"Finances",
		container.NewVBox(
			container.NewGridWithColumns(2,
				widgets.Icon("money"),
				widgets.Icon("meh_face"),
			),
		),
	)

	teamDetails := widget.NewCard(
		"",
		"Team Details",
		container.NewVBox(
			container.NewGridWithColumns(2,
				widgets.Icon("city"),
				widget.NewLabel(team.City),
			),
			container.NewHBox(
				widgets.Icon("dumbell"),
				starsFromf64(team.Roster.AvgSkill()),
			),
			container.NewHBox(
				widget.NewLabel("Avg Age"),
				widget.NewLabel(fmt.Sprintf("%.2f", team.Roster.AvgAge())),
			),
			coach,
			finances,
		),
	)

	main := container.NewHSplit(
		rosterUi(roster, ctx),
		teamDetails,
	)

	main.SetOffset(1.0)

	return NewFborder().
		Top(
			NewFborder().Left(backButton(ctx)).
				Get(
					centered(
						container.NewHBox(
							h1(team.Name),
							small(team.Country.String()),
						),
					),
				)).
		Get(
			main,
		)
}

func rosterUi(roster binding.DataList, ctx *AppContext) fyne.CanvasObject {
	return widget.NewListWithData(
		roster,
		simpleRosterListRow,
		makeSimpleRosterRowBind(ctx),
	)
}

func simpleRosterListRow() fyne.CanvasObject {
	return NewFborder().
		Get(
			container.NewMax(
				container.NewGridWithColumns(
					4,
					centered(widget.NewHyperlink("", nil)),
					widget.NewLabel("Age"),
					widget.NewLabel("Role"),
					widget.NewLabel("Nationality"),
				),
			),
		)
}

func makeSimpleRosterRowBind(ctx *AppContext) func(di binding.DataItem, co fyne.CanvasObject) {
	return func(di binding.DataItem, co fyne.CanvasObject) {
		player := vm.PlayerFromDi(di)
		c := co.(*fyne.Container)

		ctn := c.Objects[0].(*fyne.Container)
		mx := ctn.Objects[0].(*fyne.Container)
		ctr := mx.Objects[0].(*fyne.Container)
		l := ctr.Objects[0].(*widget.Hyperlink)

		l.SetText(player.String())
		mx.Objects[1].(*widget.Label).SetText(fmt.Sprintf("%d", player.Age))
		mx.Objects[2].(*widget.Label).SetText(player.Role.String())
		mx.Objects[3].(*widget.Label).SetText(player.Country.Nationality())
		l.OnTapped = func() {
			ctx.PushWithParam(PlayerDetails, player.Id)
		}
	}
}
