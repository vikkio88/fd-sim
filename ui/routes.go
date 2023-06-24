package ui

import (
	"strings"
)

type AppRoute uint8

const (
	main      string = "MAIN"
	setup     string = "SETUP"
	newGame   string = "NEW_GAME"
	loadGame  string = "LOAD_GAME"
	dashboard string = "DASHBOARD"
	email     string = "EMAIL_DETAIL"
	news      string = "NEWS_DETAIL"

	teamDetails   string = "TEAM_DETAILS"
	playerDetails string = "PLAYER_DETAILS"

	league       string = "LEAGUE"
	matchDetails string = "MATCH_DETAILS"

	test string = "TEST"
	quit string = "QUIT"

	invalid string = "INVALID_ROUTE"
)

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
	MatchDetails

	Test
	Quit
)

func getMapping() map[AppRoute]string {
	return map[AppRoute]string{
		Main:      main,
		Setup:     setup,
		NewGame:   newGame,
		LoadGame:  loadGame,
		Dashboard: dashboard,
		Email:     email,
		News:      news,

		TeamDetails:   teamDetails,
		PlayerDetails: playerDetails,

		League:       league,
		MatchDetails: matchDetails,

		Test: test,
		Quit: quit,
	}
}

func getReverseMapping() map[string]AppRoute {
	return map[string]AppRoute{
		main:      Main,
		setup:     Setup,
		newGame:   NewGame,
		loadGame:  LoadGame,
		dashboard: Dashboard,
		email:     Email,
		news:      News,

		teamDetails:   TeamDetails,
		playerDetails: PlayerDetails,

		league:       League,
		matchDetails: MatchDetails,

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
