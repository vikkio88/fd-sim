package db

import (
	"fdsim/enums"
	"time"
)

type DbEventDto struct {
	Id   string `gorm:"primarykey;size:16"`
	Type DbEventType
	// This is to get newspaper titles
	Country enums.Country
	Payload string

	TriggerDate time.Time
}

type DbEventType = uint16

const (
	DbEvPlRetiredFdTeam DbEventType = iota
	DbEvYoungJoinedFdTeam
)

func NewDbEventDto(kind DbEventType, country enums.Country, payload string, triggerDate time.Time) DbEventDto {
	return DbEventDto{
		Id:          IdGenerator("dbev"),
		Type:        kind,
		Country:     country,
		Payload:     payload,
		TriggerDate: triggerDate,
	}
}
