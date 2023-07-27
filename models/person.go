package models

import (
	"fdsim/enums"
	"fdsim/utils"
)

type Person struct {
	Name      string
	Surname   string
	Age       int
	IdealWage utils.Money
	Country   enums.Country

	Wage      utils.Money
	YContract int
}
