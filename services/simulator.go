package services

import (
	"fdsim/conf"
	"fdsim/db"
	"fdsim/enums"
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

func (sim *Simulator) persistGameState() {
	sim.db.GameR().Update(sim.game)
}

func (sim *Simulator) Simulate(days int) []*Event {
	events := []*Event{}
	for i := 1; i <= days; i++ {
		newDate := sim.game.Date.AddDate(0, 0, 1)
		//this could go to a notification chan
		fmt.Printf("Simulating day %s\n", newDate.Format(conf.DateFormatGame))
		events = sim.simulateDate(events, newDate)
	}

	// Saving New Game state
	sim.persistGameState()

	return events
}

func (sim *Simulator) simulateDate(events []*Event, newDate time.Time) []*Event {
	// Decisions is a series of queued Choosable taken during the pause stage
	events = sim.applyDecisions(newDate, events)

	// Check if there are matches during this day
	if sim.checkForMatches(newDate) {
		fmt.Printf("Had Matches\n")
		league := sim.db.LeagueR().ByIdFull(sim.game.LeagueId)

		// TODO: maybe double check that the round date is the same?
		round, _ := league.NextRound()

		fmt.Printf("Simulating Round %d\n", round.Index+1)
		events = append(events, sim.simulateRound(round, league))
		events = sim.checkIfLeagueFinished(league, events, newDate)
	}

	// here there will be logic for events triggering
	// ADD EVENT TRIGGERING HERE
	x := sim.game.IsUnemployedAndNoOfferPending()
	y := sim.rng.ChanceI(100)
	if x && y {
		randomTeam := sim.db.TeamR().All()[0]
		var money float64 = 100000
		events = append(events,
			ContractOffer.Event(
				newDate,
				models.EventParams{
					TeamId:   randomTeam.Id,
					TeamName: randomTeam.Name,
					ValueInt: 2,
					ValueF:   money,
					FdName:   sim.game.FootDirector().String(),
					//TODO: This is crap, might need to rethink where to store country
					LeagueCountry: enums.IT,
				}),
		)
	}

	// set new date
	sim.game.Date = newDate
	return events
}

func (sim *Simulator) applyDecisions(newDate time.Time, events []*Event) []*Event {
	// maybe use Decisions as queue and pop
	for _, d := range sim.game.Decisions {
		decisionEvent := ParseDecision(newDate, &d.Choice)
		if decisionEvent != nil {
			events = append(events, decisionEvent)
		}
	}
	sim.game.FreeDecisionQueue()
	return events
}

func (*Simulator) checkIfLeagueFinished(league *models.League, events []*Event, newDate time.Time) []*Event {
	if league.IsFinished() {
		firstRow := league.TableRow(0)
		events = append(
			events,
			LeagueFinished.Event(
				newDate,
				models.EventParams{
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
		models.EventParams{
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

		sim.game.Flags = e.TriggerFlags(sim.game.Flags)
		e.TriggerChanges(sim.game, sim.db)
	}

	//persist notifications on db
	sim.db.GameR().AddEmails(emails)
	sim.db.GameR().AddNews(news)
	sim.persistGameState()

	// return new ones to UI
	return emails, news
}

func (sim *Simulator) checkForMatches(date time.Time) bool {
	return sim.db.LeagueR().RoundCountByDate(date) > 0
}
