package ui

import (
	"strings"
)

type AppRoute uint8

const (
	main    string = "MAIN"
	setup   string = "SETUP"
	newGame string = "NEW_GAME"

	teamDetails   string = "TEAM_DETAILS"
	playerDetails string = "PLAYER_DETAILS"

	test string = "TEST"
	quit string = "QUIT"

	invalid string = "INVALID_ROUTE"
)

const (
	Main AppRoute = iota
	Setup
	NewGame

	TeamDetails
	PlayerDetails

	Test
	Quit
)

func getMapping() map[AppRoute]string {
	return map[AppRoute]string{
		Main:    main,
		Setup:   setup,
		NewGame: newGame,

		TeamDetails:   teamDetails,
		PlayerDetails: playerDetails,

		Test: test,
		Quit: quit,
	}
}

func getReverseMapping() map[string]AppRoute {
	return map[string]AppRoute{
		main:    Main,
		setup:   Setup,
		newGame: NewGame,

		teamDetails:   TeamDetails,
		playerDetails: PlayerDetails,

		test: Test,
		quit: Quit,
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

	return invalid
}
