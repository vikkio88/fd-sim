package db

import (
	"encoding/json"
	"fdsim/models"
	"fmt"
)

type LHistoryDto struct {
	Id          string
	Name        string
	Podium      string
	BestScorers string
	Mvp         string

	BestScorerId string
	MvpId        string
}

func (l *LHistoryDto) Award(playerId string) models.Award {
	scorer := l.BestScorerId == playerId
	mvp := l.MvpId == playerId
	lh := l.LeagueHistory()
	goals := 0
	played := 0
	score := 0.0

	var team models.TPH

	if scorer {
		topScorer := lh.BestScorers[0]
		goals = topScorer.Goals
		played = topScorer.Played
		team = *topScorer.Team
	}

	if mvp {
		score = lh.Mvp.Score
		played = lh.Mvp.Played
		team = *lh.Mvp.Team
	}

	return models.Award{
		LeagueId:   l.Id,
		LeagueName: l.Name,

		Scorer: scorer,
		Mvp:    mvp,

		Goals:  goals,
		Score:  score,
		Played: played,

		Team: team,
	}
}

func (l *LHistoryDto) LeagueHistory() *models.LeagueHistory {
	return &models.LeagueHistory{
		Id:          l.Id,
		Name:        l.Name,
		Podium:      unserialisePodium(l.Podium),
		BestScorers: unserialiseBestScorers(l.BestScorers),
		Mvp:         unserialiseMvp(l.Mvp),
	}
}

func unserialiseBestScorers(s string) []*models.PlayerHistorical {
	if s == "" {
		return []*models.PlayerHistorical{}
	}

	var result []*models.PlayerHistorical
	data := s
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Println("error while loading SCORERS", err)
		return []*models.PlayerHistorical{}
	}

	return result
}

func unserialiseMvp(s string) *models.PlayerHistorical {
	if s == "" {
		return &models.PlayerHistorical{}
	}

	var result models.PlayerHistorical
	data := s
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Println("error while loading MVP", err)
		return &models.PlayerHistorical{}
	}

	return &result
}

func unserialisePodium(s string) []*models.TPHRow {
	if s == "" {
		return []*models.TPHRow{}
	}

	var result []*models.TPHRow
	data := s
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Println("error while loading PODIUM", err)
		return []*models.TPHRow{}
	}

	return result
}

func NewLHistoryDtoFromLeague(league *models.League, mvpStat StatRowDto, scorersStat []StatRowDto) LHistoryDto {
	scorers := StatRowsPhFromDtos(scorersStat)
	mvp := StatRowsPhFromDto(mvpStat)
	lh := models.NewLeagueHistory(league, mvp, scorers)

	return DtoFromLeagueHistory(lh)
}

func DtoFromLeagueHistory(l *models.LeagueHistory) LHistoryDto {
	return LHistoryDto{
		Id:          l.Id,
		Name:        l.Name,
		Podium:      serialisePodium(l.Podium),
		BestScorers: serialiseBestScorers(l.BestScorers),
		Mvp:         serialiseMvp(l.Mvp),

		MvpId:        l.Mvp.Id,
		BestScorerId: l.BestScorers[0].Id,
	}
}

func serialiseMvp(mvp *models.PlayerHistorical) string {
	if mvp == nil {
		return ""
	}

	var result string
	data, _ := json.Marshal(mvp)
	result = string(data)

	return result
}

func serialiseBestScorers(scorers []*models.PlayerHistorical) string {
	if len(scorers) < 1 {
		return "[]"
	}

	var result string
	data, _ := json.Marshal(scorers)
	result = string(data)

	return result
}

func serialisePodium(podium []*models.TPHRow) string {
	if len(podium) < 1 {
		return "[]"
	}

	var result string
	data, _ := json.Marshal(podium)
	result = string(data)

	return result
}
