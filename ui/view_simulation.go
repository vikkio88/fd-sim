package ui

import (
	"fdsim/conf"
	"fdsim/models"
	"fdsim/services"
	"fdsim/widgets"
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
	navBtn := topNavBar(ctx)
	placeholder := widget.NewLabel("0")
	var startBtn *widget.Button

	stopBtn := widget.NewButtonWithIcon("Stop", theme.MediaStopIcon(), func() {
		state.Set("Stopping...")
		stop <- 1
		ctx.Pop()
	})
	stopBtn.Disable()

	startBtn = widget.NewButtonWithIcon("Start", theme.MediaPlayIcon(), func() {
		stopBtn.Enable()
		state.Set("Simulating...")
		go start(game, sim, dayFinished, stop, quit, state, ctx)
		placeholder.Hide()
		navBtn.(*fyne.Container).Objects[0].(*widget.Button).Disable()
		navBtn.(*fyne.Container).Objects[1].(*widget.Button).Disable()
		startBtn.Disable()
	})

	return NewFborder().
		Top(
			NewFborder().
				Left(navBtn).
				Get(
					centered(h1("Simulation")),
				),
		).
		Bottom(
			rightAligned(
				container.NewHBox(
					startBtn,
					stopBtn,
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
							widget.NewIcon(widgets.Icon("newspaper").Resource),
							placeholder,
							widget.NewLabelWithData(news),
						),
					),
					centered(
						container.NewGridWithColumns(2,
							widget.NewIcon(theme.MailComposeIcon()),
							placeholder,
							widget.NewLabelWithData(emails),
						),
					),
				),
			),
		)
}

func start(game *models.Game, sim *services.Simulator, messages chan NotificationCount, stop, quit chan int, state binding.String, ctx *AppContext) {
	for {
		select {
		case <-stop:
			{
				quit <- 0
				return
			}
		default:
			{
				events, simulated := sim.Simulate(1)
				if !simulated {
					state.Set("Stopping...")
					quit <- 1
					checkForEmailDialog(ctx.GetWindow())
					ctx.Pop()
				} else {
					emailsC, newsC := simTriggers(dateStr, news, emails, game, sim, events)
					messages <- NotificationCount{NewEmails: emailsC, NewNews: newsC}
					time.Sleep(time.Duration(conf.SimDaySpeedMs) * time.Millisecond)
				}
			}
		}
	}
}
