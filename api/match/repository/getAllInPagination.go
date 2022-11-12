package repository

import (
	"github.com/alvinwijaya/pokemon-battle-royale/api/match/filter"
	model "github.com/alvinwijaya/pokemon-battle-royale/model/match"
	"github.com/alvinwijaya/pokemon-battle-royale/model/pagination_model"
	"gorm.io/gorm"
)

func (repo *matchRepository) GetAllInPagination(db *gorm.DB, paging pagination_model.Paging, f *filter.GetAllInPaginationFilter) (*pagination_model.PaginationData, error) {
	var (
		count      int64
		totalPages int64
		err        error

		offset      = (paging.Page - 1) * paging.Limit
		defaultSort = "id DESC"
		tableName   = model.Match{}.TableName()
		result      = pagination_model.PaginationData{}
		items       = []model.Match{}
	)

	err = db.Table(tableName).Count(&count).Error
	if err != nil {
		return nil, err
	}

	totalPages = count / int64(paging.Limit)
	if count%int64(paging.Limit) != 0 {
		totalPages += 1
	}

	queryCount := db.Table(tableName).Offset(int(offset)).Limit(int(paging.Limit)).Order(defaultSort)
	query := db.Table(tableName).Preload("MatchDetails").Offset(int(offset)).Limit(int(paging.Limit)).Order(defaultSort)

	if f != nil {
		if f.StartDate != "" {
			queryCount = queryCount.Where("created_at >= ?", f.StartDate+" 00:00:00")
			query = query.Where("created_at >= ?", f.StartDate+" 00:00:00")
		}
		if f.EndDate != "" {
			queryCount = queryCount.Where("created_at <= ?", f.EndDate+" 23:59:59")
			query = query.Where("created_at <= ?", f.EndDate+" 23:59:59")
		}
	}

	err = queryCount.Count(&count).Error
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
