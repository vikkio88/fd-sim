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

func (sim *Simulator) persistGameState() {
	sim.db.GameR().Update(sim.game)
}

func (sim *Simulator) Simulate(days int) ([]*Event, bool) {
	events := []*Event{}
	for i := 1; i <= days; i++ {
		actionsDue := sim.db.GameR().GetActionsDueByDate(sim.game.Date)
		if len(actionsDue) > 0 && !sim.game.HasAllNeededDecisions(actionsDue) {

			return events, false
		}

		newDate := sim.game.Date.AddDate(0, 0, 1)
		//this could go to a notification chan
		fmt.Printf("Simulating day %s\n", newDate.Format(conf.DateFormatGame))
		events = sim.simulateDate(events, newDate)
	}

	// Saving New Game state
	sim.persistGameState()

	return events, true
}

func (sim *Simulator) simulateDate(events []*Event, newDate time.Time) []*Event {
	// Decisions is a series of queued Choosable taken during the pause stage
	events = sim.applyDecisions(newDate, events)

	// Check if there are matches during this day
	if sim.checkForMatches(newDate) {
		fmt.Printf("Had Matches\n")
		league, _ := sim.db.LeagueR().ByIdFull(sim.game.LeagueId)

		// TODO: maybe double check that the round date is the same?
		round, _ := league.NextRound()

		// TODO: remove this log
		fmt.Printf("Simulating Round %d\n", round.Index+1)
		events = append(events, sim.simulateRound(round, league))
		events = sim.checkIfLeagueFinished(league, events, newDate)
	}

	//Events triggering
	events = sim.stateTriggeredEvents(events, newDate)

	// Date triggered Events
	events = sim.dateTriggeredEvents(events, newDate)

	// Db Triggered Events
	events = sim.dbTriggeredEvents(events, newDate)

	// set new date
	sim.game.Date = newDate
	return events
}

func (sim *Simulator) dbTriggeredEvents(events []*Event, newDate time.Time) []*Event {
	newEvents := parseDbEvents(sim.db.GameR().GetEvents(newDate), sim.db)
	events = append(events, newEvents...)
	return events
}

func (sim *Simulator) dateTriggeredEvents(events []*Event, newDate time.Time) []*Event {
	// Season is over
	if newDate.Day() == 1 && newDate.Month() == time.June {
		events = append(events, SeasonOver.Event(newDate, models.EventParams{
			Country:  sim.game.BaseCountry,
			LeagueId: sim.game.LeagueId,
		}))
	}

	transfCheck := models.CalculateTransferWindowDates(sim.game.Date)
	if transfCheck.IsOpen() {
		events = sim.marketEvents(events, transfCheck, newDate)
	}

	return events
}

func (sim *Simulator) stateTriggeredEvents(events []*Event, newDate time.Time) []*Event {
	isHireable := sim.game.IsUnemployedAndNoOfferPending()

	if isHireable && sim.rng.Chance(sim.game.Fame) {
		events = triggerJobOffer(sim, events, newDate)
	}

	return events
}

func (sim *Simulator) marketEvents(events []*Event, mc models.MarketCheck, newDate time.Time) []*Event {
	if mc.OpeningDate {
		events = append(
			events,
			TransferMarketOpen.Event(newDate, models.EventParams{
				Country:  sim.game.BaseCountry,
				BoolFlag: mc.Summer,
				Label1:   mc.Opening,
				Label2:   mc.Closing,
			}),
		)
	}

	if mc.ClosingDate {
		events = append(
			events,
			TransferMarketClose.Event(newDate, models.EventParams{
				Country:  sim.game.BaseCountry,
				BoolFlag: mc.Summer,
				Label1:   mc.Opening,
				Label2:   mc.Closing,
			}),
		)
	}

	return events
}

func (sim *Simulator) applyDecisions(newDate time.Time, events []*Event) []*Event {
	// maybe use Decisions as queue and pop
	for _, d := range sim.game.EmailDecisions {
		decisionEvent := ParseDecision(newDate, &d.Choice)
		if decisionEvent != nil {
			events = append(events, decisionEvent)
		}
	}

	for _, d := range sim.game.Decisions {
		decisionEvent := ParseDecision(newDate, &d.Choice)
		if decisionEvent != nil {
			events = append(events, decisionEvent)
		}
	}
	sim.game.FreeDecisionQueue()
	return events
}

func (sim *Simulator) checkIfLeagueFinished(league *models.League, events []*Event, newDate time.Time) []*Event {
	if league.IsFinished() {
		events = sim.leagueEnd(league, events, newDate)
	}
	return events
}

func (sim *Simulator) leagueEnd(league *models.League, events []*Event, newDate time.Time) []*Event {
	firstRow := league.TableRow(0)
	secondRow := league.TableRow(1)
	thirdRow := league.TableRow(2)

	isEmployed := sim.game.IsEmployed()
	fdTeamId := ""
	fdTeamName := ""
	if isEmployed {
		fdTeamId = sim.game.Team.Id
		fdTeamName = sim.game.Team.Name
	}

	events = append(
		events,
		LeagueFinished.Event(
			newDate,
			models.EventParams{
				Country:    sim.game.BaseCountry,
				LeagueId:   league.Id,
				LeagueName: league.Name,
				TeamId:     firstRow.Team.Id,
				TeamName:   firstRow.Team.Name,
				TeamId1:    secondRow.Team.Id,
				TeamName1:  secondRow.Team.Name,
				TeamId2:    thirdRow.Team.Id,
				TeamName2:  thirdRow.Team.Name,
				IsEmployed: isEmployed,
				FdTeamId:   fdTeamId,
				FdTeamName: fdTeamName,
			},
		),
	)
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
			Country:    sim.game.BaseCountry,
			LeagueId:   league.Id,
			LeagueName: league.Name,
			RoundId:    round.Id,
			Label1:     fmt.Sprintf("%d", round.Index+1),
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
