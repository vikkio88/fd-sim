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
	tc := db.TeamsCount()
	if db.TeamsCount() > 0 {
		fmt.Printf("Had %d some teams\n", tc)
		ts := db.AllTeams()
		t := db.TeamById(ts[len(ts)-1].Id)
		fmt.Printf("t: %s - ps: %d", t.Id, t.Roster.Len())
		return
	}

	fmt.Println("Had no teams, generating")
	tg := generators.NewTeamGen(time.Now().Unix())
	ts := tg.Teams(40)
	db.InsertManyTeams(ts)

	// a := ui.NewApp()
	// a.Run()

	// a.Cleanup()
}
