package models

type Role uint8

const (
	GK Role = iota
	DF
	MF
	ST
)

const (
	gk = "Goalkeeper"
	df = "Defender"
	mf = "Midfielder"
	st = "Striker"

	invalid_role = "INVALID_ROAD"
)

func getRoleMapping() map[Role]string {
	return map[Role]string{
		GK: gk,
		DF: df,
		MF: mf,
		ST: st,
	}
}

func (r Role) String() string {
	mapping := getRoleMapping()
	if val, ok := mapping[r]; ok {
		return val
	}

	return invalid_role
}

func AllPlayerRoles() []Role {
	return []Role{
		GK,
		DF,
		MF,
		ST,
	}
}

func NewEmptyRoleCounter() map[Role]int {
	return map[Role]int{
		GK: 0,
		DF: 0,
		MF: 0,
		ST: 0,
	}
}

func NewRoleCounter(gk, df, mf, st int) map[Role]int {
	return map[Role]int{
		GK: gk,
		DF: df,
		MF: mf,
		ST: st,
	}
}
