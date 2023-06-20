package ui

import (
	"fdsim/models"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func matchDetailsView(ctx *AppContext) *fyne.Container {
	matchId := ctx.RouteParam.(string)
	match := ctx.Db.LeagueR().GetMatchById(matchId)

	home := makeTeamSide(match.Home,
		// match.LineupHome,
		match.Result,
		true, //fucking hate this
		ctx.PushWithParam)
	away := makeTeamSide(match.Away,
		// match.LineupAway,
		match.Result,
		false, //fucking hate this
		ctx.PushWithParam)
	result := widget.NewLabel("-")
	if match.Result != nil {
		result = widget.NewLabel(match.Result.String())
	}

	return NewFborder().
		Top(
			NewFborder().Left(backButton(ctx)).Get(
				h1(fmt.Sprintf("Round %d", match.RoundIndex+1)),
			),
		).
		Get(
			container.NewGridWithColumns(3,
				home,
				NewFborder().Top(centered(result)).Get(),
				away,
			),
		)
}

// TODO: could add lineup
func makeTeamSide(team *models.Team /* lineup []string,*/, result *models.Result, isHomeTeam bool, navigate func(AppRoute, any)) fyne.CanvasObject {
	// TODO: rewrite HL so you can size it rather than using buttons
	teamBtn := widget.NewButton(team.Name, func() {
		navigate(TeamDetails, team.Id)
	})

	teamBtn.Importance = widget.LowImportance

	content := container.NewMax()

	if result != nil {
		scorers := result.ScorersHome
		if !isHomeTeam {
			scorers = result.ScorersAway
		}
		content.Add(NewFborder().
			// Top(centered(widget.NewLabel("Scorers"))).
			Get(
				widget.NewList(
					func() int {
						return len(scorers)
					},
					func() fyne.CanvasObject {
						return centered(
							widget.NewHyperlink("Unknown", nil),
						)
					},
					func(lii widget.ListItemID, co fyne.CanvasObject) {
						scorerId := scorers[lii]
						scorer, ok := team.Roster.Player(scorerId)
						hl := co.(*fyne.Container).Objects[0].(*widget.Hyperlink)
						if ok {
							hl.SetText(scorer.String())
							hl.OnTapped = func() { navigate(PlayerDetails, scorerId) }
						}
					}),
			))
	}

	return NewFborder().
		Top(
			container.NewMax(
				teamBtn,
			)).
		Get(
			content,
		)
}
