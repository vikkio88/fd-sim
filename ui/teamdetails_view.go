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
			centered(widget.NewLabel(fmt.Sprintf("%s (%d)", team.Coach.String(), team.Coach.Age))),
			centered(
				widget.NewLabel(team.Coach.Country.Nationality()),
			),
			centered(
				starsFromPerc(team.Coach.Skill),
			),
			container.NewGridWithColumns(2,
				widget.NewLabel("Contract"),
				widget.NewLabel(fmt.Sprintf("%s / %d years", team.Coach.Wage.StringKMB(), team.Coach.YContract)),
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
				widget.NewLabel(team.Balance.StringKMB()),
			),
			container.NewGridWithColumns(2,
				widgets.Icon("transfers"),
				widget.NewLabel(team.TransferBudget().StringKMB()),
			),
		),
	)

	teamDetails := widget.NewCard(
		"",
		"Team Details",
		container.NewVBox(
			container.NewGridWithColumns(2,
				widgets.Icon("city"),
				widget.NewLabel(fmt.Sprintf("%s (%s)", team.City, team.Country)),
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
							widgets.Flag(team.Country),
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
					centered(widget.NewLabel("Role")),
					centered(widget.NewLabel("Nationality")),
					centered(starsFromf64(0)),
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

		l.SetText(fmt.Sprintf("%s (%d)", player.String(), player.Age))
		l.OnTapped = func() {
			ctx.PushWithParam(PlayerDetails, player.Id)
		}

		mx.Objects[1].(*fyne.Container).Objects[0].(*widget.Label).SetText(player.Role.String())
		mx.Objects[2].(*fyne.Container).Objects[0].(*widget.Label).SetText(player.Country.Nationality())
		s := mx.Objects[3].(*fyne.Container).Objects[0].(*widgets.StarRating)
		s.SetValues(vm.PercToStars(player.Skill))
	}
}
