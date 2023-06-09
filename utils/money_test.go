package utils_test

import (
	u "fdsim/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoneyCreation(t *testing.T) {
	m := u.NewMoney(u.Dollar, 20)
	assert.Equal(t, "20.00 $", m.String())
	m = u.NewMoney(u.Euro, 20)
	assert.Equal(t, "20.00 €", m.String())
	m = u.NewMoney(u.Pound, 20)
	assert.Equal(t, "20.00 £", m.String())
}
func TestMoneyMath(t *testing.T) {
	m := u.NewMoney(u.Dollar, 300)
	n := u.NewMoney(u.Dollar, 150)

	err := m.Add(n)
	assert.Nil(t, err)
	assert.Equal(t, "450.00 $", m.String())

	err = m.Sub(n)
	assert.Nil(t, err)
	assert.Equal(t, "300.00 $", m.String())

	err = m.Sub(u.NewMoney(u.Euro, 140))
	assert.NotNil(t, err)
	assert.Errorf(t, err, "Currency")
	assert.Equal(t, "300.00 $", m.String())
}

func TestMoneyWithFractional(t *testing.T) {
	m := u.NewMoneyUF(u.Dollar, 20, 75)
	assert.Equal(t, "20.75 $", m.String())
	m = u.NewMoneyUF(u.Dollar, 20, 173)
	assert.Equal(t, "21.73 $", m.String())
}

func TestMoneyFromFloat(t *testing.T) {
	m := u.NewMoneyFromF(u.Dollar, 21.54)
	assert.Equal(t, "21.54 $", m.String())
	m = u.NewMoneyFromF(u.Dollar, 21.54)
	assert.Equal(t, "21.54 $", m.String())
}

func TestCustomConstructors(t *testing.T) {
	e := u.NewEuros(21)
	e1 := u.NewEurosUF(21, 54)
	e2 := u.NewEurosFromF(21.54)
	m := u.NewMoney(u.Euro, 21)
	m1 := u.NewMoneyFromF(u.Euro, 21.54)

	assert.Equal(t, m.String(), e.String())
	assert.Equal(t, m1.String(), e1.String())
	assert.Equal(t, m1.String(), e2.String())

	d := u.NewDollars(21)
	d1 := u.NewDollarsUF(21, 54)
	d2 := u.NewDollarsFromF(21.54)
	m = u.NewMoney(u.Dollar, 21)
	m1 = u.NewMoneyFromF(u.Dollar, 21.54)

	assert.Equal(t, m.String(), d.String())
	assert.Equal(t, m1.String(), d1.String())
	assert.Equal(t, m1.String(), d2.String())

	p := u.NewPounds(21)
	p1 := u.NewPoundsUF(21, 54)
	p2 := u.NewPoundsFromF(21.54)
	m = u.NewMoney(u.Pound, 21)
	m1 = u.NewMoneyFromF(u.Pound, 21.54)

	assert.Equal(t, m.String(), p.String())
	assert.Equal(t, m1.String(), p1.String())
	assert.Equal(t, m1.String(), p2.String())
}

func TestMoneyKMB(t *testing.T) {
	eurosK := u.NewEuros(1235)
	eurosM := u.NewEurosFromF(2100235.43)
	eurosB := u.NewEurosUF(1_234_222_222, 10)
	eurosKneg := u.NewEuros(-1235)
	assert.Equal(t, "1.24k €", eurosK.StringKMB())
	assert.Equal(t, "-1.24k €", eurosKneg.StringKMB())
	assert.Equal(t, "2.10m €", eurosM.StringKMB())
	assert.Equal(t, "1.23b €", eurosB.StringKMB())
}

func TestMoneyVal(t *testing.T) {
	m := u.NewEuros(100)
	m2 := u.NewEurosFromF(m.Value())

	assert.Equal(t, m.String(), m2.String())
}

func TestCurrencyFromString(t *testing.T) {
	assert.Equal(t, u.Pound, u.CurrencyFromString("£"))
	assert.Equal(t, u.Dollar, u.CurrencyFromString("$"))
	assert.Equal(t, u.Euro, u.CurrencyFromString("€"))
}
