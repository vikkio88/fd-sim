package generators_test

import (
	"fdsim/generators"
	"fdsim/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModelsBuilder(t *testing.T) {
	g := generators.NewModelsGen(generatorsTestSeed)
	assert.Equal(t, models.MF, g.Role())
	assert.Equal(t, models.M532, g.Module())
}
