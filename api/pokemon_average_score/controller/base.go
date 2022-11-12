package controller

import (
	"github.com/alvinwijaya/pokemon-battle-royale/api/pokemon_average_score"
)

type pokemonAverageScoreController struct {
	pokemonAverageScoreService pokemon_average_score.Service
}

func NewController(pokemonAverageScoreService pokemon_average_score.Service) *pokemonAverageScoreController {
	return &pokemonAverageScoreController{
		pokemonAverageScoreService: pokemonAverageScoreService,
	}
}
