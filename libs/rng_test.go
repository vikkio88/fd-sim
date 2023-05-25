package libs_test

import (
	"fdsim/libs"
	"fdsim/utils"
	"fmt"
	"sort"
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

func TestPlusMinus(t *testing.T) {
	rng := libs.NewRng(testSeed)

	assert.Greater(t, 2, rng.PlusMinus(50))
	assert.Less(t, -2, rng.PlusMinus(50))
	val := 10
	assert.Greater(t, val+1, rng.PlusMinusVal(val, 50))
	assert.Less(t, -val-1, rng.PlusMinusVal(val, 50))
}

func TestNormalDistro(t *testing.T) {
	rng := libs.NewRng(testSeed)
	results := map[int]int{}
	keys := []int{}
	for i := 0; i < 200; i++ {
		res := rng.NormF64(0.65, 0.13)
		key := int(res * 100.0)
		keys = append(keys, key)
		results[key] += 1
	}

	sort.Ints(keys)
	// for _, k := range keys {
	// 	fmt.Printf("%d : %d\n", k, results[k])
	// }

	assert.True(t, results[keys[0]] < results[keys[len(keys)/2]])
	assert.True(t, results[keys[len(keys)-1]] < results[keys[len(keys)/2]])

	assert.Equal(t, 100, keys[len(keys)-1])
}
