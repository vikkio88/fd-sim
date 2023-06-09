package models_test

import (
	"fdsim/enums"
	"fdsim/models"
	"fdsim/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildingPlayer(t *testing.T) {
	p := models.NewPlayer("Name", "Surname", 30, enums.EN, models.GK)
	assert.Equal(t, "Name", p.Name)
	assert.Equal(t, "Surname", p.Surname)
	assert.Equal(t, "English", p.Country.Nationality())
	assert.Equal(t, 30, p.Age)
	assert.Equal(t, "Goalkeeper", p.Role.String())

	// idable
	assert.Contains(t, p.Id, "pmId_")

	// Wage and Value
	assert.Equal(t, p.IdealWage.Val, utils.Money{}.Val)
	assert.Equal(t, p.Value.Val, utils.Money{}.Val)
}

func TestPlayerIsSkillable(t *testing.T) {
	p := models.NewPlayer("Mario", "Rossi", 17, enums.DE, models.ST)

	assert.IsType(t, utils.Perc{}, p.Skill)
	assert.IsType(t, utils.Perc{}, p.Morale)

	p.SetVals(50, 99, 10)

	assert.Equal(t, 50, p.Skill.Val())
	assert.Equal(t, 99, p.Morale.Val())
	assert.Equal(t, 10, p.Fame.Val())
}
