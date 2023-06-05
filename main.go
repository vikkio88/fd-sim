package main

import (
	"fdsim/conf"
	"fdsim/db"
	"fdsim/generators"
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		panic("Need action")
	}
	action := args[0]

	db := db.NewDb(conf.DbFiles)
	tc := db.TeamR().Count()
	if tc > 0 {
		fmt.Printf("Had %d some teams\n", tc)
		if action == "list" {
			ts := db.TeamR().All()
			for _, t := range ts {
				fmt.Printf("t: %s - p#: %d skill: %.2f\n", t, t.Roster.Len(), t.Roster.AvgSkill())
			}
			fmt.Printf("teams#: %d\n", db.TeamR().Count())
			fmt.Printf("players#: %d\n", db.PlayerR().Count())
			fmt.Printf("coach#: %d\n", db.CoachR().Count())
			fmt.Printf("id#: %s\n", ts[0].Id)
		}

		if action == "delete" {
			fmt.Printf("deleting team with id#: %s\n", args[1])
			db.TeamR().DeleteOne(args[1])
			fmt.Printf("teamsAfterDel#: %d\n", db.TeamR().Count())
		}

		fmt.Printf("players#: %d\n", db.PlayerR().Count())
		fmt.Printf("free agents coach#: %d\n", len(db.CoachR().FreeAgents()))
		fmt.Printf("free agents players#: %d\n", len(db.PlayerR().FreeAgents()))
		return
	}

	fmt.Println("Had no teams, generating")
	tg := generators.NewTeamGen(time.Now().Unix())
	ts := tg.Teams(20)
	db.TeamR().Insert(ts)
	pg := generators.NewPeopleGen(time.Now().Unix())
	ps := pg.Players(4)
	db.PlayerR().Insert(ps)

	// a := ui.NewApp()
	// a.Run()

	// a.Cleanup()
}
