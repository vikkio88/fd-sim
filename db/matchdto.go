package db

import (
	"fdsim/models"
	"strings"
)

type MatchDto struct {
	Id       string `gorm:"primarykey;size:16"`
	Away     string
	Home     string
	HomeTeam TeamDto    `gorm:"foreignKey:Home"`
	AwayTeam TeamDto    `gorm:"foreignKey:Away"`
	Result   *ResultDto `gorm:"foreignKey:match_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoundId  string
	Round    RoundDto `gorm:"foreignKey:round_id"`

	LineupHome *string
	LineupAway *string
}

func DtoFromMatch(match *models.Match, roundId string) MatchDto {
	var homel, awayl *string = nil, nil
	if match.LineupHome != nil {
		homel1 := strings.Join(match.LineupHome.Ids(), pIdSeparator)
		homel = &homel1
	}

	if match.LineupAway != nil {
		awayl1 := strings.Join(match.LineupAway.Ids(), pIdSeparator)
		awayl = &awayl1
	}
	var result *ResultDto = nil
	if r, ok := match.Result(); ok {
		result = DtoFromResult(r, match.Id)
	}

	return MatchDto{
		Id:         match.Id,
		Home:       match.Home.Id,
		Away:       match.Away.Id,
		RoundId:    roundId,
		LineupHome: homel,
		LineupAway: awayl,
		Result:     result,
	}
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

func (m *MatchDto) MatchComplete() *models.MatchComplete {
	var result *models.Result = nil
	if m.Result != nil {
		result = m.Result.Result()
	}
	res := models.MatchComplete{
		Id:         m.Id,
		Home:       m.HomeTeam.Team(),
		Away:       m.AwayTeam.Team(),
		LineupHome: []string{},
		LineupAway: []string{},
		Result:     result,
		RoundIndex: m.Round.Index,
	}

	if m.LineupHome != nil {
		res.LineupHome = strings.Split(*m.LineupHome, pIdSeparator)
	}

	if m.LineupAway != nil {
		res.LineupAway = strings.Split(*m.LineupAway, pIdSeparator)
	}

	return &res
}

func (m *MatchDto) MatchResult(home, away models.TPH) *models.MatchResult {
	res := models.MatchResult{
		Id:         m.Id,
		Home:       home,
		Away:       away,
		LineupHome: []string{},
		LineupAway: []string{},
		Result:     m.Result.Result(),
	}

	if m.LineupHome != nil {
		res.LineupHome = strings.Split(*m.LineupHome, pIdSeparator)
	}

	if m.LineupAway != nil {
		res.LineupAway = strings.Split(*m.LineupAway, pIdSeparator)
	}

	return &res
}

func (m *MatchDto) MPH() models.MPH {
	return models.MPH{
		Id:   m.Id,
		Home: m.Home,
		Away: m.Away,
	}
}
