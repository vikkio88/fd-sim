package libs_test

import (
	"fdsim/libs"
	"fdsim/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testSeed        = 0
	testIntSequence = "00101011000110001000"
)

func TestRngBuilder(t *testing.T) {
	rng := libs.NewRng(testSeed)
	result := ""
	for i := 0; i < 20; i++ {
		result += fmt.Sprintf("%d", rng.UInt(0, 1))
	}

	assert.Equal(t, testIntSequence, result)

	rng = libs.NewRng(2)
	result = ""
	for i := 0; i < 20; i++ {
		result += fmt.Sprintf("%d", rng.UInt(0, 1))
	}
	assert.NotEqual(t, testIntSequence, result)
}

func TestRngChance(t *testing.T) {
	rng := libs.NewRng(testSeed)
	assert.True(t, rng.ChanceI(80))

	rng = libs.NewRng(testSeed)
	assert.True(t, rng.Chance(utils.NewPerc(80)))

	rng = libs.NewRng(1000)
	assert.False(t, rng.ChanceI(0))
	assert.True(t, rng.ChanceI(100))
}

func TestRngPickOne(t *testing.T) {
	//TODO: move this to generics once fyne supports that version of go
	rng := libs.NewRng(testSeed)
	things := []string{"some", "stuff", "maybe", "this"}

	idx := rng.UInt(0, len(things)-1)

	assert.Equal(t, "maybe", things[idx])

}
