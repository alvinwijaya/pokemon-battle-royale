package repository

import (
	"github.com/alvinwijaya/pokemon-battle-royale/model/pagination_model"
	model "github.com/alvinwijaya/pokemon-battle-royale/model/pokemon_average_score"
	"gorm.io/gorm"
)

func (repo *pokemonAverageScoreRepository) GetAllInPagination(db *gorm.DB, paging pagination_model.Paging) (*pagination_model.PaginationData, error) {
	var (
		count      int64
		totalPages int64
		err        error

		offset      = (paging.Page - 1) * paging.Limit
		defaultSort = "average_score DESC"
		tableName   = model.PokemonAverageScore{}.TableName()
		result      = pagination_model.PaginationData{}
		items       = []model.PokemonAverageScore{}
	)

	err = db.Table(tableName).Count(&count).Error
	if err != nil {
		return nil, err
	}

	totalPages = count / int64(paging.Limit)
	if count%int64(paging.Limit) != 0 {
		totalPages += 1
	}

	query := db.Table(tableName).Offset(int(offset)).Limit(int(paging.Limit)).Order(defaultSort)

	err = query.Count(&count).Error
	if err != nil {
		return nil, err
	}

	if count > 0 {
		err = query.Find(&items).Error
		if err != nil {
			return nil, err
		}
	}

	result.Count = uint64(count)
	result.CurrentPage = paging.Page
	result.TotalPages = uint64(totalPages)
	result.Params = paging
	result.Items = &items

	return &result, nil
}
