package controller

import (
	"net/http"

	"github.com/alvinwijaya/pokemon-battle-royale/helper"
	"github.com/alvinwijaya/pokemon-battle-royale/model/response_model"
	"github.com/labstack/echo"
)

func (ctrl *matchDetailController) Delete(c echo.Context) (err error) {
	var response response_model.DataResponse

	matchDetailId := helper.ConvertStringToUint64(c.Param("matchDetailId"))

	message, serviceErr := ctrl.matchDetailService.Delete(matchDetailId)
	if serviceErr != nil {
		return serviceErr.HttpError()
	}

	response.Data = map[string]string{
		"message": message,
	}

	return c.JSON(http.StatusOK, response)
}
