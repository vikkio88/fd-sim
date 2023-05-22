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

	invalid = "INVALID_ROAD"
)

func getMapping() map[Role]string {
	return map[Role]string{
		GK: gk,
		DF: df,
		MF: mf,
		ST: st,
	}
}

func getReverseMapping() map[string]Role {
	return map[string]Role{
		gk: GK,
		df: DF,
		mf: MF,
		st: ST,
	}
}

func (r Role) String() string {
	mapping := getMapping()
	if val, ok := mapping[r]; ok {
		return val
	}

	return invalid
}
