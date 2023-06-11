package db

import "fdsim/models"

type RoundDto struct {
	Id      string
	Index   int
	Matches []MatchDto

	LeagueId string
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
	}
}
