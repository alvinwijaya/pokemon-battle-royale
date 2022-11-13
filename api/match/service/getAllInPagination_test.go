package service

import (
	"errors"
	"net/http"

	"github.com/alvinwijaya/pokemon-battle-royale/api/match/filter"
	"github.com/alvinwijaya/pokemon-battle-royale/custom_error"
	model "github.com/alvinwijaya/pokemon-battle-royale/model/match"
	"github.com/alvinwijaya/pokemon-battle-royale/model/pagination_model"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func (s *MatchSuite) TestGetAllInPagination() {
	var (
		db     = gorm.DB{}
		match  = model.Match{}
		paging = pagination_model.Paging{
			Page:  1,
			Limit: 10,
		}
		paginationResponse = pagination_model.PaginationData{
			Count:       1,
			CurrentPage: 1,
			TotalPages:  1,
			Params: pagination_model.Paging{
				Page:  1,
				Limit: 10,
			},
			Items: &[]model.Match{match},
		}

		expectedResult = pagination_model.PaginationData{
			Count:       1,
			CurrentPage: 1,
			TotalPages:  1,
			Params: pagination_model.Paging{
				Page:  1,
				Limit: 10,
			},
			Items: []model.MatchResponseDTO{match.ToResponseDTO()},
		}
		filter   = filter.GetAllInPaginationFilter{}
		forceErr = errors.New("Force error")
	)

	s.dbManager.EXPECT().GetDB().Return(&db).AnyTimes()

	s.Run("Case Success", func() {
		s.matchRepository.EXPECT().GetAllInPagination(s.dbManager.GetDB(), gomock.Any(), gomock.Any()).Return(&paginationResponse, nil)

		actualResult, gotErr := s.matchService.GetAllInPagination(paging, &filter)

		s.Nil(gotErr)
		s.NotNil(actualResult)
		s.Equal(actualResult, &expectedResult)
	})

	s.Run("Case Failed, error when call matchService.GetAllInPagination", func() {
		s.matchRepository.EXPECT().GetAllInPagination(s.dbManager.GetDB(), gomock.Any(), gomock.Any()).Return(nil, forceErr)

		actualResult, gotErr := s.matchService.GetAllInPagination(paging, &filter)

		s.NotNil(gotErr)
		s.Nil(actualResult)
		s.Equal(gotErr, &custom_error.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: forceErr.Error(),
		})
	})
}
