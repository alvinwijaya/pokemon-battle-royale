package controller

import "github.com/alvinwijaya/pokemon-battle-royale/api/pokemon"

type PokemonController struct {
	pokemonService pokemon.Service
}

func NewController(pokemonService pokemon.Service) *PokemonController {
	return &PokemonController{
		pokemonService: pokemonService,
	}
}
