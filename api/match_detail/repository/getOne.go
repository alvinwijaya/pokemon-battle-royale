package repository

import (
	model "github.com/alvinwijaya/pokemon-battle-royale/model/match_detail"
	"gorm.io/gorm"
)

func (repo *matchDetailRepository) GetOne(db *gorm.DB, q map[string]interface{}) (*model.MatchDetail, error) {
	var result model.MatchDetail

	err := db.First(&result, q).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
