package generators

import (
	"fdsim/libs"
	"fdsim/models"
)

type ModelsGen struct {
	rng     *libs.Rng
	modules []models.Module
	roles   []models.Role
}

func NewModelsGen(seed int64) *ModelsGen {
	rng := libs.NewRng(seed)
	return NewModelsGenSeeded(rng)
}

func NewModelsGenSeeded(rng *libs.Rng) *ModelsGen {
	return &ModelsGen{
		rng,
		models.AllModules(),
		models.AllPlayerRoles(),
	}
}

func (m *ModelsGen) Role() models.Role {
	idx := m.rng.Index(len(m.roles))
	return m.roles[idx]
}

func (m *ModelsGen) Module() models.Module {
	idx := m.rng.Index(len(m.modules))

	return m.modules[idx]
}
