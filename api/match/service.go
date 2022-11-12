package match

import (
	"github.com/alvinwijaya/pokemon-battle-royale/api/match/filter"
	"github.com/alvinwijaya/pokemon-battle-royale/custom_error"
	model "github.com/alvinwijaya/pokemon-battle-royale/model/match"
	"github.com/alvinwijaya/pokemon-battle-royale/model/pagination_model"
)

type Service interface {
	Store(p *model.StoreMatchDTO) (string, *custom_error.ServiceError)
	GetAllInPagination(paging pagination_model.Paging, f *filter.GetAllInPaginationFilter) (*pagination_model.PaginationData, *custom_error.ServiceError)
}
