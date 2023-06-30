package models

import (
	"fdsim/conf"
	"fdsim/enums"
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

func (f FootDirector) String() string {
	return fmt.Sprintf("%s %s", f.Name, f.Surname)
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

	Date time.Time
	// The current season startdate
	StartDate time.Time

	Team     *TPH
	LeagueId string

	// This is connected to the League Country, it is here
	// so I can create News without having to fetch League
	BaseCountry enums.Country

	Decisions []*Decision
	Flags     Flags
}

type Flags struct {
	HasAContractOffer bool
}

func (g *Game) Update(name, surname string, age int, startDate time.Time) {
	name = formatName(name)
	surname = formatName(surname)

	saveName := fmt.Sprintf("%s %s", name, surname)
	g.Id = gameIdGenerator()
	g.SaveName = saveName
	g.Name = name
	g.Surname = surname
	g.Age = age
	g.Date = startDate
	g.StartDate = startDate
	g.Fame = utils.NewPerc(conf.StartingFame)
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

		Decisions: []*Decision{},
		Flags:     Flags{},
	}
}

func NewGameWithId(id, saveName, name, surname string, age int) *Game {
	return &Game{
		Idable:   NewIdable(id),
		SaveName: saveName,
		Name:     name,
		Surname:  surname,
		Age:      age,

		Decisions: []*Decision{},
	}
}

func (g *Game) FreeDecisionQueue() {
	g.Decisions = []*Decision{}
}

func (g *Game) QueueDecision(decision *Decision) {
	g.Decisions = append(g.Decisions, decision)
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

func (g *Game) IsEmployed() bool {
	return g.Team != nil
}
func (g *Game) IsUnemployedAndNoOfferPending() bool {
	_, hasContract := g.YourContract()

	return !hasContract && !g.Flags.HasAContractOffer
}
