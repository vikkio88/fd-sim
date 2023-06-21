package services

import (
	"fdsim/db"
	"fdsim/libs"
	"fdsim/models"
	"fmt"
	"time"
)

type Simulator struct {
	rng  *libs.Rng
	db   db.IDb
	game *models.Game
}

func NewSimulatorSeeded(game *models.Game, db db.IDb, rng *libs.Rng) *Simulator {
	return &Simulator{
		rng:  rng,
		db:   db,
		game: game,
	}
}

func NewSimulator(game *models.Game, db db.IDb) *Simulator {
	rng := libs.NewRng(time.Now().Unix())
	return NewSimulatorSeeded(game, db, rng)
}

func (sim *Simulator) Simulate(days int) []Event {
	events := []Event{}

	for i := 1; i <= days; i++ {
		newDate := sim.game.Date.AddDate(0, 0, i)

		if sim.checkForMatches(newDate) {
			league := sim.db.LeagueR().ByIdFull(sim.game.LeagueId)
			// maybe double check that the round date is the same?
			round, ok := league.NextRound()
			if !ok {
				// No More Matches, trigger end of the Season or keep it for the next?
				events = append(events, Event("No More Matches, season over"))
			} else {
				events = append(events, sim.simulateRound(round, league))
			}
		}

		// here there will be logic for events triggering
	}

	return events
}

func (sim *Simulator) simulateRound(round *models.Round, league *models.League) Event {
	round.Simulate(sim.rng)
	league.Update(round)

	//Stats calculation
	oldStats := sim.db.LeagueR().GetStats(league.Id)
	stats := models.StatsFromRoundResult(round, league.Id)
	newStats := models.MergeStats(oldStats, stats)
	//TODO: morale fatigue update

	// db update
	sim.db.LeagueR().PostRoundUpdate(round, league)
	sim.db.LeagueR().UpdateStats(newStats)

	return Event(fmt.Sprintf("Played round %d", round.Index+1))
}

func (sim *Simulator) checkForMatches(date time.Time) bool {
	return sim.db.LeagueR().RoundCountByDate(date) > 0
}
