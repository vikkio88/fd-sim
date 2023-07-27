package db

import (
	"fdsim/enums"
	"fdsim/models"
)

type RetiredPlayerDto struct {
	Id      string
	Name    string
	Surname string
	Country enums.Country
	Age     int
	Role    models.Role

	Stats       string
	YearRetired int
}

func NewRetiredPlayerFromDto(player PlayerDto, indexedStats map[string]PHistoryDto, year int, leagueId, leagueName string) RetiredPlayerDto {
	stats := "[]"
	if stat, ok := indexedStats[player.Id]; ok {
		stats = stat.Stats
	} else {
		// this means it hasnt played a single match
		if player.Team != nil {
			h := []PHistorySubRow{{
				LeagueId:   leagueId,
				LeagueName: leagueName,
				TeamId:     player.Team.Id,
				TeamName:   player.Team.Name,
				// this is the season starting year-1
				StartYear: year - 1,
			}}
			stats = serialisePHistoryStats(h)
		} else {
			// this is when a retired player had no team
			h := []PHistorySubRow{{
				LeagueId:   leagueId,
				LeagueName: leagueName,
				TeamId:     "",
				TeamName:   "Free Agent",
				// this is the season starting year-1
				StartYear: year - 1,
			}}
			stats = serialisePHistoryStats(h)
		}

	}

	return RetiredPlayerDto{
		Id:      player.Id,
		Name:    player.Name,
		Surname: player.Surname,
		Country: player.Country,
		Age:     player.Age,
		Role:    player.Role,

		Stats:       stats,
		YearRetired: year,
	}
}

func (rp *RetiredPlayerDto) RetiredPlayer(awardsRows []LHistoryDto, trophiesRows []TrophyDto) *models.RetiredPlayer {
	subHistoryRows := unserialisePHistoryStats(rp.Stats)
	history := make([]*models.PHistoryRow, len(subHistoryRows))
	for i, sr := range subHistoryRows {
		history[i] = sr.PHistoryRow()
	}

	awards := []models.Award{}
	if len(awardsRows) > 0 {
		for _, a := range awardsRows {
			awards = append(awards, a.Award(rp.Id))
		}
	}

	trophies := []models.Trophy{}
	if len(trophiesRows) > 0 {
		for _, t := range trophiesRows {
			trophies = append(trophies, t.Trophy())
		}
	}

	return &models.RetiredPlayer{
		Id:          rp.Id,
		Name:        rp.Name,
		Surname:     rp.Surname,
		Country:     rp.Country,
		Age:         rp.Age,
		Role:        rp.Role,
		YearRetired: rp.YearRetired,

		History:  history,
		Awards:   awards,
		Trophies: trophies,
	}
}
