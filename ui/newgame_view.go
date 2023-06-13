package ui

import (
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/models"
	"fdsim/vm"
	"fdsim/widgets"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/slices"
)

func newGameView(ctx *AppContext) *fyne.Container {
	countries := vm.GetAllCountries()
	var ts float64 = 8
	teamsNumber := binding.BindFloat(&ts)
	var teamsSlice []*models.Team
	teams := binding.NewUntypedList()
	var selectedCountry enums.Country

	pholder := container.NewCenter(widget.NewLabel("No teams yet..."))
	startGame := widget.NewButtonWithIcon("Start", theme.NavigateNextIcon(), func() {})
	startGame.Disable()

	teamGenBtn := widget.NewButton("Generate Teams", func() {
		pholder.Hide()
		ctx.Db.TruncateAll()
		vm.ClearDataUtList(teams)
		tg := generators.NewTeamGen(time.Now().Unix())
		n, _ := teamsNumber.Get()
		ts := tg.TeamsWithCountry(int(n), selectedCountry)
		teamsSlice = ts
		ctx.Db.TeamR().Insert(ts)
		for _, t := range ts {
			teams.Append(t)
		}
		startGame.Enable()
	})
	teamGenBtn.Disable()
	startGame.OnTapped = func() {
		teamGenBtn.Disable()
		// generate and save League
		name := "Serie A 2023/2024"
		league := models.NewLeague(name, teamsSlice)
		ctx.Db.LeagueR().InsertOne(&league)
		// create savegame
		startingDate := time.Date(time.Now().Year(), time.July, 1, 0, 0, 0, 0, time.UTC)
		g := models.NewGame(league.Id, "Test", "Mario", "Rossi", 35, startingDate)
		//

		ctx.Db.GameR().Create(g)
		ctx.NavigateToWithParam(Dashboard, g.Id)
	}

	ctrSelect := widget.NewSelect(countries, func(s string) {
		idx := slices.Index(countries, s)
		selectedCountry = vm.CountryFromIndex(idx)
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
			centered(widget.NewLabel("New Game")),
		).
		Bottom(
			NewFborder().Right(startGame).Get(),
		).
		Get(
			NewFborder().
				Top(inputs).
				Get(teamLst, pholder),
		)

}

func simpleTeamListRow() fyne.CanvasObject {
	return NewFborder().
		Get(
			container.NewMax(
				container.NewGridWithColumns(
					3,
					centered(widget.NewHyperlink("", nil)),
					centered(
						container.NewHBox(
							widgets.Icon("team"),
							widget.NewLabel("Roster"),
						),
					),
					centered(starsFromf64(0)),
				),
			),
		)
}

func makeSimpleTeamRowBind(ctx *AppContext) func(di binding.DataItem, co fyne.CanvasObject) {

	return func(di binding.DataItem, co fyne.CanvasObject) {
		team := vm.TeamFromDi(di)
		c := co.(*fyne.Container)

		ctn := c.Objects[0].(*fyne.Container)
		mx := ctn.Objects[0].(*fyne.Container)
		ctr := mx.Objects[0].(*fyne.Container)
		l := ctr.Objects[0].(*widget.Hyperlink)

		l.SetText(team.String())
		l.OnTapped = func() {
			ctx.PushWithParam(TeamDetails, team.Id)
		}
		mx.Objects[1].(*fyne.Container).Objects[0].(*fyne.Container).Objects[1].(*widget.Label).SetText(fmt.Sprintf("%d", team.Roster.Len()))
		mx.Objects[2].(*fyne.Container).Objects[0].(*widgets.StarRating).SetValues(vm.PercFToStars(team.Roster.AvgSkill()))
	}
}
