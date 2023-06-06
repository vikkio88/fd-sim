package ui

import (
	"fdsim/utils"
	"fdsim/viewmodels"
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
	return NewFborder().
		Top(
			NewFborder().Left(backButton(ctx)).
				Get(
					centered(
						container.NewHBox(
							widget.NewLabel(team.String()),
							small(team.Country.String()),
						),
					),
				)).
		Get(
			container.NewGridWithColumns(2,
				rosterUi(roster, ctx),
				container.NewVBox(stars(utils.NewPerc(10))),
			),
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
					widget.NewLabel("Nantionality"),
				),
			),
		)
}

func makeSimpleRosterRowBind(ctx *AppContext) func(di binding.DataItem, co fyne.CanvasObject) {
	return func(di binding.DataItem, co fyne.CanvasObject) {
		player := viewmodels.PlayerFromDi(di)
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
