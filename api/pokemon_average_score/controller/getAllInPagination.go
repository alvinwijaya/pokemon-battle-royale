package controller

import (
	"net/http"

	"github.com/alvinwijaya/pokemon-battle-royale/model/pagination_model"
	"github.com/alvinwijaya/pokemon-battle-royale/model/response_model"
	"github.com/labstack/echo"
)

func (ctrl *pokemonAverageScoreController) GetAllInPagination(c echo.Context) (err error) {
	var (
		result response_model.PaginationResponse
		paging pagination_model.Paging
	)

	paginationData, serviceErr := ctrl.pokemonAverageScoreService.GetAllInPagination(*paging.FromContext(c))
	if serviceErr != nil {
		return serviceErr.HttpError()
	}

	result.Data = *paginationData

	return c.JSON(http.StatusOK, result)
}
