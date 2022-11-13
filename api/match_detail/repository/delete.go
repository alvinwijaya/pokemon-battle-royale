package repository

import (
	"time"

	model "github.com/alvinwijaya/pokemon-battle-royale/model/match_detail"
	"gorm.io/gorm"
)

func (repo *matchDetailRepository) Delete(db *gorm.DB, data *model.MatchDetail) (*model.MatchDetail, error) {
	now := time.Now()
	data.DeletedAt = &now

	if err := db.Updates(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
