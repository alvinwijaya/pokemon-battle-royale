package pokemon_average_score

import (
	"github.com/alvinwijaya/pokemon-battle-royale/model/pagination_model"
	model "github.com/alvinwijaya/pokemon-battle-royale/model/pokemon_average_score"
	"gorm.io/gorm"
)

type Repository interface {
	GetByPokemonIDs(db *gorm.DB, pokemonIDs []uint64) (*[]model.PokemonAverageScore, error)
	Store(db *gorm.DB, data model.PokemonAverageScore) (*model.PokemonAverageScore, error)
	UpdateByID(db *gorm.DB, id uint64, p map[string]interface{}) (*model.PokemonAverageScore, error)
	GetAllInPagination(db *gorm.DB, paging pagination_model.Paging) (*pagination_model.PaginationData, error)
}
