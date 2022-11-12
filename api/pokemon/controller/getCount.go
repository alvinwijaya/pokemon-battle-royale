package controller

import (
	"net/http"

	"github.com/alvinwijaya/pokemon-battle-royale/model/response_model"
	"github.com/labstack/echo"
)

func (ctrl *PokemonController) GetCount(c echo.Context) (err error) {
	var response response_model.DataResponse

	count, serviceErr := ctrl.pokemonService.GetCount()
	if serviceErr != nil {
		return serviceErr.HttpError()
	}

	response.Data = map[string]uint64{
		"count": count,
	}

	return c.JSON(http.StatusOK, response)
}
