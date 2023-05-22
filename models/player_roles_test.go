package models_test

import (
	"fdsim/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoleStringify(t *testing.T) {
	assert.Equal(t, "Goalkeeper", models.GK.String())
	assert.Equal(t, "Defender", models.DF.String())
	assert.Equal(t, "Midfielder", models.MF.String())
	assert.Equal(t, "Striker", models.ST.String())
}
