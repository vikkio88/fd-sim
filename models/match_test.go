package models_test

import (
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchBuilder(t *testing.T) {
	g := generators.NewTeamGen(0)
	home := g.Team(enums.IT)
	away := g.Team(enums.EN)

	m := models.NewMatch(home, away)
	assert.Equal(t, home.Id, m.Home.Id)
	assert.Equal(t, away.Id, m.Away.Id)
}
