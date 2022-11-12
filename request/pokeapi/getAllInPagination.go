package pokeapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alvinwijaya/pokemon-battle-royale/config"
	"github.com/alvinwijaya/pokemon-battle-royale/custom_error"
	"github.com/alvinwijaya/pokemon-battle-royale/model/pokeapi"
)

func (req *pokeapiRequester) GetAllInPagination() (*pokeapi.PokeapiPaginationResponseDTO, *custom_error.ServiceError) {
	var (
		response *http.Response
		result   pokeapi.PokeapiPaginationResponseDTO
		err      error
	)

	response, err = http.Get(config.GetConfig().PokeapiUrl + "?limit=" + config.GetConfig().PokeapiPaginationLimit)
	if err != nil {
		return nil, &custom_error.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &custom_error.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, &custom_error.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return &result, nil
}
