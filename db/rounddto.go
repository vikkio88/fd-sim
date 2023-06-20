package db

import (
	"fdsim/models"
	"time"
)

type RoundDto struct {
	Id        string `gorm:"primarykey;size:16"`
	Index     int
	Matches   []MatchDto `gorm:"foreignKey:round_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WasPlayed bool
	Date      time.Time

	LeagueId string
}

func DtoFromRound(r *models.Round, leagueId string) RoundDto {
	//build matches
	ms := make([]MatchDto, len(r.Matches))
	for i, m := range r.Matches {
		ms[i] = DtoFromMatch(m, r.Id)
	}

	return RoundDto{
		Id:        r.Id,
		Index:     r.Index,
		LeagueId:  leagueId,
		Matches:   ms,
		Date:      r.Date,
		WasPlayed: r.WasPlayed,
	}
}
func DtoFromRoundPH(r models.RPH, leagueId string) RoundDto {
	//build matches
	ms := make([]MatchDto, len(r.Matches))
	for i, m := range r.Matches {
		ms[i] = DtoFromMatchPH(m, r.Id)
	}
	return RoundDto{
		Id:       r.Id,
		Index:    r.Index,
		LeagueId: leagueId,
		Date:     r.Date,
		Matches:  ms,
	}
}

func (r *RoundDto) Round(teams map[string]*models.Team) *models.RoundResult {
	matches := make([]*models.MatchResult, len(r.Matches))
	for i, m := range r.Matches {
		matches[i] = m.MatchResult(teams[m.Home].PH(), teams[m.Away].PH())
	}

	return &models.RoundResult{
		Id:        r.Id,
		Index:     r.Index,
		Matches:   matches,
		WasPlayed: r.WasPlayed,
	}
}

func (r *RoundDto) RoundPH() models.RPH {
	ms := make([]models.MPH, len(r.Matches))
	for i, m := range r.Matches {
		ms[i] = m.MPH()
	}

	return models.RPH{
		Id:      r.Id,
		Index:   r.Index,
		Date:    r.Date,
		Matches: ms,
	}
}

func RoundsPHFromDto(rdtos []RoundDto) []models.RPH {
	rds := make([]models.RPH, len(rdtos))
	for i, r := range rdtos {
		rds[i] = r.RoundPH()
	}
	return rds
}
