package models

import "fdsim/utils"

type skillable struct {
	Skill  utils.Perc
	Morale utils.Perc
	Fame   utils.Perc
}
type sPH struct {
	Skill  int
	Morale int
	Fame   int
}

func (s *skillable) SetVals(skill, morale, fame int) {
	s.Skill.SetVal(skill)
	s.Morale.SetVal(morale)
	s.Fame.SetVal(fame)
}
func (s *skillable) PH() sPH {
	return sPH{
		s.Skill.Val(),
		s.Morale.Val(),
		s.Fame.Val(),
	}
}
