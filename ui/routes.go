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
	Profile
	Calendar
	Press
	Chat
	Email
	News
	TeamMgmt

	TeamDetails
	PlayerDetails

	League
	LeagueHistory
	RoundDetails
	MatchDetails

	Simulation

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
		Profile:   e.Profile,
		Calendar:  e.Calendar,
		Press:     e.Press,
		Chat:      e.Chat,
		Email:     e.Email,
		News:      e.News,
		TeamMgmt:  e.TeamMgmt,

		TeamDetails:   e.TeamDetails,
		PlayerDetails: e.PlayerDetails,

		League:        e.League,
		LeagueHistory: e.LeagueHistory,
		RoundDetails:  e.RoundDetails,
		MatchDetails:  e.MatchDetails,

		Simulation: e.Simulation,

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
		e.Profile:   Profile,
		e.Calendar:  Calendar,
		e.Press:     Press,
		e.Chat:      Chat,
		e.Email:     Email,
		e.News:      News,
		e.TeamMgmt:  TeamMgmt,

		e.TeamDetails:   TeamDetails,
		e.PlayerDetails: PlayerDetails,

		e.League:        League,
		e.LeagueHistory: LeagueHistory,
		e.RoundDetails:  RoundDetails,
		e.MatchDetails:  MatchDetails,

		e.Simulation: Simulation,

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
