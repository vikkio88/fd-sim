package models

import (
	"fdsim/enums"
	"fdsim/utils"
	"fmt"
	"math"
	"time"

	"github.com/oklog/ulid/v2"
)

const teamInMemoryId = "tmId"

func teamIdGenerator() string {
	return fmt.Sprintf("%s_%s", teamInMemoryId, ulid.Make())
}

type Team struct {
	Idable
	Name    string
	City    string
	Country enums.Country
	Roster  *Roster
	Coach   *Coach

	Balance       utils.Money
	TransferRatio float64

	//TODO: add familiarity with a module
}

type TeamDetailed struct {
	Team
	History []*THistoryRow
}

// Team Placeholder
type TPH struct {
	Id   string
	Name string
}

func NewTeam(name, city string, country enums.Country) *Team {
	return &Team{
		Idable:  NewIdable(teamIdGenerator()),
		Name:    name,
		City:    city,
		Country: country,
		Roster:  NewRoster(),
	}
}

func (t *Team) OfferAcceptanceChance(offer utils.Money, playerId string) utils.Perc {
	p, ok := t.Roster.Player(playerId)
	if !ok {
		//TODO: log error no player in this team
		return utils.NewPerc(0)
	}

	offerVal := offer.Value()
	value := p.Value.Value()
	//TODO: check if player is in transferable
	acceptancePercentage := 40.0

	if t.Balance.Value()-t.Wages().Value() < 0 {
		acceptancePercentage += 30.0
	}

	if math.Abs(offerVal-value) < 0.05*value || offerVal > value {
		acceptancePercentage += 45.0
	}

	if offerVal < value {
		exponentialDecrease := math.Pow(2, (value-offerVal)/value)
		acceptancePercentage -= 3.0 * exponentialDecrease
	} else {
		exponentialIncrease := math.Pow(2, (offerVal-value)/value)
		acceptancePercentage += 3.0 * exponentialIncrease
	}

	if p.Age < 25 {
		acceptancePercentage -= float64(p.Age) / 2
	} else if p.Age > 29 {
		acceptancePercentage += float64(p.Age) / 3
	}

	ycontractDifference := math.Abs(float64(p.YContract - 1))
	if ycontractDifference <= 1 {
		acceptancePercentage += 25.0
	} else if ycontractDifference <= 2 {
		acceptancePercentage += 2.0
	}

	return utils.NewPerc(int(acceptancePercentage))
}

func (t *Team) Wages() utils.Money {
	r := t.Roster.Wages()
	return utils.NewEurosFromF(r.Value() + t.Coach.Wage.Value())
}

func (t *Team) TransferBudget() utils.Money {
	val := t.Balance.Value() * t.TransferRatio
	return utils.NewEurosFromF(val)
}

func (t *Team) String() string {
	return fmt.Sprintf("%s (%s)", t.Name, t.Country)
}

func (t *Team) StringShort() string {
	return fmt.Sprintf("%s", t.Name)
}

func (t *Team) Lineup() *Lineup {
	module := M442
	rngSeed := time.Now().Unix()
	skill := 50
	if t.Coach != nil {
		module = t.Coach.Module
		rngSeed = t.Coach.RngSeed
		skill = t.Coach.Skill.Val()
	}

	return t.Roster.Lineup(module, rngSeed, skill)
}

func (t *Team) PH() TPH {
	return TPH{
		Id:   t.Id,
		Name: t.Name,
	}
}
