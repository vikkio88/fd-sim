package services

import (
	"fdsim/models"
	"time"
)

func triggerJobOffer(sim *Simulator, events []*Event, newDate time.Time) []*Event {
	randomTeam := sim.db.TeamR().OneByFame(sim.game.Fame)

	// this should be linked to fame
	var money float64 = float64(sim.rng.UInt(1, 500)) * 1000.0

	years := sim.rng.UInt(1, 3)
	events = append(events,
		ContractOffer.Event(
			newDate,
			models.EventParams{
				TeamId:   randomTeam.Id,
				TeamName: randomTeam.Name,
				ValueInt: years,
				ValueF:   money,
				FdName:   sim.game.FootDirector().String(),
				Country:  sim.game.BaseCountry,
			}),
	)
	return events
}
