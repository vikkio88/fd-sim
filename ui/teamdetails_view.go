package ui

import (
	"fdsim/viewmodels"
	"fmt"

	"fyne.io/fyne/v2"
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
				Get(widget.NewLabel(team.String()))).
		Get(
			widget.NewListWithData(
				roster,
				simpleRosterListRow,
				makeSimpleRosterRowBind(ctx),
			),
		)
}

func simpleRosterListRow() fyne.CanvasObject {
	return NewFborder().
		Left(widget.NewHyperlink("", nil)).
		Get()
}

func makeSimpleRosterRowBind(ctx *AppContext) func(di binding.DataItem, co fyne.CanvasObject) {
	return func(di binding.DataItem, co fyne.CanvasObject) {
		player := viewmodels.PlayerFromDi(di)
		c := co.(*fyne.Container)
		l := c.Objects[0].(*widget.Hyperlink)
		l.SetText(player.String())
		l.OnTapped = func() {
			fmt.Println(player.Id)
		}
	}
}
