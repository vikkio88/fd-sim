package ui

import (
	"fdsim/conf"
	"fdsim/models"
	"fdsim/services"
	"fdsim/vm"
	"fdsim/widgets"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func dashboardView(ctx *AppContext) *fyne.Container {
	gameId := ctx.RouteParam.(string)
	game := ctx.InitGameState(gameId)
	dateStr := binding.NewString()
	dateStr.Set(game.Date.Format(conf.GameDateFormat))

	emails := binding.NewUntypedList()
	news := binding.NewUntypedList()

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
		container.NewTabItemWithIcon("News", theme.DocumentIcon(), makeNewsTab(news)),
		container.NewTabItemWithIcon("Emails", theme.MailComposeIcon(), makeEmailsTab(emails)),
	)
	main := container.NewGridWithColumns(2, navigation, newsMailsTabs)
	sim := services.NewSimulator(game, ctx.Db)

	nextDay := widget.NewButtonWithIcon("Next Day", theme.MediaSkipNextIcon(), func() {
		events := sim.Simulate(1)
		simTriggers(dateStr, news, emails, game, sim, events)

	})

	nextWeek := widget.NewButtonWithIcon("Next Week", theme.MediaFastForwardIcon(), func() {
		events := sim.Simulate(7)
		simTriggers(dateStr, news, emails, game, sim, events)
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
				Right(widget.NewLabelWithData(dateStr)).
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

func makeNewsTab(news binding.UntypedList) fyne.CanvasObject {
	// if news.Length() < 1 {
	// 	return widget.NewLabel("No news...")
	// }

	list := widget.NewListWithData(
		news,
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			news := vm.NewsFromDi(di)
			co.(*widget.Label).SetText(news.String())
			// if is unread
			co.(*widget.Label).TextStyle = fyne.TextStyle{Bold: true}
		})
	return list
}

func makeEmailsTab(emails binding.UntypedList) fyne.CanvasObject {
	// if emails.Length() < 1 {
	// 	return widget.NewLabel("No emails...")
	// }

	list := widget.NewListWithData(
		emails,
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			email := vm.EmailFromDi(di)
			co.(*widget.Label).SetText(email.String())
		})
	return list
}

func simTriggers(dateStr binding.String, news, emails binding.UntypedList, game *models.Game, sim *services.Simulator, events []*services.Event) {
	dateStr.Set(game.Date.Format(conf.GameDateFormat))
	newEmails, newNews := sim.SettleEventsTriggers(events)
	for _, e := range newEmails {
		emails.Prepend(e)
	}
	for _, n := range newNews {
		news.Prepend(n)
	}
}
