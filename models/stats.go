package models

import "time"

type StatRowPH struct {
	Index int
	StatRow
}
type StatRow struct {
	PlayerId string
	Player   *PNPH
	TeamId   string
	Team     *TPH
	LeagueId string
	Played   int
	Goals    int
	Score    float64
}

func (sr *StatRow) Merge(o *StatRow) {
	sr.Played += o.Played
	sr.Goals += o.Goals
	sr.Score += o.Score
}

func NewStatRow(player, team, league string, played, goals int, score float64) *StatRow {
	return &StatRow{
		PlayerId: player,
		TeamId:   team,
		LeagueId: league,
		Played:   played,
		Goals:    goals,
		Score:    score,

		Player: nil,
		Team:   nil,
	}
}

func NewStatRowBase(player, team, league string) *StatRow {
	return &StatRow{
		PlayerId: player,
		TeamId:   team,
		LeagueId: league,
		Played:   1,

		Player: nil,
		Team:   nil,
	}
}

type StatsMap map[string]*StatRow

func StatsFromRoundResult(round *Round, leagueId string) StatsMap {
	// // the length is at least 11 x 2 pear each match
	// stats := make([]*StatRow, 2*11*len(round.Matches))
	stats := map[string]*StatRow{}
	for _, m := range round.Matches {
		r, ok := m.Result()
		if !ok {
			continue
		}

		for _, pId := range m.LineupHome.Ids() {
			stats[pId] = NewStatRowBase(pId, m.Home.Id, leagueId)
			if score, ok := r.ScoreHome[pId]; ok {
				stats[pId].Score = score
			} else {
				panic("THIS SHOULD NEVER HAPPEN")
			}
		}

		for _, pId := range m.LineupAway.Ids() {
			stats[pId] = NewStatRowBase(pId, m.Away.Id, leagueId)
			if score, ok := r.ScoreAway[pId]; ok {
				stats[pId].Score = score
			} else {
				panic("THIS SHOULD NEVER HAPPEN")
			}
		}

		for _, pId := range r.ScorersHome {
			if row, ok := stats[pId]; ok {
				row.Goals++
			}
		}

		for _, pId := range r.ScorersAway {
			if row, ok := stats[pId]; ok {
				row.Goals++
			}
		}

	}

	return stats
}

func MergeStats(existing, new StatsMap) StatsMap {
	changed := StatsMap{}

	for pId, newRow := range new {
		if row, exists := existing[pId]; exists {
			row.Merge(newRow)
			changed[pId] = row
		} else {
			changed[pId] = newRow
		}

	}

	return changed
}

// Player History
type PHistoryRow struct {
	PlayerId string

	LeagueId   string
	LeagueName string
	TeamId     string
	TeamName   string
	Played     int
	Goals      int
	Score      float64

	HalfSeason   bool
	TransferCost *string
	StartYear    int
}

func NewPHistoryRow(stat *StatRow, leagueName string, gameDate time.Time) *PHistoryRow {
	wasTransferedHalfSeason := false
	// If I am creating this row on January is because it was transfered half season
	if gameDate.Month() == time.January {
		wasTransferedHalfSeason = true
	}
	return &PHistoryRow{
		PlayerId:   stat.PlayerId,
		LeagueId:   stat.LeagueId,
		LeagueName: leagueName,
		TeamId:     stat.TeamId,
		TeamName:   stat.Team.Name,
		Played:     stat.Played,
		Goals:      stat.Goals,
		Score:      stat.Score,
		StartYear:  gameDate.Year(),
		HalfSeason: wasTransferedHalfSeason,
	}
}

// Team History
type THistoryRow struct {
	TeamId string

	LeagueId      string
	LeagueName    string
	FinalPosition int
	Played        int
	Wins          int
	Draws         int
	Losses        int
	Points        int
	GoalScored    int
	GoalConceded  int
	Year          int
}

func NewTHistoryRow(stat *TPHRow, leagueId, leagueName string, gameDate time.Time) *THistoryRow {
	row := stat.Row
	return &THistoryRow{
		TeamId: stat.Team.Id,

		LeagueId:      leagueId,
		LeagueName:    leagueName,
		Played:        row.Played,
		Wins:          row.Wins,
		Draws:         row.Draws,
		Losses:        row.Losses,
		Points:        row.Played,
		GoalScored:    row.Played,
		GoalConceded:  row.Played,
		FinalPosition: stat.Index,

		Year: gameDate.Year(),
	}
}

type FDStatRow struct {
	TeamId    string
	TeamName  string
	HiredDate time.Time

	PlayersSigned int
	PlayersSold   int
	CoachesSigned int
	CoachesSacked int

	MaxSpent    float64
	TotalSpent  float64
	MaxCashed   float64
	TotalCashed float64
}

func NewFDStatRow(date time.Time, teamId, teamName string) *FDStatRow {
	return &FDStatRow{
		TeamId:    teamId,
		TeamName:  teamName,
		HiredDate: date,
	}
}
