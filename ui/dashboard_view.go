package ui

import (
	"fdsim/conf"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func dashboardView(ctx *AppContext) *fyne.Container {
	// gamestate
	gameId := ctx.RouteParam.(string)
	game := ctx.Db.GameR().ById(gameId)
	fd := game.FootDirector()
	saveBtn := widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {})
	exitBtn := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		dialog.ShowConfirm("Exit Game", "Are you sure?", func(b bool) {
			if !b {
				return
			}

			ctx.NavigateTo(Main)
		}, ctx.GetWindow())
	})
	return NewFborder().
		Top(
			NewFborder().
				Left(
					container.NewGridWithRows(2,
						container.NewHBox(
							exitBtn,
							saveBtn,
						),
					),
				).
				Right(widget.NewLabel(game.Date.Format(conf.GameDateFormat))).
				Get(centered(
					container.NewVBox(
						widget.NewLabel(fmt.Sprintf("%s %s (%d)", fd.Name, fd.Surname, fd.Age)),
						starsFromPerc(fd.Fame),
					),
				)),
		).
		Get(
			container.NewCenter(
				widget.NewButton("League", func() {
					// ctx.PushWithParam(League, )
					ctx.Push(League)
				}),
			),
		)
	// 	container.NewVBox(
	// 		widget.NewLabel("Dashboard"),
	// 		widget.NewLabel(fmt.Sprintf("%s %s (%d)", fd.Name, fd.Surname, fd.Age)),
	// 		starsFromPerc(fd.Fame),
	// 		widget.NewButton("League", func() {
	// 			// ctx.PushWithParam(League, )
	// 			ctx.Push(League)
	// 		}),
	// 	),
	// )
}
