package pokeapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/alvinwijaya/pokemon-battle-royale/config"
	"github.com/alvinwijaya/pokemon-battle-royale/custom_error"
	"github.com/alvinwijaya/pokemon-battle-royale/model/pokeapi"
)

func (req *pokeapiRequester) GetByID(id uint64) (*pokeapi.PokeapiGetByIDResponseDTO, *custom_error.ServiceError) {
	var (
		response *http.Response
		result   pokeapi.PokeapiGetByIDResponseDTO
		err      error
	)

	response, err = http.Get(config.GetConfig().PokeapiUrl + "/" + strconv.Itoa(int(id)))
	if err != nil {
		return nil, &custom_error.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, &custom_error.ServiceError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("invalid pokemonId %v", id),
		}
	}

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
