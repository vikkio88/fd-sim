package ui

import "fyne.io/fyne/v2"

func makeRouteMap(ctx *AppContext) map[AppRoute]func() *fyne.Container {
	return map[AppRoute]func() *fyne.Container{
		Main:     func() *fyne.Container { return mainView(ctx) },
		Setup:    func() *fyne.Container { return setupView(ctx) },
		NewGame:  func() *fyne.Container { return newGameView(ctx) },
		LoadGame: func() *fyne.Container { return loadGameView(ctx) },

		Dashboard: func() *fyne.Container { return dashboardView(ctx) },
		Profile:   func() *fyne.Container { return profileView(ctx) },
		Calendar:  func() *fyne.Container { return calendarView(ctx) },
		Press:     func() *fyne.Container { return pressView(ctx) },
		Chat:      func() *fyne.Container { return chatView(ctx) },
		Email:     func() *fyne.Container { return notificationView(ctx, Email) },
		News:      func() *fyne.Container { return notificationView(ctx, News) },
		TeamMgmt:  func() *fyne.Container { return teamMgmtView(ctx) },

		TeamDetails:   func() *fyne.Container { return teamDetailsView(ctx) },
		PlayerDetails: func() *fyne.Container { return playerDetailsView(ctx) },
		League:        func() *fyne.Container { return leagueView(ctx) },
		LeagueHistory: func() *fyne.Container { return leaguehistoryView(ctx) },
		RoundDetails:  func() *fyne.Container { return roundDetailsView(ctx) },
		MatchDetails:  func() *fyne.Container { return matchDetailsView(ctx) },

		Simulation: func() *fyne.Container { return simulationView(ctx) },

		//TEST ROUTE
		Test: func() *fyne.Container { return testView(ctx) },
	}
}
