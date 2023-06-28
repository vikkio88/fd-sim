package ui

import (
	"fdsim/models"
	"fdsim/services"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type NotificationCount struct {
	NewEmails int
	NewNews   int
}

func simulationView(ctx *AppContext) *fyne.Container {
	game, _ := ctx.GetGameState()
	sim := services.NewSimulator(game, ctx.Db)
	emailsNum := 0
	newsNum := 0
	emails := binding.NewString()
	news := binding.NewString()
	state := binding.NewString()

	stop := make(chan int)
	quit := make(chan int)
	dayFinished := make(chan NotificationCount)

	go func() {
		for {
			select {
			case msg := <-dayFinished:
				emailsNum += msg.NewEmails
				newsNum += msg.NewNews
				emails.Set(fmt.Sprintf("%d", emailsNum))
				news.Set(fmt.Sprintf("%d", newsNum))
			case <-quit:
				return
			}
		}
	}()
	backBtn := backButton(ctx)
	var startBtn *widget.Button
	startBtn = widget.NewButtonWithIcon("Start", theme.MediaPlayIcon(), func() {
		state.Set("Simulating...")
		go start(game, sim, dayFinished, stop, quit)
		backBtn.Disable()
		startBtn.Disable()
	})

	return NewFborder().
		Top(
			NewFborder().
				Left(backBtn).
				Get(
					centered(h1("Simulation")),
				),
		).
		Bottom(
			rightAligned(
				container.NewHBox(
					startBtn,
					widget.NewButtonWithIcon("Stop", theme.MediaStopIcon(), func() {
						state.Set("Stopping...")
						stop <- 1
						ctx.Pop()
					}),
				),
			),
		).
		Get(
			centered(
				container.NewVBox(
					container.NewVBox(
						centered(
							widget.NewLabelWithData(state),
						),
						widget.NewLabelWithData(dateStr),
					),
					centered(
						container.NewGridWithColumns(2,
							widget.NewIcon(theme.DocumentIcon()),
							widget.NewLabelWithData(news),
						),
					),
					centered(
						container.NewGridWithColumns(2,
							widget.NewIcon(theme.MailComposeIcon()),
							widget.NewLabelWithData(emails),
						),
					),
				),
			),
		)
}

func start(game *models.Game, sim *services.Simulator, messages chan NotificationCount, stop, quit chan int) {
	for {
		select {
		case <-stop:
			{
				quit <- 0
				return
			}
		default:
			{
				events := sim.Simulate(1)
				emailsC, newsC := simTriggers(dateStr, news, emails, game, sim, events)
				messages <- NotificationCount{NewEmails: emailsC, NewNews: newsC}
				time.Sleep(time.Duration(2) * time.Second)
			}
		}
	}
}
