package services

import (
	"fdsim/conf"
	"fdsim/data"
	"fdsim/enums"
	"fdsim/models"
	"fmt"
	"strings"
	"time"
)

type EventType uint8

const (
	// Needs LeagueId, RoundId and RoundIndex
	RoundPlayed EventType = iota
	// Needs LeagueId and LeagueName, TeamId and TeamName for Winner
	LeagueFinished

	Null
)

const (
	roundPlayed    string = "RoundPlayed"
	leagueFinished string = "LeagueFinished"

	null string = "null_event"
)

func getMapping() map[EventType]string {
	return map[EventType]string{
		RoundPlayed:    roundPlayed,
		LeagueFinished: leagueFinished,
	}
}

func getReverseMapping() map[string]EventType {
	return map[string]EventType{
		roundPlayed:    RoundPlayed,
		leagueFinished: LeagueFinished,
	}
}

func EventTypeFromString(route string) EventType {
	route = strings.ToUpper(route)
	mapping := getReverseMapping()
	if val, ok := mapping[route]; ok {
		return val
	}

	return Null
}

func (a EventType) String() string {
	mapping := getMapping()
	if val, ok := mapping[a]; ok {
		return val
	}

	return null
}

type EventParams struct {
	LeagueId      string
	LeagueName    string
	LeagueCountry enums.Country
	RoundId       string
	MatchId       string
	TeamId1       string
	TeamId2       string
	PlayerId      string
	CoachId       string
	Label1        string
	Label2        string
	Label3        string
	Label4        string
}

func (a EventType) Event(date time.Time, params EventParams) *Event {
	switch a {
	case RoundPlayed:
		{
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
	case LeagueFinished:
		{
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

	}

	return nil
}
