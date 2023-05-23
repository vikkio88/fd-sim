package models

import "fdsim/enums"

type Person struct {
	Name    string
	Surname string
	Age     int
	Country enums.Country
}
