package ui

import (
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/viewmodels"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/slices"
)

func newGameView(ctx *AppContext) *fyne.Container {
	countries := viewmodels.GetAllCountries()
	var ts float64 = 2
	teamsNumber := binding.BindFloat(&ts)
	teams := binding.NewUntypedList()
	var selectedCountry enums.Country

	pholder := container.NewCenter(widget.NewLabel("No teams yet..."))
	teamGenBtn := widget.NewButton("Generate Teams", func() {
		pholder.Hide()
		ctx.Db.TruncateAll()
		viewmodels.ClearDataUtList(teams)
		tg := generators.NewTeamGen(time.Now().Unix())
		n, _ := teamsNumber.Get()
		ts := tg.TeamsWithCountry(int(n), selectedCountry)
		ctx.Db.TeamR().Insert(ts)
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
		simpleTeamListRow,
		makeSimpleTeamRowBind(ctx),
	)

	inputs := NewFborder().Right(teamGenBtn).Get(
		container.NewGridWithColumns(2,
			ctrSelect,
			container.NewVBox(
				centered(numberText),
				numberSelect,
			),
		),
	)
	return NewFborder().
		Top(
			centered(widget.NewLabel("New Game"))).
		Get(
			NewFborder().
				Top(inputs).
				Get(teamLst, pholder),
		)

}

func simpleTeamListRow() fyne.CanvasObject {
	return NewFborder().Right(widget.NewButtonWithIcon("", theme.ZoomInIcon(), func() {})).Get(widget.NewLabel(""))
}

func makeSimpleTeamRowBind(ctx *AppContext) func(di binding.DataItem, co fyne.CanvasObject) {

	return func(di binding.DataItem, co fyne.CanvasObject) {
		team := viewmodels.TeamFromDi(di)
		c := co.(*fyne.Container)
		l := c.Objects[0].(*widget.Label)
		b := c.Objects[1].(*widget.Button)
		l.SetText(team.String())
		b.OnTapped = func() {
			ctx.PushWithParam(TeamDetails, team.Id)
		}
	}
}
