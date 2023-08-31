package db

import (
	"encoding/json"
	"fdsim/enums"
	"fdsim/models"
	"time"
)

type DbEventDto struct {
	Id   string `gorm:"primarykey;size:16"`
	Type DbEventType
	// This is to get newspaper titles
	Country     enums.Country
	Payload     string
	EventParams string

	TriggerDate time.Time
}

type DbEventType = uint16

const (
	DbEvPlRetiredFdTeam DbEventType = iota
	DbEvYoungJoinedFdTeam
	DbEvPlayerLeftFreeAgent
	DbEvPlayersSkillChanged
	DbEvIndividualAwards

	DbEvTeamAcceptedOffer
	DbEvTeamRefusedOffer

	DbEvPlayerAcceptedContract
	DbEvPlayerRefusedContract
)

func NewDbEventDto(kind DbEventType, country enums.Country, payload string, evParams models.EventParams, triggerDate time.Time) DbEventDto {
	evParamBytes, _ := json.Marshal(evParams)

	return DbEventDto{
		Id:          IdGenerator("dbev"),
		Type:        kind,
		Country:     country,
		Payload:     payload,
		EventParams: string(evParamBytes),

		TriggerDate: triggerDate,
	}
}
