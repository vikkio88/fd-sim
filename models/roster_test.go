package models_test

import (
	"fdsim/enums"
	"fdsim/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRosterBuilder(t *testing.T) {
	r := models.NewRoster()
	p := models.NewPlayer("a", "a", 20, enums.DE, models.DF)
	p.SetVals(10, 10, 10)
	r.AddPlayer(&p)
	assert.Equal(t, 1, r.Len())
	assert.Equal(t, 10.0, r.AvgSkill())

	p2 := models.NewPlayer("b", "b", 20, enums.IT, models.DF)
	p2.SetVals(50, 10, 10)
	r.AddPlayer(&p2)
	assert.Equal(t, 2, r.Len())
	assert.Equal(t, 30.0, r.AvgSkill())
}
