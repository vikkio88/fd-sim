package enums

type Country uint8

const (
	IT Country = iota
	EN
	FR
	DE
	ES
)

const (
	it = "Italy"
	en = "England"
	fr = "France"
	de = "Germany"
	es = "Spain"

	invalid = "INVALID_COUNTRY"
)

func getMapping() map[Country]string {
	return map[Country]string{
		IT: it,
		EN: en,
		FR: fr,
		DE: de,
		ES: es,
	}
}

func getNationalityMapping() map[Country]string {
	return map[Country]string{
		IT: "Italian",
		EN: "English",
		FR: "French",
		DE: "German",
		ES: "Spanish",
	}
}

func (r Country) String() string {
	mapping := getMapping()
	if val, ok := mapping[r]; ok {
		return val
	}

	return invalid
}

func (r Country) Nationality() string {
	mapping := getNationalityMapping()
	if val, ok := mapping[r]; ok {
		return val
	}

	return invalid
}
