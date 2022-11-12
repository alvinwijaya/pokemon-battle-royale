package service

import (
	"github.com/alvinwijaya/pokemon-battle-royale/api/pokemon"
	"github.com/alvinwijaya/pokemon-battle-royale/request/pokeapi"
)

type pokemonService struct {
	pokeapiReq pokeapi.Interface
}

func NewService(
	pokeapiReq pokeapi.Interface,
) pokemon.Service {
	return &pokemonService{
		pokeapiReq: pokeapiReq,
	}
}
