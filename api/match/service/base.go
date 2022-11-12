package service

import (
	"github.com/alvinwijaya/pokemon-battle-royale/api/match"
	"github.com/alvinwijaya/pokemon-battle-royale/api/match_detail"
	"github.com/alvinwijaya/pokemon-battle-royale/api/pokemon_average_score"
	"github.com/alvinwijaya/pokemon-battle-royale/db"
	"github.com/alvinwijaya/pokemon-battle-royale/request/pokeapi"
)

type matchService struct {
	dbManager                     db.DatabaseManager
	matchRepository               match.Repository
	matchDetailRepository         match_detail.Repository
	pokemonAverageScoreRepository pokemon_average_score.Repository
	pokeapiReq                    pokeapi.Interface
}

func NewService(
	dbManager db.DatabaseManager,
	matchRepository match.Repository,
	matchDetailRepository match_detail.Repository,
	pokemonAverageScoreRepository pokemon_average_score.Repository,
	pokeapiReq pokeapi.Interface,
) match.Service {
	return &matchService{
		dbManager:                     dbManager,
		matchRepository:               matchRepository,
		matchDetailRepository:         matchDetailRepository,
		pokemonAverageScoreRepository: pokemonAverageScoreRepository,
		pokeapiReq:                    pokeapiReq,
	}
}
