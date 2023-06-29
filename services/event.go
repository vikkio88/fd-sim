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
		Description: "Empty",
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
