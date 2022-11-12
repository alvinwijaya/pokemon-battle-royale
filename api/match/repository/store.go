package repository

import (
	model "github.com/alvinwijaya/pokemon-battle-royale/model/match"
	"gorm.io/gorm"
)

func (repo *matchRepository) Store(db *gorm.DB, data model.Match) (*model.Match, error) {
	if err := db.Create(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
