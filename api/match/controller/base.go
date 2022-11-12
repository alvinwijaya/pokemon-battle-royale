package controller

import "github.com/alvinwijaya/pokemon-battle-royale/api/match"

type matchController struct {
	matchService match.Service
}

func NewController(matchService match.Service) *matchController {
	return &matchController{
		matchService: matchService,
	}
}
