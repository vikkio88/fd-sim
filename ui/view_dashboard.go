package ui

import (
	"fdsim/conf"
	"fdsim/libs"
	"fdsim/models"
	"fdsim/widgets"
	"fmt"
	"time"

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

	// nextDay := widget.NewButtonWithIcon("Next Day", theme.MediaSkipNextIcon(), func() {})
	nextDay := widget.NewButtonWithIcon("Simulate Round", theme.MediaSkipNextIcon(), func() {
		//TODO: move this to a simulation helper
		league := ctx.Db.LeagueR().ByIdFull(game.LeagueId)
		r, ok := league.NextRound()
		if !ok {
			dialog.ShowInformation("Finished!", "No more rounds to play!", ctx.GetWindow())
			return
		}
		rng := libs.NewRng(time.Now().Unix())
		r.Simulate(rng)
		league.Update(r)
		oldStats := ctx.Db.LeagueR().GetStats(league.Id)
		stats := models.StatsFromRoundResult(r, game.LeagueId)
		newStats := models.MergeStats(oldStats, stats)

		ctx.Db.LeagueR().PostRoundUpdate(r, league)
		ctx.Db.LeagueR().UpdateStats(newStats)

		dialog.ShowInformation("Finished!", fmt.Sprintf("Simulated round %d", r.Index+1), ctx.GetWindow())
	})
	nextWeek := widget.NewButtonWithIcon("Next Week", theme.MediaFastForwardIcon(), func() {})

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
