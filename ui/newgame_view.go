package ui

import (
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/viewmodels"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/slices"
)

func NewGameView(ctx *AppContext) *fyne.Container {
	countries := viewmodels.GetAllCountries()
	var ts float64 = 2
	teamsNumber := binding.BindFloat(&ts)
	teams := binding.NewUntypedList()
	var selectedCountry enums.Country

	pholder := container.NewCenter(widget.NewLabel("No teams yet..."))
	teamGenBtn := widget.NewButton("Generate Teams", func() {
		pholder.Hide()
		viewmodels.ClearDataUtList(teams)
		tg := generators.NewTeamGen(time.Now().Unix())
		n, _ := teamsNumber.Get()
		ts := tg.TeamsWithCountry(int(n), selectedCountry)
		for _, t := range ts {
			teams.Append(t)
		}
	})
	teamGenBtn.Disable()
	ctrSelect := widget.NewSelect(countries, func(s string) {
		idx := slices.Index(countries, s)
		selectedCountry = viewmodels.CountryFromIndex(idx)
		teamGenBtn.Enable()
	})
	ctrSelect.PlaceHolder = "Select Country"
	numberSelect := widget.NewSliderWithData(2.0, 20.0, teamsNumber)
	numberSelect.Step = 2.0
	numberText := widget.NewLabelWithData(binding.FloatToStringWithFormat(teamsNumber, "Teams: %.0f"))

	teamLst := widget.NewListWithData(
		teams,
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			team := viewmodels.TeamFromDi(di)
			l := co.(*widget.Label)
			l.SetText(team.String())
		},
	)

	return container.NewBorder(
		centered(widget.NewLabel("New Game")),
		nil,
		nil,
		nil,
		container.NewBorder(
			container.NewBorder(
				nil,
				nil,
				nil,
				teamGenBtn,
				container.NewGridWithColumns(2,
					ctrSelect,
					container.NewVBox(
						centered(numberText),
						numberSelect,
					),
				),
			),
			nil,
			nil,
			nil,
			pholder,
			teamLst,
		),
	)
}
