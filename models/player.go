package models

import (
	"fdsim/enums"
	"fdsim/utils"
	"fmt"

	"github.com/oklog/ulid/v2"
	"golang.org/x/exp/slices"
)

const playerInMemoryId = "pmId"

func playerIdGenerator() string {
	return fmt.Sprintf("%s_%s", playerInMemoryId, ulid.Make())
}

type Player struct {
	Idable
	Person
	Role  Role
	Value utils.Money
	skillable
}
type PlayerDetailed struct {
	Player
	History  []*PHistoryRow
	Team     *TPH
	Awards   []Award
	Trophies []Trophy
	Offers   []Offer
}

func (p *PlayerDetailed) WageAcceptanceChance(offer utils.Money, offeringTeamId string) utils.Perc {
	return utils.NewPerc(0)
}

func (p *PlayerDetailed) GetOfferFromTeamId(teamId string) (*Offer, bool) {
	if p.Offers == nil || len(p.Offers) < 1 {
		return nil, false
	}

	idx := slices.IndexFunc(p.Offers, func(o Offer) bool {
		return o.OfferingTeam.Id == teamId
	})

	if idx != -1 {
		return &p.Offers[idx], true
	}
	return nil, false
}

type PlayerHistorical struct {
	Id      string
	Name    string
	Surname string
	Team    *TPH
	Played  int
	Goals   int
	Score   float64
}

func NewPlayerHistoricalFromStatRowPH(row *StatRowPH) *PlayerHistorical {
	return &PlayerHistorical{
		Id:      row.StatRow.Player.Id,
		Name:    row.StatRow.Player.Name,
		Surname: row.StatRow.Player.Surname,
		Team:    row.Team,
		Played:  row.Played,
		Goals:   row.Goals,
		Score:   row.Score,
	}
}

func NewPlayer(name, surname string, age int, country enums.Country, role Role) Player {
	return Player{
		Idable: NewIdable(playerIdGenerator()),
		Person: Person{
			Name:    name,
			Surname: surname,
			Age:     age,
			Country: country,
		},
		Role: role,
		//TODO: add familiarity with a module
	}
}

func (p *Player) String() string {
	return fmt.Sprintf("%s %s", p.Name, p.Surname)
}

func (p *Player) StringShort() string {
	return fmt.Sprintf("%c. %s", p.Name[0], p.Surname)
}

// get placeholder
func (p *Player) PH() PPH {
	return PPH{
		Id:  p.Id,
		sPH: p.skillable.PH(),
		Age: p.Age,
	}
}

func (p *Player) PHName() *PNPH {
	return &PNPH{
		Id:      p.Id,
		Name:    p.Name,
		Surname: p.Surname,
	}
}

// Player Placeholder name
type PNPH struct {
	Id      string
	Name    string
	Surname string
}

// Player Placeholder with Name and Some holding values, to Trigger events on db side
type PNPHVals struct {
	PNPH

	ValueI  int
	ValueI1 int
	ValueI2 int

	ValueF  float64
	ValueF1 float64
	ValueF2 float64

	ValueS  string
	ValueS1 string
	ValueS2 string
}

func (p *PNPH) String() string {
	return fmt.Sprintf("%s %s", p.Name, p.Surname)
}

// Player Placeholder without names
type PPH struct {
	sPH
	Id  string
	Age int
	// TODO: track injuries so we know whether can be choose or not for lineup
}

func NewRolePPHMap() map[Role][]PPH {
	result := map[Role][]PPH{}
	for _, role := range AllPlayerRoles() {
		result[role] = []PPH{}
	}

	return result
}

type RetiredPlayer struct {
	Id      string
	Name    string
	Surname string
	Country enums.Country
	Age     int
	Role    Role

	History     []*PHistoryRow
	Awards      []Award
	Trophies    []Trophy
	YearRetired int
}

func (p *RetiredPlayer) String() string {
	return fmt.Sprintf("%s %s", p.Name, p.Surname)
}
