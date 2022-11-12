package pokemon_average_score

import (
	"github.com/alvinwijaya/pokemon-battle-royale/custom_error"
	"github.com/alvinwijaya/pokemon-battle-royale/model/pagination_model"
)

type Service interface {
	GetAllInPagination(paging pagination_model.Paging) (*pagination_model.PaginationData, *custom_error.ServiceError)
}
