package generators

import (
	"fdsim/conf"
	"fdsim/data"
	"fdsim/enums"
	"fdsim/libs"
	"fdsim/models"
	"fdsim/utils"
)

type PeopleGen struct {
	rng        *libs.Rng
	eGen       *EnumsGen
	mGen       *ModelsGen
	plAgeRange utils.IntRange
	cAgeRange  utils.IntRange
}

func NewPeopleGenNested(rng *libs.Rng, enGen *EnumsGen) *PeopleGen {
	p := NewPeopleGenSeeded(rng)
	p.eGen = enGen
	return p
}

func NewPeopleGenSeeded(rng *libs.Rng) *PeopleGen {
	pAr := utils.NewIntRangeFromStr(conf.PlayerAgeRange)
	cAr := utils.NewIntRangeFromStr(conf.CoachAgeRange)
	return &PeopleGen{
		rng:        rng,
		plAgeRange: pAr,
		cAgeRange:  cAr,
	}
}

func NewPeopleGen(seed int64) *PeopleGen {
	rng := libs.NewRng(seed)
	return NewPeopleGenSeeded(rng)
}

func (p *PeopleGen) getModelsGen() *ModelsGen {
	if p.mGen != nil {
		return p.mGen
	}

	p.mGen = NewModelsGenSeeded(p.rng)

	return p.mGen
}

func (p *PeopleGen) getEnumsGen() *EnumsGen {
	if p.eGen != nil {
		return p.eGen
	}

	p.eGen = NewEnumsGenSeeded(p.rng)

	return p.eGen
}

func (p *PeopleGen) getName(country enums.Country) string {
	return data.GetNames(country)[0]
}

func (p *PeopleGen) getSurname(country enums.Country) string {
	return data.GetSurnames(country)[0]
}

func (p *PeopleGen) getSkill() int {
	return p.rng.NormPercVal()
}

func (p *PeopleGen) getMorale() int {
	return p.rng.UInt(20, 100)
}

func (p *PeopleGen) getFame(skill int) int {
	return p.rng.UInt(skill+p.rng.PlusMinusVal(10, 50), 100)
}

func (p *PeopleGen) Champion(country enums.Country) *models.Player {
	mGen := p.getModelsGen()
	name := p.getName(country)
	surname := p.getSurname(country)
	age := p.rng.UInt(p.plAgeRange.Min, p.plAgeRange.Max)
	pl := models.NewPlayer(name, surname, age, country, mGen.Role())

	skill := p.rng.UInt(85, 100)
	morale := p.getMorale()
	fame := p.getFame(skill)

	pl.SetVals(skill, morale, fame)

	return &pl
}

func (p *PeopleGen) PlayerWithRole(country enums.Country, role models.Role) *models.Player {
	name := p.getName(country)
	surname := p.getSurname(country)
	age := p.rng.UInt(p.plAgeRange.Min, p.plAgeRange.Max)
	pl := models.NewPlayer(name, surname, age, country, role)

	skill := p.getSkill()
	morale := p.getMorale()
	fame := p.getFame(skill)

	pl.SetVals(skill, morale, fame)

	return &pl
}

func (p *PeopleGen) Player(country enums.Country) *models.Player {
	mGen := p.getModelsGen()

	name := p.getName(country)
	surname := p.getSurname(country)
	age := p.rng.UInt(p.plAgeRange.Min, p.plAgeRange.Max)
	pl := models.NewPlayer(name, surname, age, country, mGen.Role())

	skill := p.getSkill()
	morale := p.getMorale()
	fame := p.getFame(skill)

	pl.SetVals(skill, morale, fame)

	return &pl
}

func (p *PeopleGen) Coach(country enums.Country) *models.Coach {
	mGen := p.getModelsGen()
	name := p.getName(country)
	surname := p.getSurname(country)

	age := p.rng.UInt(p.cAgeRange.Min, p.cAgeRange.Max)
	c := models.NewCoach(name, surname, age, country, mGen.Module())

	skill := p.getSkill()
	morale := p.getMorale()
	fame := p.getFame(skill)

	c.SetVals(skill, morale, fame)

	return &c
}
