package service

import (
	"net/http"

	"github.com/alvinwijaya/pokemon-battle-royale/custom_error"
	"github.com/alvinwijaya/pokemon-battle-royale/model/pagination_model"
	model "github.com/alvinwijaya/pokemon-battle-royale/model/pokemon_average_score"
)

func (s *pokemonAverageScoreService) GetAllInPagination(paging pagination_model.Paging) (*pagination_model.PaginationData, *custom_error.ServiceError) {
	var (
		result *pagination_model.PaginationData
		err    error

		responseItems = []model.PokemonAverageScoreResponseDTO{}
	)

	result, err = s.pokemonAverageScoreRepository.GetAllInPagination(s.dbManager.GetDB(), paging)
	if err != nil {
		return nil, &custom_error.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	items := result.Items.(*[]model.PokemonAverageScore)
	if items != nil {
		for _, item := range *items {
			responseItems = append(responseItems, item.ToResponseDTO())
		}
		result.Items = responseItems
	}

	return result, nil
}
