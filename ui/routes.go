package ui

import (
	e "fdsim/enums"
	"strings"
)

type AppRoute uint8

const (
	Main AppRoute = iota
	Setup
	NewGame
	LoadGame
	Dashboard
	Email
	News

	TeamDetails
	PlayerDetails

	League
	RoundDetails
	MatchDetails

	Test
	Quit
)

func getMapping() map[AppRoute]string {
	return map[AppRoute]string{
		Main:      e.Main,
		Setup:     e.Setup,
		NewGame:   e.NewGame,
		LoadGame:  e.LoadGame,
		Dashboard: e.Dashboard,
		Email:     e.Email,
		News:      e.News,

		TeamDetails:   e.TeamDetails,
		PlayerDetails: e.PlayerDetails,

		League:       e.League,
		RoundDetails: e.RoundDetails,
		MatchDetails: e.MatchDetails,

		Test: e.Test,
		Quit: e.Quit,
	}
}

func getReverseMapping() map[string]AppRoute {
	return map[string]AppRoute{
		e.Main:      Main,
		e.Setup:     Setup,
		e.NewGame:   NewGame,
		e.LoadGame:  LoadGame,
		e.Dashboard: Dashboard,
		e.Email:     Email,
		e.News:      News,

		e.TeamDetails:   TeamDetails,
		e.PlayerDetails: PlayerDetails,

		e.League:       League,
		e.RoundDetails: RoundDetails,
		e.MatchDetails: MatchDetails,

		e.Test: Test,
		e.Quit: Quit,
	}
}

func RouteFromString(route string) AppRoute {
	route = strings.ToUpper(route)
	mapping := getReverseMapping()
	if val, ok := mapping[route]; ok {
		return val
	}

	return Main
}

func (a AppRoute) String() string {
	mapping := getMapping()
	if val, ok := mapping[a]; ok {
		return val
	}

	return e.Invalid
}
