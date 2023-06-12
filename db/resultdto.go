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

func (r *ResultDto) Result() *models.Result {
	return models.NewResult(
		r.GoalsHome, r.GoalsAway,
		strings.Split(r.ScorersHome, pIdSeparator),
		strings.Split(r.ScorersAway, pIdSeparator),
	)
}
