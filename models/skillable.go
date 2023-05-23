package models

import "fdsim/utils"

type skillable struct {
	Skill  utils.Perc
	Morale utils.Perc
	Fame   utils.Perc
}

func (s *skillable) SetVals(skill, morale, fame int) {
	s.Skill.SetVal(skill)
	s.Morale.SetVal(morale)
	s.Fame.SetVal(fame)
}
