package generators_test

import (
	"fdsim/enums"
	"fdsim/generators"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnumsBuilder(t *testing.T) {
	g := generators.NewEnumsGen(generatorsTestSeed)
	assert.Equal(t, enums.ES, g.Country())
}
