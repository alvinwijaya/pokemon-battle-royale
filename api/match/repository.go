package match

import (
	"github.com/alvinwijaya/pokemon-battle-royale/api/match/filter"
	model "github.com/alvinwijaya/pokemon-battle-royale/model/match"
	"github.com/alvinwijaya/pokemon-battle-royale/model/pagination_model"
	"gorm.io/gorm"
)

type Repository interface {
	Store(db *gorm.DB, data model.Match) (*model.Match, error)
	GetAllInPagination(db *gorm.DB, paging pagination_model.Paging, f *filter.GetAllInPaginationFilter) (*pagination_model.PaginationData, error)
}
