package service

import (
	"github.com/alvinwijaya/pokemon-battle-royale/api/match_detail"
	"github.com/alvinwijaya/pokemon-battle-royale/api/pokemon_average_score"
	"github.com/alvinwijaya/pokemon-battle-royale/db"
)

type matchDetailService struct {
	dbManager                     db.DatabaseManager
	matchDetailRepository         match_detail.Repository
	pokemonAverageScoreRepository pokemon_average_score.Repository
}

func NewService(
	dbManager db.DatabaseManager,
	matchDetailRepository match_detail.Repository,
	pokemonAverageScoreRepository pokemon_average_score.Repository,
) match_detail.Service {
	return &matchDetailService{
		dbManager:                     dbManager,
		matchDetailRepository:         matchDetailRepository,
		pokemonAverageScoreRepository: pokemonAverageScoreRepository,
	}
}
