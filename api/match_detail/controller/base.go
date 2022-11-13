package controller

import "github.com/alvinwijaya/pokemon-battle-royale/api/match_detail"

type matchDetailController struct {
	matchDetailService match_detail.Service
}

func NewController(matchDetailService match_detail.Service) *matchDetailController {
	return &matchDetailController{
		matchDetailService: matchDetailService,
	}
}
