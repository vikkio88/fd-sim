package services

import (
	"fdsim/models"
	"time"
)

type Event struct {
	Date         time.Time
	Description  string
	TriggerEmail *models.Email
	TriggerNews  *models.News
}

func NewEvent(date time.Time, description string) *Event {
	return &Event{
		Date:        date,
		Description: description,
	}
}
