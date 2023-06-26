package services

import (
	"fdsim/conf"
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

func (sim *Simulator) Simulate(days int) []*Event {
	// TODO: Here apply actions maybe?

	events := []*Event{}
	for i := 1; i <= days; i++ {
		newDate := sim.game.Date.AddDate(0, 0, 1)
		fmt.Printf("Simulating day %s\n", newDate.Format(conf.DateFormatGame))

		if sim.checkForMatches(newDate) {
			fmt.Printf("Had Matches\n")
			league := sim.db.LeagueR().ByIdFull(sim.game.LeagueId)
			// maybe double check that the round date is the same?
			round, _ := league.NextRound()
			fmt.Printf("Simulating Round %d\n", round.Index+1)
			events = append(events, sim.simulateRound(round, league))
			events = sim.checkIfLeagueFinished(league, events, newDate)
		}

		// here there will be logic for events triggering

		// set new date
		sim.game.Date = newDate
	}

	// Saving New Game state
	sim.db.GameR().Update(sim.game)

	return events
}

func (*Simulator) checkIfLeagueFinished(league *models.League, events []*Event, newDate time.Time) []*Event {
	if league.IsFinished() {
		firstRow := league.TableRow(0)
		events = append(
			events,
			LeagueFinished.Event(
				newDate,
				EventParams{
					LeagueId:      league.Id,
					LeagueName:    league.Name,
					LeagueCountry: league.Country,
					TeamId1:       firstRow.Team.Id,
					Label1:        firstRow.Team.Name,
				},
			),
		)
	}
	return events
}

func (sim *Simulator) simulateRound(round *models.Round, league *models.League) *Event {
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

	return RoundPlayed.Event(
		round.Date,
		EventParams{
			LeagueId:      league.Id,
			LeagueName:    league.Name,
			LeagueCountry: league.Country,
			RoundId:       round.Id,
			Label1:        fmt.Sprintf("%d", round.Index+1),
		},
	)
}

func (sim *Simulator) SettleEventsTriggers(events []*Event) ([]*models.Email, []*models.News) {
	emails := []*models.Email{}
	news := []*models.News{}
	for _, e := range events {
		if e.TriggerNews != nil {
			news = append(news, e.TriggerNews)
		}

		if e.TriggerEmail != nil {
			emails = append(emails, e.TriggerEmail)
		}
	}

	//persist notifications on db
	sim.db.GameR().AddEmails(emails)
	sim.db.GameR().AddNews(news)

	// return new ones to UI
	return emails, news
}

func (sim *Simulator) checkForMatches(date time.Time) bool {
	return sim.db.LeagueR().RoundCountByDate(date) > 0
}
