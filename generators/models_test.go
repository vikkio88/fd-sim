package generators_test

import (
	"fdsim/generators"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModelsBuilder(t *testing.T) {
	generators.NewModelsGen(generatorsTestSeed)
	assert.Equal(t, true, true)

	//TODO: add tests
}
