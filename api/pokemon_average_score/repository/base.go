package repository

import "github.com/alvinwijaya/pokemon-battle-royale/api/pokemon_average_score"

type pokemonAverageScoreRepository struct {
}

func NewRepository() pokemon_average_score.Repository {
	return &pokemonAverageScoreRepository{}
}
