package services

import (
	"fdsim/conf"
	"fdsim/data"
	"fdsim/db"
	"fdsim/enums"
	"fdsim/models"
	"fdsim/utils"
	"fmt"
	"time"
)

func leagueFinishedEvent(params EventParams, date time.Time) *Event {
	leagueId := params.LeagueId
	teamId := params.TeamId1
	leagueName := params.LeagueName
	teamName := params.Label1
	event := NewEvent(date, fmt.Sprintf("%s Finished", leagueName))
	title := fmt.Sprintf("%s won %s!", teamName, leagueName)

	event.TriggerNews = models.NewNews(
		title,
		data.GetNewspaper(params.LeagueCountry),
		fmt.Sprintf(
			"Today the %s League officially finished, and the winner was %s."+
				"\nFinal Table: %s Winner %s",
			leagueName, teamName, conf.LinkBodyPH, conf.LinkBodyPH,
		),
		date,
		[]models.Link{
			models.NewLink(leagueName, enums.League, &leagueId),
			teamLink(teamName, teamId),
		},
	)
	return event
}

func roundPlayedEvent(params EventParams, date time.Time) *Event {
	roundIndex := params.Label1
	roundId := params.RoundId
	leagueId := params.LeagueId
	leagueName := params.LeagueName
	desc := fmt.Sprintf("%s - Round %s played", leagueName, roundIndex)
	event := NewEvent(date, desc)

	event.TriggerNews = models.NewNews(
		desc,
		data.GetNewspaper(params.LeagueCountry),
		fmt.Sprintf(
			"The %sth round of %s  was played today %s, "+
				"Here you can see the updated League table:"+
				" %s Here you can see the round results %s",
			roundIndex,
			leagueName,
			date.Format(conf.DateFormatGame),
			conf.LinkBodyPH,
			conf.LinkBodyPH,
		),
		date,
		[]models.Link{
			models.NewLink(fmt.Sprintf("%s Table", leagueName), enums.League, &leagueId),
			models.NewLink("Round Results", enums.RoundDetails, &roundId),
		},
	)
	return event
}

func contractAccepted(params EventParams, date time.Time) *Event {
	teamId := params.TeamId1
	teamName := params.Label1
	ycontract := params.valueInt
	money := utils.NewEurosFromF(params.valueF)

	title := fmt.Sprintf("%s contract accepted", teamName)

	event := NewEvent(date, title)
	event.TriggerEmail = models.NewEmail(
		emailAddrFromTeamName(teamName),
		fmt.Sprintf("Welcome to %s", teamName),
		fmt.Sprintf(
			"Thanks for joining us, we are delighted to have you on board."+
				"Please check our info here: %s CEO of %s",
			conf.LinkBodyPH,
			teamName,
		),
		date,
		[]models.Link{
			teamLink(teamName, teamId),
		},
	)

	event.TriggerFlags = func(f models.Flags) models.Flags {
		f.HasAContractOffer = false
		return f
	}

	event.TriggerChanges = func(game *models.Game, db db.IDb) {
		game.YContract = uint8(ycontract)
		game.Wage = money
		game.Team = &models.TPH{Id: teamId, Name: teamName}
	}

	return event
}

func contractOffered(params EventParams, date time.Time) *Event {
	teamId := params.TeamId1
	teamName := params.Label1
	money := utils.NewEurosFromF(params.valueF)
	years := params.valueInt

	title := fmt.Sprintf("%s contract offer", teamName)

	event := NewEvent(date, title)
	event.TriggerEmail = models.NewEmailWithAction(
		emailAddrFromTeamName(teamName),
		title,
		fmt.Sprintf(
			"We are willing to offer you %s per year for a lentgth of %d year(s)."+
				"Please consider us for your next job LINK",
			money.StringKMB(),
			years,
		),
		date,
		[]models.Link{
			teamLink(teamName, teamId),
		},
		MakeActionableFromType(
			models.ActionRespondContract,
			date,
			ActionParameter{
				TeamId:   &teamId,
				Label:    &teamName,
				ValueInt: &years,
				ValueF:   &params.valueF,
			},
		),
	)

	event.TriggerFlags = func(f models.Flags) models.Flags {
		f.HasAContractOffer = true
		return f
	}

	return event

}
