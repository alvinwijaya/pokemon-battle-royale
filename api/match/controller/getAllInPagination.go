package controller

import (
	"net/http"

	"github.com/alvinwijaya/pokemon-battle-royale/api/match/filter"
	"github.com/alvinwijaya/pokemon-battle-royale/model/pagination_model"
	"github.com/alvinwijaya/pokemon-battle-royale/model/response_model"
	"github.com/labstack/echo"
)

func (ctrl *matchController) GetAllInPagination(c echo.Context) (err error) {
	var (
		result response_model.PaginationResponse
		paging pagination_model.Paging
	)

	f := filter.GetAllInPaginationFilter{}
	f = f.FromContext(c)

	paginationData, serviceErr := ctrl.matchService.GetAllInPagination(*paging.FromContext(c), &f)
	if serviceErr != nil {
		return serviceErr.HttpError()
	}

	result.Data = *paginationData

	return c.JSON(http.StatusOK, result)
}
