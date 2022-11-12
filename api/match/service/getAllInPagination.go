package service

import (
	"net/http"

	"github.com/alvinwijaya/pokemon-battle-royale/api/match/filter"
	"github.com/alvinwijaya/pokemon-battle-royale/custom_error"
	model "github.com/alvinwijaya/pokemon-battle-royale/model/match"
	"github.com/alvinwijaya/pokemon-battle-royale/model/pagination_model"
)

func (s *matchService) GetAllInPagination(paging pagination_model.Paging, f *filter.GetAllInPaginationFilter) (*pagination_model.PaginationData, *custom_error.ServiceError) {
	var (
		result *pagination_model.PaginationData
		err    error

		responseItems = []model.MatchResponseDTO{}
	)

	result, err = s.matchRepository.GetAllInPagination(s.dbManager.GetDB(), paging, f)
	if err != nil {
		return nil, &custom_error.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	items := result.Items.(*[]model.Match)
	if items != nil {
		for _, item := range *items {
			responseItems = append(responseItems, item.ToResponseDTO())
		}
		result.Items = responseItems
	}

	return result, nil
}
