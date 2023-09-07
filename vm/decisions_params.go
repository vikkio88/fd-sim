package vm

import "fdsim/models"

func MakePlayerContractOfferParams(fdTeamId string, params ChatParams, offer float64, ycontract int) models.EventParams {
	return models.EventParams{
		PlayerId:   params.Player.Id,
		PlayerName: params.Player.String(),
		ValueF:     offer,
		ValueInt:   ycontract,
		FdTeamId:   fdTeamId,
	}
}

func MakeBidForPlayerParams(fdTeamId string, params ChatParams, offer float64) models.EventParams {
	return models.EventParams{
		TeamId:     params.Team.Id,
		TeamName:   params.Team.Name,
		PlayerId:   params.Player.Id,
		PlayerName: params.Player.String(),
		ValueF:     offer,
		FdTeamId:   fdTeamId,
	}
}
