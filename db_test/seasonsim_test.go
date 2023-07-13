package db_test

import (
	d "fdsim/db"
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/libs"
	"fdsim/models"
	"fdsim/utils"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestPostSeason(t *testing.T) {
	t.Skip("Slow")
	os.Remove("test.db")
	date := utils.NewDate(2023, time.August, 20)
	game := &models.Game{Idable: models.NewIdable("FakeGameId"), Date: date}
	tg := generators.NewTeamGen(0)
	ts := tg.TeamsWithCountryUnique(20, enums.IT)
	// ts := tg.TeamsWithCountryUnique(8, enums.IT)
	l := models.NewLeague(ts, date)
	l.UpdateLocales("Serie A 2023/2024", enums.IT)

	//Setting up game with team
	game.LeagueId = l.Id
	// giving ourselves a random team
	tph := ts[0].PH()
	game.Team = &tph
	game.Wage = utils.NewEurosUF(1000, 0)
	game.Age = 30
	game.Name = "Testio"
	game.Surname = "Mc Test"
	game.SaveName = "TestSave"
	game.Fame = utils.NewPerc(80)
	game.YContract = 1
	game.OnEmployed = func() {}
	game.OnUnEmployed = func() {}

	db := d.NewDb("test.db")
	db.TruncateAll()

	db.LeagueR().InsertOne(l)
	db.GameR().AddStatRow(
		models.NewFDStatRow(game.Date, game.Team.Id, game.Team.Name),
	)

	// Simulating a whole season
	for {
		r, hasMore := l.NextRound()
		if !hasMore {
			break
		}

		r.Simulate(libs.NewRng(0))
		l.Update(r)
		oldStats := db.LeagueR().GetStats(l.Id)
		stats := models.StatsFromRoundResult(r, l.Id)
		newStats := models.MergeStats(oldStats, stats)

		db.LeagueR().PostRoundUpdate(r, l)
		db.LeagueR().UpdateStats(newStats)

		game.Date = r.Date
	}

	// Adding 2 months more or less otherwise it finishes in January
	game.Date = game.Date.Add(time.Duration(24*60) * time.Hour)

	// this will already modify game to the new league
	l2 := db.LeagueR().PostSeason(game, l.Name)

	// End of First Season
	game.Date = utils.NewDate(2024, time.August, 20)

	for {
		r, hasMore := l2.NextRound()
		if !hasMore {
			break
		}

		r.Simulate(libs.NewRng(0))
		l2.Update(r)
		oldStats := db.LeagueR().GetStats(l2.Id)
		stats := models.StatsFromRoundResult(r, l2.Id)
		newStats := models.MergeStats(oldStats, stats)

		db.LeagueR().PostRoundUpdate(r, l2)
		db.LeagueR().UpdateStats(newStats)

		game.Date = r.Date
	}
	// Adding 2 months more or less otherwise it finishes in January
	game.Date = game.Date.Add(time.Duration(24*60) * time.Hour)

	game.Wage.Add(utils.NewEurosUF(25_000, 0))

	// End of Second Season
	db.LeagueR().PostSeason(game, l2.Name)
	fmt.Println("Finished")
}
