package pokeapi

import (
	"github.com/alvinwijaya/pokemon-battle-royale/custom_error"
	"github.com/alvinwijaya/pokemon-battle-royale/model/pokeapi"
)

type Interface interface {
	GetAllInPagination() (*pokeapi.PokeapiPaginationResponseDTO, *custom_error.ServiceError)
	GetByID(id uint64) (*pokeapi.PokeapiGetByIDResponseDTO, *custom_error.ServiceError)
}
