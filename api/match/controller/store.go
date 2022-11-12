package controller

import (
	"net/http"

	model "github.com/alvinwijaya/pokemon-battle-royale/model/match"
	"github.com/alvinwijaya/pokemon-battle-royale/model/response_model"
	"github.com/labstack/echo"
)

func (ctrl *matchController) Store(c echo.Context) (err error) {
	var response response_model.DataResponse

	payloadData := new(model.StoreMatchDTO)

	if err := c.Bind(payloadData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	message, serviceErr := ctrl.matchService.Store(payloadData)
	if serviceErr != nil {
		return serviceErr.HttpError()
	}

	response.Data = map[string]string{
		"message": message,
	}

	return c.JSON(http.StatusCreated, response)
}
