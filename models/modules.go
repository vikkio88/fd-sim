package models

type Module uint8

const (
	M442 Module = iota
	M433
	M343
	M352
	M532
)

const (
	m442 = "4-4-2"
	m433 = "4-3-3"
	m343 = "3-4-3"
	m352 = "3-5-2"
	m532 = "5-3-2"

	invalid_module = "INVALID_MODULE"
)

func getModuleMapping() map[Module]string {
	return map[Module]string{
		M442: m442,
		M433: m433,
		M343: m343,
		M352: m352,
		M532: m532,
	}
}

func (m Module) Validate(lineup map[Role][]PPH) bool {
	conf := m.Conf()
	for r, c := range conf {
		if ps, ok := lineup[r]; ok {
			if len(ps) != c {
				return false
			}
		}

	}

	return true
}

func (r Module) Conf() map[Role]int {
	switch r {
	case M442:
		return NewRoleCounter(1, 4, 4, 2)
	case M433:
		return NewRoleCounter(1, 4, 3, 3)
	case M343:
		return NewRoleCounter(1, 3, 4, 3)
	case M352:
		return NewRoleCounter(1, 3, 5, 2)
	case M532:
		return NewRoleCounter(1, 5, 3, 2)
	}

	return NewEmptyRoleCounter()
}
func (r Module) String() string {
	mapping := getModuleMapping()
	if val, ok := mapping[r]; ok {
		return val
	}

	return invalid_module
}

func AllModules() []Module {
	return []Module{
		M442,
		M433,
		M343,
		M352,
		M532,
	}
}
