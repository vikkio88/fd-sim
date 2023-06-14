package models

import (
	"fdsim/utils"
	"fmt"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
)

const gameInMemoryId = "gameId"

func gameIdGenerator() string {
	return fmt.Sprintf("%s_%s", gameInMemoryId, ulid.Make())
}

type FootDirector struct {
	Name    string
	Surname string
	Age     int
	Fame    utils.Perc
}

type YourContract struct {
	Team       TPH
	Wage       utils.Money
	YContract  uint8
	Board      utils.Perc
	Supporters utils.Perc
}

type Game struct {
	Idable
	SaveName string
	Name     string
	Surname  string
	Age      int
	Fame     utils.Perc

	Wage       utils.Money
	YContract  uint8
	Board      utils.Perc
	Supporters utils.Perc
	Date       time.Time

	Team     *TPH
	LeagueId string
}

func (g *Game) Update(name, surname string, age int, date time.Time) {
	name = formatName(name)
	surname = formatName(surname)

	saveName := fmt.Sprintf("%s %s", name, surname)
	g.Id = gameIdGenerator()
	g.SaveName = saveName
	g.Name = name
	g.Surname = surname
	g.Age = age
	g.Date = date
	g.Fame = utils.NewPerc(25) //TODO: move to config
}

func formatName(name string) string {
	name = strings.ToLower(name)
	name = strings.Title(name)
	return name
}

func NewGameWithLeagueId(leagueId, saveName, name, surname string, age int, date time.Time) *Game {
	return &Game{
		Idable:   NewIdable(gameIdGenerator()),
		SaveName: saveName,
		Name:     name,
		Surname:  surname,
		Age:      age,
		Date:     date,
		LeagueId: leagueId,
	}
}

func NewGameWithId(id, saveName, name, surname string, age int) *Game {
	return &Game{
		Idable:   NewIdable(id),
		SaveName: saveName,
		Name:     name,
		Surname:  surname,
		Age:      age,
	}
}

func (g *Game) FootDirector() FootDirector {
	return FootDirector{
		Name:    g.Name,
		Surname: g.Surname,
		Age:     g.Age,
		Fame:    g.Fame,
	}
}
func (g *Game) YourContract() (*YourContract, bool) {
	if g.Team == nil {
		return nil, false
	}
	return &YourContract{
		Team:       *g.Team,
		Wage:       g.Wage,
		YContract:  g.YContract,
		Board:      g.Board,
		Supporters: g.Supporters,
	}, true
}
