package generators

import (
	"fdsim/data"
	"fdsim/enums"
	"fdsim/libs"
	"fdsim/models"
	"fdsim/utils"
	"fmt"
)

func getTeamConfigGeneration() map[models.Role]int {
	return map[models.Role]int{
		models.GK: 2,
		models.DF: 6,
		models.MF: 5,
		models.ST: 4,
	}
}

type TeamGen struct {
	rng  *libs.Rng
	eGen *EnumsGen
	pGen *PeopleGen

	config map[models.Role]int
}

func NewTeamGenSeeded(rng *libs.Rng) *TeamGen {
	eGen := NewEnumsGenSeeded(rng)
	p := NewPeopleGenNested(rng, eGen)
	return &TeamGen{
		rng:    rng,
		eGen:   eGen,
		pGen:   p,
		config: getTeamConfigGeneration(),
	}

}

func NewTeamGen(seed int64) *TeamGen {
	rng := libs.NewRng(seed)
	return NewTeamGenSeeded(rng)
}

func (t *TeamGen) cityName(country enums.Country) string {
	cities := data.GetCities(country)
	idx := t.rng.Index(len(cities))
	return cities[idx]
}

func (t *TeamGen) teamNamePattern(country enums.Country) string {
	tnp := data.GetTeamNamePattern(country)
	idx := t.rng.Index(len(tnp))
	return tnp[idx]
}

func (t *TeamGen) TeamsWithCountry(count int, country enums.Country) []*models.Team {
	teams := make([]*models.Team, count)
	for i := 0; i < count; i++ {
		teams[i] = t.Team(country)
	}

	return teams
}

func (t *TeamGen) TeamsWithCountryUnique(count int, country enums.Country) []*models.Team {
	teams := make([]*models.Team, count)
	teamsNames := map[string]string{}
	for i := 0; i < count; i++ {
		team := t.Team(country)
		for {
			if _, exists := teamsNames[team.Name]; !exists {
				teamsNames[team.Name] = team.Id
				break
			}

			city := t.cityName(country)
			teamNameP := t.teamNamePattern(country)
			team.City = city
			team.Name = fmt.Sprintf(teamNameP, city)
		}

		teams[i] = team
	}

	return teams
}

func (t *TeamGen) Teams(count int) []*models.Team {
	teams := make([]*models.Team, count)
	for i := 0; i < count; i++ {
		teams[i] = t.Team(t.eGen.Country())
	}

	return teams
}

func (t *TeamGen) Team(country enums.Country) *models.Team {
	city := t.cityName(country)
	teamNamePattern := t.teamNamePattern(country)
	team := models.NewTeam(fmt.Sprintf(teamNamePattern, city), city, country)
	players := []*models.Player{}

	for role, count := range t.config {
		for i := 0; i < count; i++ {
			plCountry := t.getCountry(country)
			p := t.pGen.PlayerWithRole(plCountry, role)
			players = append(players, p)
		}
	}

	if t.rng.ChanceI(60) {
		additional := t.rng.UInt(1, 7)
		for i := 0; i < additional; i++ {
			plCountry := t.getCountry(country)
			p := t.pGen.Player(plCountry)
			players = append(players, p)
		}
	}

	// add a random Champions
	if t.rng.NormPercVal() > 70 {
		additional := t.rng.UInt(1, 6)
		for i := 0; i < additional; i++ {
			plCountry := t.getCountry(country)
			p := t.pGen.Champion(plCountry)
			players = append(players, p)
		}
	}

	team.Roster.AddPlayers(players)
	cCountry := t.getCountry(country)
	team.Coach = t.pGen.Coach(cCountry)
	wages := team.Roster.Wages().Value()
	wages += team.Coach.Wage.Value()
	balance := utils.NewEurosFromF(wages)
	balance.Modify(float64(t.rng.PlusMinus(55)) * t.rng.PercR(0, 20))
	team.Balance = balance
	team.TransferRatio = t.rng.PercR(30, 50)
	return team
}

func (t *TeamGen) getCountry(c enums.Country) enums.Country {
	result := c
	if t.rng.ChanceI(50) {
		result = t.eGen.Country()
	}

	return result
}
