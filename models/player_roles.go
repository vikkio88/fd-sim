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
