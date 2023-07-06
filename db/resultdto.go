package db

import (
	"encoding/json"
	"fdsim/models"
	"fmt"
	"strings"
)

type ResultDto struct {
	MatchId     string
	Match       MatchDto `gorm:"foreignKey:MatchId"`
	GoalsHome   int
	GoalsAway   int
	ScorersHome string
	ScorersAway string
	ScoreHome   string
	ScoreAway   string
}

func DtoFromResult(r *models.Result, matchId string) *ResultDto {
	return &ResultDto{
		MatchId:     matchId,
		GoalsHome:   r.GoalsHome,
		GoalsAway:   r.GoalsAway,
		ScorersHome: strings.Join(r.ScorersHome, pIdSeparator),
		ScorersAway: strings.Join(r.ScorersAway, pIdSeparator),
		ScoreHome:   serialiseScoreMap(r.ScoreHome),
		ScoreAway:   serialiseScoreMap(r.ScoreAway),
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
		unserialiseScoreMap(r.ScoreHome),
		unserialiseScoreMap(r.ScoreAway),
	)
}

func serialiseScoreMap(playerScoreMap models.PlayerScoreMap) string {
	var result string
	data, _ := json.Marshal(playerScoreMap)
	result = string(data)

	return result
}

func unserialiseScoreMap(s string) models.PlayerScoreMap {
	var res models.PlayerScoreMap

	err := json.Unmarshal([]byte(s), &res)
	if err != nil {
		fmt.Println("Error while getting flags")
		return models.PlayerScoreMap{}
	}

	return res
}

func getScorers(scorersJoined string) []string {
	if scorersJoined == "" {
		return []string{}
	}

	return strings.Split(scorersJoined, pIdSeparator)
}
