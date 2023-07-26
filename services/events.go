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
	SeasonOver

	// TransferMarket
	TransferMarketOpen
	TransferMarketClose

	// Offered Contract to User
	ContractOffer
	ContractAccepted
	// Getting Sacked
	Sacked

	//TODO: Remove Testing Actions
	TestingActionYes
	TestingActionNo

	// Those Events are created by other events and triggered on a certain date
	DbStoredEvent

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

func (ev EventType) Event(date time.Time, params models.EventParams) *Event {
	switch ev {
	case ContractOffer:
		return contractOffered(params, date)
	case ContractAccepted:
		return contractAccepted(params, date)
	case RoundPlayed:
		return roundPlayedEvent(params, date)
	case LeagueFinished:
		return leagueFinishedEvent(params, date)
	case SeasonOver:
		return seasonOverEvent(params, date)
	case TransferMarketOpen:
		return transferMarketOpen(params, date)
	case TransferMarketClose:
		return transferMarketClose(params, date)

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
			teamName := params.Label1

			desc := fmt.Sprintf("You Replied NO")
			event := NewEvent(date, desc)

			event.TriggerEmail = models.NewEmailNoLinks(
				"spam@spam.com",
				"Yeah why not?",
				fmt.Sprintf("Some spam as you said NO to us, %s", teamName),
				date)
			return event
		}

	}

	return nil
}
