package db

import (
	"fdsim/models"
	"strings"
)

type ResultDto struct {
	MatchId     string
	Match       MatchDto `gorm:"foreignKey:MatchId"`
	GoalsHome   int
	GoalsAway   int
	ScorersHome string
	ScorersAway string
}

func DtoFromResult(r *models.Result, matchId string) *ResultDto {
	return &ResultDto{
		MatchId:     matchId,
		GoalsHome:   r.GoalsHome,
		GoalsAway:   r.GoalsAway,
		ScorersHome: strings.Join(r.ScorersHome, pIdSeparator),
		ScorersAway: strings.Join(r.ScorersAway, pIdSeparator),
	}
}

func ResultsMapPHFromDtos(res []ResultDto) models.ResultsPHMap {
	result := models.ResultsPHMap{}

	for _, r := range res {
		result[r.MatchId] = r.ResultPH()
	}
	return result
}

func (r *ResultDto) ResultPH() *models.ResultPH {
	return &models.ResultPH{
		MatchId:   r.MatchId,
		GoalsHome: r.GoalsHome,
		GoalsAway: r.GoalsAway,
	}
}
func (r *ResultDto) Result() *models.Result {

	return models.NewResult(
		r.GoalsHome, r.GoalsAway,
		getScorers(r.ScorersHome),
		getScorers(r.ScorersAway),
	)
}

func getScorers(scorersJoined string) []string {
	if scorersJoined == "" {
		return []string{}
	}

	return strings.Split(scorersJoined, pIdSeparator)
}
