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

	//TODO: Remove Testing Actions
	TestingActionYes
	TestingActionNo

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
		return roundPlayedEvent(params, date)
	case LeagueFinished:
		return leagueFinishedEvent(params, date)

		//TODO: Remove Tests Action
	case TestingActionYes:
		{
			teamId := params.TeamId1
			teamName := params.Label1

			desc := fmt.Sprintf("You Replied YES")
			event := NewEvent(date, desc)

			event.TriggerNews = models.NewNews(
				desc,
				data.GetNewspaper(params.LeagueCountry),
				fmt.Sprintf(
					"You replied YES, team was %s",
					conf.LinkBodyPH,
				),
				date,
				[]models.Link{
					models.NewLink(teamName, enums.TeamDetails, &teamId),
				},
			)
			return event
		}
	case TestingActionNo:
		{
			teamId := params.TeamId1
			teamName := params.Label1

			desc := fmt.Sprintf("You Replied NO")
			event := NewEvent(date, desc)

			event.TriggerNews = models.NewNews(
				desc,
				data.GetNewspaper(params.LeagueCountry),
				fmt.Sprintf(
					"You replied NO, team was %s",
					conf.LinkBodyPH,
				),
				date,
				[]models.Link{
					models.NewLink(teamName, enums.TeamDetails, &teamId),
				},
			)
			return event
		}

	}

	return nil
}