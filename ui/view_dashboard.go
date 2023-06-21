package ui

import (
	"fdsim/conf"
	"fdsim/services"
	"fdsim/widgets"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func dashboardView(ctx *AppContext) *fyne.Container {
	gameId := ctx.RouteParam.(string)
	game := ctx.InitGameState(gameId)
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

	toLeague := widget.NewButtonWithIcon("League", theme.ListIcon(), func() {
		ctx.PushWithParam(League, game.LeagueId)
	})

	toCalendar := widget.NewButtonWithIcon("Calendar", theme.GridIcon(), func() {

	})
	toCalendar.Disable()

	toPersonal := widget.NewButtonWithIcon("Profile", theme.AccountIcon(), func() {

	})
	toTeamMgmt := widget.NewButtonWithIcon("Team", widgets.Icon("team").Resource, func() {

	})
	toTeamMgmt.Disable()

	navigation := container.NewCenter(
		container.NewPadded(
			container.NewGridWithRows(2,
				container.NewGridWithColumns(2,
					toLeague, toCalendar,
				),
				container.NewGridWithColumns(2,
					toPersonal, toTeamMgmt,
				),
			),
		),
	)
	newsMailsTabs := container.NewAppTabs(
		container.NewTabItemWithIcon("News", theme.DocumentIcon(), widget.NewLabel("Here there will be news...")),
		container.NewTabItemWithIcon("Emails", theme.MailComposeIcon(), widget.NewLabel("Here there will be emails...")),
	)
	main := container.NewGridWithColumns(2, navigation, newsMailsTabs)
	sim := services.NewSimulator(ctx.gameState, ctx.Db)

	nextDay := widget.NewButtonWithIcon("Next Day", theme.MediaSkipNextIcon(), func() {
		events := sim.Simulate(1)
		fmt.Println(events)
		dialog.ShowInformation("Done", "Done", ctx.GetWindow())
	})

	nextWeek := widget.NewButtonWithIcon("Next Week", theme.MediaFastForwardIcon(), func() {
		events := sim.Simulate(7)
		fmt.Println(events)
		dialog.ShowInformation("Done", "Done", ctx.GetWindow())
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
		Bottom(
			NewFborder().Right(
				container.NewHBox(
					nextDay,
					nextWeek,
				)).Get(),
		).
		Get(
			main,
		)
}
