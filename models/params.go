package models

import "fdsim/enums"

type EventParams struct {
	Country       enums.Country
	LeagueId      string
	LeagueName    string
	LeagueCountry enums.Country
	RoundId       string
	MatchId       string
	TeamId        string
	TeamName      string
	TeamId1       string
	TeamName1     string
	TeamId2       string
	TeamName2     string
	PlayerId      string
	PlayerName    string
	PlayerId1     string
	PlayerName1   string
	PlayerId2     string
	PlayerName2   string
	CoachId       string
	CoachName     string
	Label1        string
	Label2        string
	Label3        string
	Label4        string
	ValueInt      int
	ValueInt1     int
	ValueF        float64
	ValueF1       float64
	BoolFlag      bool
	BoolFlag2     bool
	FdName        string
	IsEmployed    bool
	FdTeamId      string
	FdTeamName    string
}

func EP() EventParams {
	return EventParams{}
}
