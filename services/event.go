package services

import (
	"fdsim/db"
	"fdsim/models"
	"time"
)

type Event struct {
	Date           time.Time
	Description    string
	TriggerEmail   *models.Email
	TriggerNews    *models.News
	TriggerFlags   func(models.Flags) models.Flags
	TriggerChanges func(game *models.Game, db db.IDb)
}

func NewEmptyEvent() *Event {
	return &Event{
		Description:    "Empty",
		TriggerFlags:   func(f models.Flags) models.Flags { return f },
		TriggerChanges: func(game *models.Game, db db.IDb) {},
	}
}
func NewEvent(date time.Time, description string) *Event {
	return &Event{
		Date:        date,
		Description: description,
		// basic trigger flag returns the same flags
		TriggerFlags: func(f models.Flags) models.Flags { return f },
		// complex triggers that changes the game state
		TriggerChanges: func(game *models.Game, db db.IDb) {},
	}
}

// DbTriggered events have json serialised payload
func parseDbEvents(dbEvents []db.DbEventDto) []*Event {
	result := make([]*Event, len(dbEvents))

	for i, dbe := range dbEvents {
		ev := getEventFromDbEvent(dbe)
		if ev != nil {
			result[i] = ev
		}
	}
	return result
}
