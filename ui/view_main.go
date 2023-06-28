package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func mainView(ctx *AppContext) *fyne.Container {
	return NewFborder().
		Top(centered(h1("FDSim"))).
		Get(
			container.NewCenter(
				container.NewVBox(
					// widget.NewButton("Test Simulation Page",
					// 	func() {
					// 		ctx.NavigateTo(Simulation)
					// 	},
					// ),
					widget.NewButton("New Game",
						func() {
							ctx.NavigateTo(NewGame)
						},
					),
					widget.NewButtonWithIcon("Load Game", theme.LoginIcon(),
						func() {
							ctx.Push(LoadGame)
						},
					),
					widget.NewButtonWithIcon("Setup", theme.SettingsIcon(),
						func() {
							ctx.Push(Setup)
						},
					),
					widget.NewSeparator(),
					widget.NewButtonWithIcon("Quit", theme.LogoutIcon(),
						func() {
							ctx.NavigateTo(Quit)
						},
					),
				),
			),
		)
}
