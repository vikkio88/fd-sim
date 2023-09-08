package vm

import (
	"fdsim/models"

	"golang.org/x/exp/slices"
)

type PendingDecisions struct {
	queue map[models.ActionType][]string
}

func NewPendingDecisions() *PendingDecisions {
	return &PendingDecisions{
		queue: map[models.ActionType][]string{},
	}
}

func (p *PendingDecisions) Add(at models.ActionType, id string) {
	if ids, ok := p.queue[at]; ok {
		ids = append(ids, id)
	}

	p.queue[at] = []string{id}
}

func (p *PendingDecisions) Has(at models.ActionType, id string) bool {
	if ids, ok := p.queue[at]; ok {
		return slices.Contains(ids, id)

	}

	return false
}

func (p *PendingDecisions) Free() {
	p.queue = map[models.ActionType][]string{}
}
