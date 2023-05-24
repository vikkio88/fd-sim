package generators

import (
	"fdsim/data"
	"fdsim/enums"
	"fdsim/libs"
	"fdsim/models"
)

type TeamGen struct {
	rng  *libs.Rng
	eGen *EnumsGen
	pGen *PeopleGen
}

func NewTeamGenSeeded(rng *libs.Rng) *TeamGen {
	eGen := NewEnumsGenSeeded(rng)
	p := NewPeopleGenNested(rng, eGen)
	return &TeamGen{
		rng:  rng,
		eGen: eGen,
		pGen: p,
	}

}

func NewTeamGen(seed int64) *TeamGen {
	rng := libs.NewRng(seed)
	return NewTeamGenSeeded(rng)
}

func (t *TeamGen) cityName(country enums.Country) string {
	return data.GetCities(country)[0]
}

func (t *TeamGen) Team(country enums.Country) models.Team {
	city := t.cityName(country)
	team := models.NewTeam(city, city, country)
	players := []*models.Player{
		t.pGen.PlayerWithRole(country, models.GK),
		t.pGen.PlayerWithRole(country, models.GK),
		t.pGen.PlayerWithRole(country, models.DF),
		t.pGen.PlayerWithRole(country, models.DF),
		t.pGen.PlayerWithRole(country, models.MF),
		t.pGen.PlayerWithRole(country, models.MF),
		t.pGen.PlayerWithRole(country, models.ST),
		t.pGen.PlayerWithRole(country, models.ST),
	}

	team.Roster.AddPlayers(players)
	team.Coach = t.pGen.Coach(country)

	return team
}
