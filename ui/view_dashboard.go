package ui

import (
	"fdsim/conf"
	"fdsim/models"
	"fdsim/services"
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

	newsx, emailsx := loadNotifications(ctx.Db)
	news = newsx
	emails = emailsx

	dateStr = binding.NewString()
	dateStr.Set(game.Date.Format(conf.DateFormatGame))

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
		ctx.Push(Calendar)
	})

	toPersonal := widget.NewButtonWithIcon("Profile", theme.AccountIcon(), func() {
		ctx.Push(Profile)
	})

	toTeamMgmt := widget.NewButtonWithIcon("Team", widgets.Icon("team").Resource, func() {
		ctx.Push(TeamMgmt)
	})

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
	navigate := ctx.PushWithParam
	newsMailsTabs := makeNotificationsTabs(ctx, navigate)
	main := container.NewGridWithColumns(2, navigation, newsMailsTabs)
	sim := services.NewSimulator(game, ctx.Db)

	trigEm := widget.NewButtonWithIcon("Trig Email", theme.InfoIcon(), func() {
		randomTeam := ctx.Db.TeamR().GetRandom()

		email := models.NewEmail(
			fmt.Sprintf("hr@%s.com", randomTeam.Name),
			"Testing Email",
			fmt.Sprintf("random email, from this team %s", conf.LinkBodyPH),
			game.Date,
			[]models.Link{
				models.NewLink(randomTeam.Name, TeamDetails.String(), &randomTeam.Id),
			},
		)

		ctx.Db.GameR().AddEmails([]*models.Email{email})
		emails.Prepend(email)
	})

	trigNw := widget.NewButtonWithIcon("Trig News", theme.InfoIcon(), func() {
		randomTeam := ctx.Db.TeamR().GetRandom()

		newsI := models.NewNews(
			fmt.Sprintf("%s did something", randomTeam.Name),
			"Random Newspaper",
			fmt.Sprintf("random news, from this team %s", conf.LinkBodyPH),
			game.Date,
			[]models.Link{
				models.NewLink(randomTeam.Name, TeamDetails.String(), &randomTeam.Id),
			},
		)

		ctx.Db.GameR().AddNews([]*models.News{newsI})
		news.Prepend(newsI)
	})

	startSim := widget.NewButtonWithIcon("Simulate", theme.MediaPlayIcon(), func() {
		ctx.Push(Simulation)
	})

	nextDay := widget.NewButtonWithIcon("Next Day", theme.MediaSkipNextIcon(), func() {
		events, simulated := sim.Simulate(1)
		if simulated {
			simTriggers(dateStr, news, emails, game, sim, events)
		} else {
			dialog.ShowInformation("Check your Emails", "Some Emails need a reply!", ctx.w)
		}

	})

	nextWeek := widget.NewButtonWithIcon("Next Week", theme.MediaFastForwardIcon(), func() {
		events, simulated := sim.Simulate(7)
		if simulated {
			simTriggers(dateStr, news, emails, game, sim, events)
		} else {
			dialog.ShowInformation("Check your Emails", "Some Emails need a reply!", ctx.w)
		}
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
					trigEm,
					trigNw,
					nextDay,
					nextWeek,
					startSim,
				)).Get(),
		).
		Get(
			main,
		)
}

func simTriggers(dateStr binding.String, news, emails binding.UntypedList, game *models.Game, sim *services.Simulator, events []*services.Event) (int, int) {
	dateStr.Set(game.Date.Format(conf.DateFormatGame))
	newEmails, newNews := sim.SettleEventsTriggers(events)
	for _, e := range newEmails {
		emails.Prepend(e)
	}
	for _, n := range newNews {
		news.Prepend(n)
	}

	// this is used in Simulation as it returns the numbers of news notifications
	return len(newEmails), len(newNews)
}
