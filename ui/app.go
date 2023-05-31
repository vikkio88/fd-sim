package ui

import (
	"fdsim/conf"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type App struct {
	isLogEnabled bool
	ctx          *AppContext
	views        map[AppRoute]func() *fyne.Container
	application  fyne.App
	window       *fyne.Window
}

func NewApp() App {
	a := app.NewWithID(conf.AppId)
	w := a.NewWindow(conf.WindowTitle)
	w.Resize(fyne.NewSize(
		conf.WindowWidth,
		conf.WindowHeight,
	))
	w.SetFixedSize(conf.WindowFixed)
	isLogEnabled := conf.EnableConsoleLog

	ctx := setupContext(w)
	ctx.Version = conf.Version

	//a.Settings().SetTheme(&ui.MuscurdigTheme{})

	return App{
		ctx:          &ctx,
		isLogEnabled: isLogEnabled,
		application:  a,
		window:       &w,
		views: map[AppRoute]func() *fyne.Container{
			Main: func() *fyne.Container { return mainView(&ctx) },
		},
	}
}

// TODO: for the next project this might be better as a Container
// or Factory with Cache and a stack to simulate pop push routes
func (a *App) getView() *fyne.Container {
	key := a.ctx.CurrentRoute()

	if content, ok := a.views[key]; ok {
		return content()
	}

	return a.views[Main]()
}
func (a *App) setView() {
	(*a.window).SetContent(a.getView())
}

func (a *App) log(msg string) {
	if a.isLogEnabled {
		fmt.Printf("%s - %s\n", time.Now().Format("15:04:05"), msg)
	}
}

func (a *App) Run() {
	a.ctx.OnRouteChange(func() {
		val := a.ctx.CurrentRoute()
		a.log(fmt.Sprintf("route state changed %s", val))
		if val == Quit {
			a.application.Quit()
		}

		a.setView()
	})
	a.setView()
	(*a.window).ShowAndRun()

	a.log("exiting...")
}

func (a *App) Cleanup() {
	a.log("Running cleanup")
	a.log("cleanup finished")
}

func setupContext(w fyne.Window) AppContext {
	initialRoute := Main

	return NewAppContext(initialRoute, w)
}
