package vm

import "fdsim/enums"

func GetAllCountries() []string {
	ce := enums.AllCountries()
	countries := make([]string, len(ce))
	for i, c := range ce {
		countries[i] = c.String()
	}

	return countries
}

func CountryFromIndex(idx int) enums.Country {
	ce := enums.AllCountries()
	return ce[idx]
}
