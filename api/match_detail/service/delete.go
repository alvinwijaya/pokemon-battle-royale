package service

import (
	"net/http"

	"github.com/alvinwijaya/pokemon-battle-royale/custom_error"
	model "github.com/alvinwijaya/pokemon-battle-royale/model/match_detail"
	pas_model "github.com/alvinwijaya/pokemon-battle-royale/model/pokemon_average_score"
)

func (s *matchDetailService) Delete(id uint64) (string, *custom_error.ServiceError) {
	var (
		matchDetail         *model.MatchDetail
		err                 error
		pokemonAverageScore *pas_model.PokemonAverageScore
	)

	matchDetail, err = s.matchDetailRepository.GetOne(s.dbManager.GetDB(),
		map[string]interface{}{
			"id": id,
		},
	)
	if err != nil {
		return "", &custom_error.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	tx := s.dbManager.StartTransaction()

	_, err = s.matchDetailRepository.Delete(tx, matchDetail)
	if err != nil {
		return "", &custom_error.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	pokemonAverageScore, err = s.pokemonAverageScoreRepository.GetOne(tx,
		map[string]interface{}{
			"pokemon_id": matchDetail.PokemonID,
		},
	)
	if err != nil {
		return "", &custom_error.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if pokemonAverageScore != nil {
		_, err = s.pokemonAverageScoreRepository.UpdateByID(tx, pokemonAverageScore.ID, pokemonAverageScore.ToUpdateEventAnnul(matchDetail.Score))
		if err != nil {
			return "", &custom_error.ServiceError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}

	s.dbManager.CommitTransaction(tx)

	return "Successfully deleted.", nil
}
