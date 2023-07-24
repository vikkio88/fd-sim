package models

import (
	"fdsim/enums"
	"fdsim/utils"
	"fmt"

	"github.com/oklog/ulid/v2"
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
	History []*PHistoryRow
	Team    *TPH
	Awards  []Award
}

type Award struct {
	//TODO Maybe to a LPH?
	LeagueId   string
	LeagueName string

	Scorer bool
	Mvp    bool
	Score  float64
	Goals  int
	Played int
	Team   TPH
}

func (a *Award) StatString() string {

	if a.Mvp && a.Scorer {
		score := "-"
		if a.Played > 0 {
			score = fmt.Sprintf("%.2f", a.Score/float64(a.Played))
		}

		return fmt.Sprintf("%d (%d) s: %s", a.Goals, a.Played, score)
	}

	if a.Mvp {
		score := "-"
		if a.Played > 0 {
			score = fmt.Sprintf("%.2f", a.Score/float64(a.Played))
		}
		return fmt.Sprintf("%s (%d)", score, a.Played)
	}

	return fmt.Sprintf("%d (%d)", a.Goals, a.Played)
}

func (a *Award) String() string {
	awards := ""
	if a.Mvp {
		awards += "MVP"
	}

	if a.Scorer {
		if awards != "" {
			awards += ", Top Scorer"
		} else {
			awards += "Top Scorer"
		}
	}
	return awards
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

// Player Placeholder name
type PNPH struct {
	Id      string
	Name    string
	Surname string
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
	YearRetired int
}

func (p *RetiredPlayer) String() string {
	return fmt.Sprintf("%s %s", p.Name, p.Surname)
}
