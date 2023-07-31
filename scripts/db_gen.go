package scripts

import (
	d "fdsim/db"
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/libs"
	"fdsim/models"
	"fdsim/utils"
	"fmt"
	"os"
	"time"
)

type DBGen struct{}

func (DBGen) Run(seed int64, teams int) {
	os.Remove("test.db")
	date := utils.NewDate(2023, time.August, 20)
	game := &models.Game{Idable: models.NewIdable("FakeGameId"), Date: date}
	tg := generators.NewTeamGen(seed)
	ts := tg.TeamsWithCountryUnique(teams, enums.IT)
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
	game.YContract = 2
	game.OnEmployed = func() {}
	game.OnUnEmployed = func() {}

	db := d.NewDb("test.db")
	db.TruncateAll()

	db.LeagueR().InsertOne(l)
	db.GameR().AddStatRow(
		models.NewFDStatRow(game.Date, game.Team.Id, game.Team.Name),
	)

	fmt.Println("Simulating season 1")
	l2 := simulateSeason(l, db, game, false)
	// Moving up to the next season
	game.Date = utils.NewDate(2024, time.August, 20)

	fmt.Println("Simulating season 2")
	l3 := simulateSeason(l2, db, game, false)
	// Moving up to the next season
	game.Date = utils.NewDate(2025, time.August, 20)

	fmt.Println("Simulating season 3")
	// stopping before post season
	simulateSeason(l3, db, game, true)

	// Setting date after post season trigger
	// this is commented if the :61 is true
	// game.Date = utils.NewDate(game.Date.Year(), time.July, 1)
	// this is 31st because :61 is true
	game.Date = utils.NewDate(game.Date.Year(), time.May, 30)
	db.GameR().Update(game)

	fmt.Println("Finished")
}

// Simulating a whole season
func simulateSeason(l *models.League, db d.IDb, game *models.Game, stopBeforePostSeason bool) *models.League {
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
	// if teams number is small
	if len(l.Teams()) < 10 {
		game.Date = game.Date.Add(time.Duration(24*60) * time.Hour)
	}

	if stopBeforePostSeason {
		fmt.Println("\t skipping post season")
		return l
	}

	return db.LeagueR().PostSeason(game)
}
