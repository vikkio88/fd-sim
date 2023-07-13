package ui

import (
	"fdsim/conf"
	"fdsim/db"
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
	ctx.Db = db.NewDb(conf.DbFiles)
	ctx.Version = conf.Version

	//a.Settings().SetTheme(&ui.MuscurdigTheme{})

	return App{
		ctx:         &ctx,
		application: a,
		window:      &w,
		views:       makeRouteMap(&ctx),

		isLogEnabled: isLogEnabled,
	}
}

// TODO: for the next project this might be better as a Container
// or Factory with Cache and a stack to simulate pop push routes
func (a *App) getView() *fyne.Container {
	key := a.ctx.CurrentRoute()
	a.log(fmt.Sprintf("key: %s", key.String()))

	if a.ctx.RouteMode == Pop {
		a.log("POP cached view")
		//TODO: this is crap, I need to make everything fyne.CanvasObject
		return a.ctx.NavStack.GetPopContent().(*fyne.Container)
	}

	if a.ctx.RouteMode == Replace {
		a.log("REPLACE")
		a.ctx.NavStack.ClearAll()
	}

	if content, ok := a.views[key]; ok {
		a.log(
			fmt.Sprintf("view found %s: rendering view anew", key.String()),
		)
		return content()
	} else {
		a.log(fmt.Sprintf("view %s not found", key.String()))
	}

	return a.views[Main]()
}

func (a *App) setView() {
	if a.ctx.RouteMode == Push {
		a.log("caching old view")
		a.ctx.CacheViewOnStack((*a.window).Content())
	}

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
	//TODO: remove this as it is for testing views
	//TODO: check why Main seems to be rendered twice
	initialRoute := Main

	// initialRoute := Test

	return NewAppContext(initialRoute, w)
}
