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

type TransferMarketInfo struct {
	TransferBudget utils.Money
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

	EmailDecisions map[string]*Decision
	Decisions      []*Decision
	Flags          Flags
	ActionsExps    []time.Time

	// Adding some callbacks to trigger some ui changes
	OnEmployed   func()
	OnUnEmployed func()
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

		EmailDecisions: map[string]*Decision{},
		Decisions:      []*Decision{},
		Flags:          NewFlags(),

		OnEmployed:   func() {},
		OnUnEmployed: func() {},
	}
}

func NewFlags() Flags {
	return Flags{
		HasAContractOffer: false,
	}
}

func NewGameWithId(id, saveName, name, surname string, age int) *Game {
	return &Game{
		Idable:   NewIdable(id),
		SaveName: saveName,
		Name:     name,
		Surname:  surname,
		Age:      age,

		EmailDecisions: map[string]*Decision{},
		Decisions:      []*Decision{},

		OnEmployed:   func() {},
		OnUnEmployed: func() {},
	}
}

func (g *Game) FreeDecisionQueue() {
	g.EmailDecisions = map[string]*Decision{}
	g.Decisions = []*Decision{}
}

func (g *Game) QueueDecision(decision *Decision) {
	if decision.EmailId != "" {
		g.EmailDecisions[decision.EmailId] = decision
	} else {
		g.Decisions = append(g.Decisions, decision)
	}
}

func (g *Game) IsFDTeam(teamId string) bool {
	return g.Team != nil && g.Team.Id == teamId
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

func (g *Game) HasAllNeededDecisions(emailIds []*Idable) bool {
	for _, idable := range emailIds {
		if _, ok := g.EmailDecisions[idable.Id]; !ok {
			return false
		}
	}

	return true
}

func (g *Game) UnsetTeamContract() {
	g.YContract = 0
	g.Wage = utils.NewEurosFromF(0)
	g.Team = nil

	g.OnUnEmployed()
}

func (g *Game) SetTeamContract(yContract int, wage utils.Money, team *TPH) {
	g.YContract = uint8(yContract)
	g.Wage = wage
	g.Team = team

	g.OnEmployed()
}

func (g *Game) IsTransferWindowOpen() (bool, MarketCheck) {
	mc := CalculateTransferWindowDates(g.Date)

	return mc.IsOpen(), mc
}

func (g *Game) IsEmployed() bool {
	return g.Team != nil
}

func (g *Game) GetTeamOrEmpty() *TPH {
	if g.Team != nil {
		return g.Team
	}

	return nil
}

func (g *Game) GetTeamIdOrEmpty() string {
	if g.Team != nil {
		return g.Team.Id
	}

	return ""
}

func (g *Game) IsUnemployedAndNoOfferPending() bool {
	_, hasContract := g.YourContract()

	return !hasContract && !g.Flags.HasAContractOffer
}

type MarketCheck struct {
	IsOpeningDate bool
	IsClosingDate bool
	Summer        bool
	Winter        bool
	OpeningDate   time.Time
	ClosingDate   time.Time
}

func (m MarketCheck) IsOpen() bool {
	return m.IsOpeningDate || (m.Summer || m.Winter)
}

func MakeMarketWindows(date time.Time) []time.Time {
	thisYear := date.Year()

	return []time.Time{
		utils.NewDate(thisYear, conf.SummerMarketWindowStart, 1),
		utils.NewDate(thisYear, conf.SummerMarketWindowEnd, 31),
		utils.NewDate(thisYear, conf.WinterMarketWindowStart, 1),
		utils.NewDate(thisYear, conf.WinterMarketWindowEnd, 31),
	}
}

func CalculateTransferWindowDates(date time.Time) MarketCheck {
	dates := MakeMarketWindows(date)

	if date.Equal(dates[0]) {
		return MarketCheck{
			IsOpeningDate: true, Summer: true,
			OpeningDate: dates[0],
			ClosingDate: dates[1],
		}
	}

	if date.Equal(dates[1]) {
		return MarketCheck{IsClosingDate: true, Summer: true,
			OpeningDate: dates[0],
			ClosingDate: dates[1],
		}
	}

	if date.Equal(dates[2]) {
		return MarketCheck{IsOpeningDate: true, Winter: true,
			OpeningDate: dates[2],
			ClosingDate: dates[3],
		}
	}

	if date.Equal(dates[3]) {
		return MarketCheck{IsClosingDate: true, Winter: true,
			OpeningDate: dates[2],
			ClosingDate: dates[3],
		}
	}

	if date.After(dates[0]) && date.Before(dates[1]) {
		return MarketCheck{Summer: true,
			OpeningDate: dates[2],
			ClosingDate: dates[3],
		}
	}

	if date.After(dates[2]) && date.Before(dates[3]) {
		return MarketCheck{Winter: true}
	}

	return MarketCheck{}
}
