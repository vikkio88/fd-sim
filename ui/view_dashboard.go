package ui

import (
	"fdsim/conf"
	d "fdsim/db"
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
	"golang.org/x/exp/slices"
)

func dashboardView(ctx *AppContext) *fyne.Container {
	gameId := ctx.RouteParam.(string)
	game := ctx.InitGameState(gameId)
	news, emails := loadNotifications(ctx.Db)

	dateStr := binding.NewString()
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
	navigate := ctx.PushWithParam
	newsMailsTabs := container.NewAppTabs(
		container.NewTabItemWithIcon("News", theme.DocumentIcon(), makeNewsTab(news, ctx.Db, navigate)),
		container.NewTabItemWithIcon("Emails", theme.MailComposeIcon(), makeEmailsTab(emails, ctx.Db, navigate)),
	)
	main := container.NewGridWithColumns(2, navigation, newsMailsTabs)
	sim := services.NewSimulator(game, ctx.Db)

	trigNotifyNoDb := widget.NewButtonWithIcon("Notif NODB", theme.InfoIcon(), func() {
		email := models.NewEmail("Bla@bla.com", "Some Stuff", "Do things", game.Date, []models.Link{})
		emails.Prepend(email)

		n := models.NewNews("Something Happened", "Corriere della Sera", "Some Stuff", game.Date, []models.Link{})
		news.Prepend(n)
	})

	trigNotifyDb := widget.NewButtonWithIcon("Notif DB", theme.InfoIcon(), func() {
		email := models.NewEmail(
			"someguylink@bla.com", "Here you go some Links!",
			fmt.Sprintf("Do some stuff things %s other text", conf.LinkBodyPH),
			game.Date,
			[]models.Link{models.NewLink("Test", Test.String(), nil)},
		)
		emails.Prepend(email)
		ctx.Db.GameR().AddEmails([]*models.Email{email})

		n := models.NewNews("Something Happened", "Corriere della Sera", "Some Stuff", game.Date, []models.Link{})
		news.Prepend(n)
		ctx.Db.GameR().AddNews([]*models.News{n})
	})

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
					trigNotifyNoDb,
					trigNotifyDb,
					nextDay,
					nextWeek,
				)).Get(),
		).
		Get(
			main,
		)
}

func loadNotifications(db d.IDb) (binding.UntypedList, binding.UntypedList) {
	emailsDb := db.GameR().GetEmails()
	newsDb := db.GameR().GetNews()
	emails := binding.NewUntypedList()
	for _, e := range emailsDb {
		emails.Prepend(e)
	}

	news := binding.NewUntypedList()
	for _, n := range newsDb {
		news.Prepend(n)
	}
	return news, emails
}

func makeNewsTab(news binding.UntypedList, db d.IDb, navigate func(AppRoute, any)) fyne.CanvasObject {
	list := widget.NewListWithData(
		news,
		func() fyne.CanvasObject {
			return container.NewMax(
				NewFborder().
					Left(widget.NewIcon(theme.DocumentIcon())).
					Right(
						widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {}),
					).
					Get(
						container.NewVBox(
							widget.NewLabel(""),
							// widget.NewLabel(""),
						),
					),
			)
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			newsI := vm.NewsFromDi(di)
			newsContainer := co.(*fyne.Container).Objects[0].(*fyne.Container)
			newsInfoCtr := newsContainer.Objects[0].(*fyne.Container) //.Objects[0].(*fyne.Container)
			mainLbl := newsInfoCtr.Objects[0].(*widget.Label)
			mainLbl.SetText(newsI.String())
			mainLbl.TextStyle = fyne.TextStyle{Bold: !newsI.Read}
			mainLbl.Refresh()

			leftIcon := newsContainer.Objects[1].(*widget.Icon)
			if newsI.Read {
				leftIcon.SetResource(theme.FileTextIcon())
			}
			deleteBtn := newsContainer.Objects[2].(*widget.Button)
			deleteBtn.OnTapped = func() {
				db.GameR().DeleteNews(newsI.Id)
				items, _ := news.Get()
				index := slices.IndexFunc(items, func(item any) bool {
					e := item.(*models.News)
					return e.Id == newsI.Id
				})
				items = append(items[:index], items[index+1:]...)
				news.Set(items)
			}
		})

	list.OnSelected = func(id widget.ListItemID) {
		di, _ := news.GetItem(id)
		news := vm.NewsFromDi(di)
		if !news.Read {
			news.Read = true
			list.Refresh()
			db.GameR().MarkNewsAsRead(news.Id)
		}
		list.UnselectAll()
		navigate(News, news.Id)
	}
	return list
}

func makeEmailsTab(emails binding.UntypedList, db d.IDb, navigate func(AppRoute, any)) fyne.CanvasObject {
	list := widget.NewListWithData(
		emails,
		func() fyne.CanvasObject {
			return container.NewMax(
				NewFborder().
					Left(widget.NewIcon(theme.MailComposeIcon())).
					Right(
						widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {}),
					).
					Get(
						container.NewVBox(
							widget.NewLabel(""),
						),
					),
			)
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			email := vm.EmailFromDi(di)
			emailCtr := co.(*fyne.Container).Objects[0].(*fyne.Container)
			mailInfoCtr := emailCtr.Objects[0].(*fyne.Container) //.Objects[0].(*fyne.Container)
			mainLbl := mailInfoCtr.Objects[0].(*widget.Label)
			mainLbl.SetText(email.String())
			mainLbl.TextStyle = fyne.TextStyle{Bold: !email.Read}
			mainLbl.Refresh()

			leftIcon := emailCtr.Objects[1].(*widget.Icon)
			if email.Read {
				leftIcon.SetResource(widgets.Icon("email_read").Resource)
			}
			deleteBtn := emailCtr.Objects[2].(*widget.Button)
			deleteBtn.OnTapped = func() {
				db.GameR().DeleteEmail(email.Id)
				items, _ := emails.Get()
				index := slices.IndexFunc(items, func(item any) bool {
					e := item.(*models.Email)
					return e.Id == email.Id
				})
				items = append(items[:index], items[index+1:]...)
				emails.Set(items)
			}
		})

	list.OnSelected = func(id widget.ListItemID) {
		di, _ := emails.GetItem(id)
		email := vm.EmailFromDi(di)
		if !email.Read {
			email.Read = true
			list.Refresh()
			db.GameR().MarkNewsAsRead(email.Id)
		}
		list.UnselectAll()
		navigate(Email, email.Id)
	}
	return list
}

func simTriggers(dateStr binding.String, news, emails binding.UntypedList, game *models.Game, sim *services.Simulator, events []*services.Event) {
	dateStr.Set(game.Date.Format(conf.DateFormatGame))
	newEmails, newNews := sim.SettleEventsTriggers(events)
	for _, e := range newEmails {
		emails.Prepend(e)
	}
	for _, n := range newNews {
		news.Prepend(n)
	}
}
