package utils

import (
	"fmt"
	"math"
)

type Money struct {
	Val      int64
	Currency Currency
}

const MULTIPLIER_100 int64 = 100
const MULTIPLIERF_100 float64 = 100

func NewEuros(unit int64) Money {
	return NewMoney(Euro, unit)
}
func NewEurosUF(unit int64, fractional int64) Money {
	return NewMoneyUF(Euro, unit, fractional)
}
func NewEurosFromF(amount float64) Money {
	return NewMoneyFromF(Euro, amount)
}

func NewPounds(unit int64) Money {
	return NewMoney(Pound, unit)
}
func NewPoundsUF(unit int64, fractional int64) Money {
	return NewMoneyUF(Pound, unit, fractional)
}
func NewPoundsFromF(amount float64) Money {
	return NewMoneyFromF(Pound, amount)
}

func NewDollars(unit int64) Money {
	return NewMoney(Dollar, unit)
}
func NewDollarsUF(unit int64, fractional int64) Money {
	return NewMoneyUF(Dollar, unit, fractional)
}
func NewDollarsFromF(amount float64) Money {
	return NewMoneyFromF(Dollar, amount)
}

func NewMoney(currency Currency, unit int64) Money {
	return Money{
		Val:      unit * MULTIPLIER_100,
		Currency: currency,
	}
}

func NewMoneyUF(currency Currency, unit, fractional int64) Money {
	return Money{
		Val:      unit*MULTIPLIER_100 + fractional,
		Currency: currency,
	}
}

func NewMoneyFromF(currency Currency, amount float64) Money {
	val, _ := math.Modf(amount * 100)

	return Money{Val: int64(val), Currency: currency}
}

func (m Money) StringKMB() string {
	value := float64(m.Val) / MULTIPLIERF_100

	return fmt.Sprintf("%s %s", formatToKMB(value), m.Currency.String())
}

func (m Money) String() string {
	return fmt.Sprintf("%.2f %s", float64(m.Val)/MULTIPLIERF_100, m.Currency.String())
}

func (m Money) Value() float64 {
	return float64(m.Val) / MULTIPLIERF_100
}

func (m *Money) Add(n Money) error {
	if m.Currency != n.Currency {
		return fmt.Errorf("Currencies")
	}

	m.Val += n.Val
	return nil
}

func (m *Money) Sub(n Money) error {
	if m.Currency != n.Currency {
		return fmt.Errorf("Currencies")
	}

	m.Val -= n.Val
	return nil
}
func (m Money) Cmp(n Money) int {
	if m.Val > n.Val {
		return 1
	}

	if m.Val == n.Val {
		return 0
	}

	return -1
}

func (m Money) SameCurrency(n Money) bool {
	return m.Currency == n.Currency
}

type Currency uint8

const (
	Dollar Currency = iota
	Euro
	Pound
)

func CurrencyFromString(s string) Currency {
	switch s {
	case "$":
		return Dollar
	case "€":
		return Euro
	case "£":
		return Pound
	}

	return Dollar
}

func (c Currency) String() string {
	switch c {
	case Dollar:
		return "$"
	case Euro:
		return "€"
	case Pound:
		return "£"
	}
	return ""
}

type ErrorDifferentCurrency struct {
	left  Currency
	right Currency
}

func NewErrorDifferentCurrency(c Currency, c1 Currency) ErrorDifferentCurrency {
	return ErrorDifferentCurrency{
		c,
		c1,
	}
}

func (e ErrorDifferentCurrency) Error() string {
	return fmt.Sprintf("The currencies are not compatible (%s and %s)", e.left, e.right)
}

type ErrorInsufficientFunds struct {
}

func NewErrorInsufficientFunds() ErrorInsufficientFunds {
	return ErrorInsufficientFunds{}
}

func (e ErrorInsufficientFunds) Error() string {
	return "Insufficient Funds"
}

func formatToKMB(value float64) string {
	res := ""
	if value > 999999999 || value < -999999999 {
		res = fmt.Sprintf("%0.2fb", value/1_000_000_000.)
	} else if value > 999999 || value < -999999 {
		res = fmt.Sprintf("%0.2fm", value/1_000_000.)
	} else if value > 999 || value < -999 {
		res = fmt.Sprintf("%0.2fk", value/1_000)
	}

	return res
}
