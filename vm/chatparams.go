package vm

import "fdsim/models"

type ChatParams struct {
	IsPlayerOffer bool
	IsChat        bool
	Player        *models.PlayerDetailed
	Team          *models.TPH
	Coach         *models.Coach
	ValueF        float64
	ValueF1       float64
	ValueI        *int
}

func (c *ChatParams) IsSimpleChat() bool {
	return c.IsChat && !c.IsPlayerOffer
}

func (c *ChatParams) IsCoachChat() bool {
	return c.Coach != nil
}
