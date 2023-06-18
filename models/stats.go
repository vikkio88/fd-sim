package models

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
	// TODO: Avg this
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
		for _, pId := range m.LineupHome.Ids() {
			stats[pId] = NewStatRowBase(pId, m.Home.Id, leagueId)
		}

		for _, pId := range m.LineupAway.Ids() {
			stats[pId] = NewStatRowBase(pId, m.Away.Id, leagueId)
		}

		r, ok := m.Result()
		if ok {
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
