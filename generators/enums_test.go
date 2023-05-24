package generators_test

import (
	"fdsim/generators"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnumsBuilder(t *testing.T) {
	generators.NewEnumsGen(generatorsTestSeed)
	assert.Equal(t, true, true)

	//TODO: add tests
}
