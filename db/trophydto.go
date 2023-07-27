package db

import "fdsim/models"

type TrophyDto struct {
	PlayerId   string
	LeagueId   string
	LeagueName string
	TeamId     string
	TeamName   string
	Year       int
}

func NewTrophyDto(playerId, leagueId, leagueName, teamId, teamName string, year int) TrophyDto {
	return TrophyDto{
		PlayerId:   playerId,
		LeagueId:   leagueId,
		LeagueName: leagueName,
		TeamId:     teamId,
		TeamName:   teamName,
		Year:       year,
	}
}

func (t *TrophyDto) Trophy() models.Trophy {
	return models.Trophy{
		LeagueId:   t.LeagueId,
		LeagueName: t.LeagueName,
		Team:       models.TPH{Id: t.TeamId, Name: t.TeamName},
		Year:       t.Year,
	}
}
