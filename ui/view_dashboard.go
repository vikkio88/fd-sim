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
	loadGlobals(ctx.Db)

	dateStr = binding.NewString()
	dateStr.Set(game.Date.Format(conf.DateFormatGame))

	pendingDecisionIndicator := container.NewHBox(widget.NewIcon(theme.WarningIcon()), widget.NewLabel("Some actions are pending"))
	pendingDecisionIndicator.Hide()
	hasPendingDecisions.AddListener(binding.NewDataListener(func() {
		if has, err := hasPendingDecisions.Get(); err == nil {
			if has {
				pendingDecisionIndicator.Show()
			} else {
				pendingDecisionIndicator.Hide()
			}
		}

	}))

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

	toPersonal := widget.NewButtonWithIcon("Your Profile", theme.AccountIcon(), func() {
		ctx.Push(Profile)
	})

	toTeamMgmt := widget.NewButtonWithIcon("Manage Team", widgets.Icon("team").Resource, func() {
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

	trigTest := widget.NewButtonWithIcon("Trig Test", theme.InfoIcon(), func() {
		ctx.PushWithParam(League, "leId_01H5TDM0QNBFPGQDZZH968QTPx")
	})

	startSim := widget.NewButtonWithIcon("Simulate", theme.MediaPlayIcon(), func() {
		ctx.Push(Simulation)
	})

	nextDay := widget.NewButtonWithIcon("Next Day", theme.MediaSkipNextIcon(), func() {
		events, simulated := sim.Simulate(1)
		if simulated {
			simTriggers(dateStr, news, emails, game, sim, events)
			freePendingDecisions()
		} else {
			checkForEmailDialog(ctx.GetWindow())
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
							widget.NewLabelWithData(dateStr),
						),
					),
				).
				Right(
					centered(
						container.NewVBox(
							container.NewHBox(
								widget.NewIcon(theme.AccountIcon()),
								widget.NewLabel(fmt.Sprintf("%s %s (%d)", fd.Name, fd.Surname, fd.Age)),
							),
							starsFromPerc(fd.Fame),
						),
					),
				).
				Get(),
		).
		Bottom(
			NewFborder().
				Right(
					container.NewHBox(
						trigEm,
						trigTest,
						nextDay,
						startSim,
					)).Get(pendingDecisionIndicator),
		).
		Get(
			main,
		)
}

func checkForEmailDialog(window fyne.Window) {
	dialog.ShowInformation("Check your Emails", "Some Emails need a reply!", window)
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
