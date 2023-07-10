package ui

import (
	"fdsim/db"
	"fdsim/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

type ViewMode uint8

const (
	Replace ViewMode = iota
	Push
	Pop
)

type AppContext struct {
	Version string
	//TODO: maybe those 3 props can be handled by Navstack
	Route      binding.String
	RouteParam any
	RouteMode  ViewMode
	//

	Db db.IDb

	gameState *models.Game

	NavStack *NavStack

	w fyne.Window
}

func NewAppContext(initialRoute AppRoute, window fyne.Window) AppContext {
	route := initialRoute.String()
	return AppContext{
		Route:    binding.BindString(&route),
		NavStack: NewNavStack(),
		w:        window,
	}
}

func (c *AppContext) GetClipboard() fyne.Clipboard {
	return c.w.Clipboard()
}

func (c *AppContext) GetWindow() fyne.Window {
	return c.w
}

func (c *AppContext) OnRouteChange(callback func()) {
	c.Route.AddListener(binding.NewDataListener(callback))
}

func (c *AppContext) CurrentRoute() AppRoute {
	r, _ := c.Route.Get()
	return RouteFromString(r)
}

func (c *AppContext) NavigateTo(route AppRoute) {
	c.RouteParam = nil
	c.RouteMode = Replace
	c.Route.Set(route.String())
}

func (c *AppContext) NavigateToWithParam(route AppRoute, param any) {
	c.RouteParam = param
	c.RouteMode = Replace
	c.Route.Set(route.String())
}

func (c *AppContext) PushWithParam(route AppRoute, param any) {
	c.NavStack.Push(NewNavStackItem(c.CurrentRoute(), c.RouteParam))
	c.RouteParam = param
	c.RouteMode = Push
	c.Route.Set(route.String())
}

func (c *AppContext) Push(route AppRoute) {
	c.NavStack.Push(NewNavStackItem(c.CurrentRoute(), c.RouteParam))
	c.RouteMode = Push
	c.RouteParam = nil
	// TODO: maybe change this to Route itself as it breaks if you forget to map
	c.Route.Set(route.String())
}

func (c *AppContext) Pop() {
	nsi, ok := c.NavStack.Pop()
	if !ok {
		//TODO: instead of panic maybe should revert to a base view?
		panic("Trying to pop even tho you have no views on the stack")
	}

	c.RouteMode = Pop
	c.RouteParam = nsi.routeParam
	c.Route.Set(nsi.route.String())
}

func (c *AppContext) CacheViewOnStack(content fyne.CanvasObject) {
	//TODO: maybe return an error or something if this Peek fails
	i, ok := c.NavStack.Peek()
	if !ok {
		return
	}

	i.SetContent(content)
}

// GameState utils
func (c *AppContext) InitGameState(gameId string) *models.Game {
	if c.gameState == nil || c.gameState.Id != gameId {
		c.gameState = c.Db.GameR().ById(gameId)
		c.gameState.OnEmployed = func() {
			fdTeamId = c.gameState.Team.Id
		}

		c.gameState.OnUnEmployed = func() {
			fdTeamId = ""
		}

		if c.gameState.Team != nil {
			fdTeamId = c.gameState.Team.Id
		}
	}

	return c.gameState
}

// returns the save game and a bool telling you whether it has been init
func (c *AppContext) GetGameState() (*models.Game, bool) {
	if c.gameState == nil {
		return nil, false
	}
	return c.gameState, true
}

type NavigateWithParamFunc func(AppRoute, any)
