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

func (p *PeopleGen) getContractInfo(idealWage utils.Money) (utils.Money, int) {
	wage := idealWage.Value() * p.rng.PercR(80, 110)
	contractYears := p.rng.UInt(1, 5)
	return utils.NewEurosFromF(wage), contractYears
}

func (p *PeopleGen) GetWage(skill, age int, isCoach bool) utils.Money {
	if age < 19 {
		return utils.NewEuros(int64(p.rng.UInt(5_000, 150_000)))
	}

	var baseVal int

	switch {
	case skill >= 95:
		baseVal = p.rng.UInt(5_000, 15_000)
	case skill >= 90:
		baseVal = p.rng.UInt(3_000, 7_000)
	case skill >= 80:
		baseVal = p.rng.UInt(600, 5_000)
	case skill >= 70:
		baseVal = p.rng.UInt(500, 1_000)
	case skill >= 60:
		baseVal = p.rng.UInt(200, 450)
	case skill >= 50:
		baseVal = p.rng.UInt(90, 200)
	default:
		baseVal = p.rng.UInt(1, 200)
	}

	if isCoach {
		return utils.NewEuros(int64((baseVal) * 250.0))
	}

	return utils.NewEuros(int64((baseVal * 1000)))

}

func (p *PeopleGen) GetValue(skill, age int) utils.Money {
	var baseVal int

	switch {
	case skill >= 95:
		baseVal = p.rng.UInt(90_000, 250_000)
	case skill >= 90:
		baseVal = p.rng.UInt(80_000, 120_000)
	case skill >= 80:
		baseVal = p.rng.UInt(50_000, 90_000)
	case skill >= 60:
		baseVal = p.rng.UInt(200, 800)
	case skill >= 50:
		baseVal = p.rng.UInt(100, 600)
	default:
		baseVal = p.rng.UInt(10, 100)
	}

	return utils.NewEuros(int64(baseVal * 1000))
}

func (p *PeopleGen) getName(country enums.Country) string {
	names := data.GetNames(country)
	idx := p.rng.Index(len(names))
	return names[idx]
}

func (p *PeopleGen) getSurname(country enums.Country) string {
	surnames := data.GetSurnames(country)
	idx := p.rng.Index(len(surnames))
	return surnames[idx]
}

func (p *PeopleGen) getSkill() int {
	return p.rng.NormPercVal()
}

func (p *PeopleGen) getMorale() int {
	return p.rng.UInt(20, 100)
}

func (p *PeopleGen) getFame(skill int) int {
	return p.rng.UInt(20, skill+p.rng.PlusMinusVal(10, 20))
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
	pl.Value = p.GetValue(pl.Skill.Val(), pl.Age)

	pl.IdealWage = p.GetWage(pl.Skill.Val(), pl.Age, false)
	wage, contract := p.getContractInfo(pl.IdealWage)
	pl.Wage = wage
	pl.YContract = contract

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
	pl.IdealWage = p.GetWage(pl.Skill.Val(), pl.Age, false)
	pl.Value = p.GetValue(pl.Skill.Val(), pl.Age)

	wage, contract := p.getContractInfo(pl.IdealWage)
	pl.Wage = wage
	pl.YContract = contract

	return &pl
}

func (p *PeopleGen) YoungPlayersWithRole(count int, role models.Role) []*models.Player {
	players := make([]*models.Player, count)
	for i := 0; i < count; i++ {
		country := p.getEnumsGen().Country()

		name := p.getName(country)
		surname := p.getSurname(country)
		age := p.rng.UInt(p.plAgeRange.Min, 19)
		pl := models.NewPlayer(name, surname, age, country, role)

		skill := p.rng.UInt(40, 65)
		morale := p.getMorale()
		fame := p.getFame(skill)

		pl.SetVals(skill, morale, fame)
		pl.Value = p.GetValue(pl.Skill.Val(), pl.Age)

		pl.IdealWage = p.GetWage(pl.Skill.Val(), pl.Age, false)
		wage, _ := p.getContractInfo(pl.IdealWage)
		pl.Wage = wage
		// young players always 1 year
		pl.YContract = 1
		players[i] = &pl

	}

	return players
}

func (p *PeopleGen) Players(count int) []*models.Player {
	players := make([]*models.Player, count)
	for i := 0; i < count; i++ {
		players[i] = p.Player(p.getEnumsGen().Country())
	}

	return players
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
	pl.Value = p.GetValue(pl.Skill.Val(), pl.Age)

	pl.IdealWage = p.GetWage(pl.Skill.Val(), pl.Age, false)
	wage, contract := p.getContractInfo(pl.IdealWage)
	pl.Wage = wage
	pl.YContract = contract

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
	c.IdealWage = p.GetWage(c.Skill.Val(), c.Age, true)
	wage, contract := p.getContractInfo(c.IdealWage)
	c.Wage = wage
	c.YContract = contract

	c.RngSeed = int64(p.rng.UInt(0, conf.BigInt))

	return &c
}
