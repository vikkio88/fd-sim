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

const (
	a_day = time.Duration(1) * time.Hour * 24
)

func leagueFinishedEvent(params models.EventParams, date time.Time) *Event {
	leagueId := params.LeagueId
	leagueName := params.LeagueName
	teamId := params.TeamId
	teamName := params.TeamName
	event := NewEvent(date, fmt.Sprintf("%s Finished", leagueName))
	title := fmt.Sprintf("%s won %s!", teamName, leagueName)

	event.TriggerNews = models.NewNews(
		title,
		data.GetNewspaper(params.Country),
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

func roundPlayedEvent(params models.EventParams, date time.Time) *Event {
	roundIndex := params.Label1
	roundId := params.RoundId
	leagueId := params.LeagueId
	leagueName := params.LeagueName
	desc := fmt.Sprintf("%s - Round %s played", leagueName, roundIndex)
	event := NewEvent(date, desc)

	event.TriggerNews = models.NewNews(
		desc,
		data.GetNewspaper(params.Country),
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

func contractAccepted(params models.EventParams, date time.Time) *Event {
	teamId := params.TeamId
	teamName := params.TeamName
	ycontract := params.ValueInt
	fdName := params.FdName
	money := utils.NewEurosFromF(params.ValueF)

	title := fmt.Sprintf("%s contract accepted", teamName)

	event := NewEvent(date, title)
	event.TriggerEmail = models.NewEmail(
		emailAddrFromTeamName(teamName),
		fmt.Sprintf("Welcome to %s", teamName),
		fmt.Sprintf(
			"Thanks Mr %s for joining us, we are delighted to have you on board."+
				"Please check our info here: %s CEO of %s",
			fdName,
			conf.LinkBodyPH,
			teamName,
		),
		date,
		[]models.Link{
			teamLink(teamName, teamId),
		},
	)

	event.TriggerNews = models.NewNews(
		fmt.Sprintf("%s hired from %s as Football Director", fdName, teamName),
		data.GetNewspaper(params.Country),
		fmt.Sprintf(
			"A new football director, %s, got hired today from %s. Seems like he signed a %d year(s) contract. %s",
			fdName, teamName, ycontract, conf.LinkBodyPH,
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
		game.SetTeamContract(ycontract, money, &models.TPH{Id: teamId, Name: teamName})
	}

	return event
}

func contractOffered(params models.EventParams, date time.Time) *Event {
	teamId := params.TeamId
	teamName := params.TeamName
	money := utils.NewEurosFromF(params.ValueF)
	years := params.ValueInt

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
			params,
		),
	)
	event.TriggerEmail.SetExpiry(date.Add(2 * a_day))

	event.TriggerFlags = func(f models.Flags) models.Flags {
		f.HasAContractOffer = true
		return f
	}

	return event
}

func transferMarketOpen(params models.EventParams, date time.Time) *Event {
	event := NewEvent(date, "market open")

	session := "Winter"
	// is Summer
	if params.BoolFlag {
		session = "Summer"
	}

	event.TriggerNews = models.NewNews(
		"Transfer Market is open!",
		data.GetNewspaper(params.Country),
		fmt.Sprintf(
			"Today the transfer market for the %s session officially opened!\nIt will stay open until %s",
			session,
			params.Label2,
		),
		date,
		[]models.Link{},
	)

	return event
}
func transferMarketClose(params models.EventParams, date time.Time) *Event {
	event := NewEvent(date, "market closed")

	session := "Winter"
	// is Summer
	if params.BoolFlag {
		session = "Summer"
	}

	event.TriggerNews = models.NewNews(
		"Transfer Market is now closed.",
		data.GetNewspaper(params.Country),
		fmt.Sprintf(
			"Today the transfer market for the %s session is closed!",
			session,
		),
		date,
		[]models.Link{},
	)

	return event
}
