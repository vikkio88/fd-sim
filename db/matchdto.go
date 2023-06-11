package db

import "fdsim/models"

type MatchDto struct {
	Id       string
	Away     string
	Home     string
	HomeTeam TeamDto `gorm:"foreignKey:Home"`
	AwayTeam TeamDto `gorm:"foreignKey:Away"`
	Result   *ResultDto
	RoundId  string
}

func DtoFromMatchPH(match models.MPH, roundId string) MatchDto {
	return MatchDto{
		Id:      match.Id,
		Home:    match.Home,
		Away:    match.Away,
		Result:  nil,
		RoundId: roundId,
	}
}
