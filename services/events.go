package services

import (
	"fdsim/models"
	"fmt"
	"strings"
	"time"
)

type EventType uint8

const (
	// Objects[] { index, id of the round }
	RoundPlayed EventType = iota
	// Objects[] { leagueName, leagueid, teamId winner }
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

func (a EventType) Event(date time.Time, objects []string) *Event {
	switch a {
	case RoundPlayed:
		{
			desc := fmt.Sprintf("Round %s played", objects[0])
			event := NewEvent(date, desc)

			event.TriggerNews = models.NewNews(desc, "Sportsweek", desc, date, []models.Link{})
			return event
		}
	case LeagueFinished:
		{
			desc := "League Finished"
			event := NewEvent(date, desc)

			event.TriggerNews = models.NewNews(desc, "Sportsweek", desc, date,
				[]models.Link{
					models.NewLink(objects[0], "LEAGUE", &objects[1]),
				},
			)
			return event

		}

	}

	return nil
}
