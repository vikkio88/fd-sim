package main

import (
	"fdsim/conf"
	"fdsim/db"
	"fdsim/generators"
	"fmt"
	"time"
)

func main() {
	db := db.NewDb(conf.DbFiles)
	tc := db.TeamR().Count()
	if tc > 0 {
		fmt.Printf("Had %d some teams\n", tc)
		ts := db.TeamR().All()
		for _, t := range ts {
			fmt.Printf("t: %s - p#: %d skill: %.2f\n", t, t.Roster.Len(), t.Roster.AvgSkill())
		}

		fmt.Printf("players#: %d\n", db.PlayerR().Count())
		return
	}

	fmt.Println("Had no teams, generating")
	tg := generators.NewTeamGen(time.Now().Unix())
	ts := tg.Teams(40)
	db.TeamR().Insert(ts)

	// a := ui.NewApp()
	// a.Run()

	// a.Cleanup()
}
