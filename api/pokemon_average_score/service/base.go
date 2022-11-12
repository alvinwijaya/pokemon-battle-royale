package service

import (
	"github.com/alvinwijaya/pokemon-battle-royale/api/pokemon_average_score"
	"github.com/alvinwijaya/pokemon-battle-royale/db"
)

type pokemonAverageScoreService struct {
	dbManager                     db.DatabaseManager
	pokemonAverageScoreRepository pokemon_average_score.Repository
}

func NewService(
	dbManager db.DatabaseManager,
	pokemonAverageScoreRepository pokemon_average_score.Repository,
) pokemon_average_score.Service {
	return &pokemonAverageScoreService{
		dbManager:                     dbManager,
		pokemonAverageScoreRepository: pokemonAverageScoreRepository,
	}
}
