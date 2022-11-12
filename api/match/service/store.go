package service

import (
	"fmt"
	"net/http"

	"github.com/alvinwijaya/pokemon-battle-royale/config"
	"github.com/alvinwijaya/pokemon-battle-royale/custom_error"
	model "github.com/alvinwijaya/pokemon-battle-royale/model/match"
	"github.com/alvinwijaya/pokemon-battle-royale/model/match_detail"
	detail_model "github.com/alvinwijaya/pokemon-battle-royale/model/match_detail"
	pas_model "github.com/alvinwijaya/pokemon-battle-royale/model/pokemon_average_score"
)

func (s *matchService) Store(p *model.StoreMatchDTO) (string, *custom_error.ServiceError) {
	var (
		match                *model.Match
		matchDetailToStore   []detail_model.MatchDetail
		pokemonAverageScores *[]pas_model.PokemonAverageScore
		serviceErr           *custom_error.ServiceError
		err                  error

		mapPokemonIDToDetail              = map[uint64]match_detail.StoreMatchDetailDTO{}
		pokemonIDs                        = []uint64{}
		mapPokemonIDToPokemonAverageScore = map[uint64]pas_model.PokemonAverageScore{}
	)

	detailCount := config.GetConfig().MatchParticipantCount
	if len(p.MatchDetails) != detailCount {
		return "", &custom_error.ServiceError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("match detail count != %v", detailCount),
		}
	}

	//Validate pokemonIds in details
	for _, detail := range p.MatchDetails {
		_, ok := mapPokemonIDToDetail[detail.PokemonID]
		if ok {
			return "", &custom_error.ServiceError{
				Code:    http.StatusConflict,
				Message: fmt.Sprintf("error duplicate pokemonId %v", detail.PokemonID),
			}
		}

		mapPokemonIDToDetail[detail.PokemonID] = detail
		pokemonIDs = append(pokemonIDs, detail.PokemonID)

		_, serviceErr = s.pokeapiReq.GetByID(detail.PokemonID)
		if serviceErr != nil {
			return "", serviceErr
		}
	}

	// Get existing average scores by pokemonIds
	pokemonAverageScores, err = s.pokemonAverageScoreRepository.GetByPokemonIDs(s.dbManager.GetDB(), pokemonIDs)
	if err != nil {
		return "", &custom_error.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	for _, pokemonAverageScore := range *pokemonAverageScores {
		mapPokemonIDToPokemonAverageScore[pokemonAverageScore.PokemonID] = pokemonAverageScore
	}

	tx := s.dbManager.StartTransaction()

	// Store match data
	match, err = s.matchRepository.Store(tx, p.ToStore())
	if err != nil {
		s.dbManager.RollbackTransaction(tx)
		return "", &custom_error.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	for _, detail := range p.MatchDetails {
		matchDetailToStore = append(matchDetailToStore, detail.ToStore(match.ID))
	}

	// Bulk store match details
	err = s.matchDetailRepository.BulkStore(tx, &matchDetailToStore)
	if err != nil {
		s.dbManager.RollbackTransaction(tx)
		return "", &custom_error.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	for _, pokemonID := range pokemonIDs {
		_, ok := mapPokemonIDToPokemonAverageScore[pokemonID]
		if ok {
			// Case if already exists average score
			_, err = s.pokemonAverageScoreRepository.UpdateByID(tx, mapPokemonIDToPokemonAverageScore[pokemonID].ID, mapPokemonIDToPokemonAverageScore[pokemonID].ToUpdateEventScore(mapPokemonIDToDetail[pokemonID].Score))
			if err != nil {
				s.dbManager.RollbackTransaction(tx)
				return "", &custom_error.ServiceError{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				}
			}
		} else {
			// Case if not exists, insert new average score
			pokemonAverageScore := pas_model.PokemonAverageScore{
				PokemonID:    pokemonID,
				AverageScore: float64(mapPokemonIDToDetail[pokemonID].Score),
				MatchCount:   1,
			}
			_, err = s.pokemonAverageScoreRepository.Store(tx, pokemonAverageScore)
			if err != nil {
				s.dbManager.RollbackTransaction(tx)
				return "", &custom_error.ServiceError{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				}
			}
		}
	}

	s.dbManager.CommitTransaction(tx)

	return "Successfully created.", nil
}
