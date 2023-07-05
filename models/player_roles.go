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

// Short
const (
	gk_s = "GK"
	df_s = "DF"
	mf_s = "MF"
	st_s = "ST"
)

func getRoleMapping() map[Role]string {
	return map[Role]string{
		GK: gk,
		DF: df,
		MF: mf,
		ST: st,
	}
}

func getRoleMappingShort() map[Role]string {
	return map[Role]string{
		GK: gk_s,
		DF: df_s,
		MF: mf_s,
		ST: st_s,
	}
}

func (r Role) String() string {
	mapping := getRoleMapping()
	if val, ok := mapping[r]; ok {
		return val
	}

	return invalid_role
}

func (r Role) StringShort() string {
	mapping := getRoleMappingShort()
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
