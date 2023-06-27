package services

import (
	"fdsim/conf"
	"fdsim/data"
	"fdsim/enums"
	"fdsim/models"
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
			models.NewLink(teamName, enums.TeamDetails, &teamId),
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
