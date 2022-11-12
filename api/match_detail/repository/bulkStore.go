package repository

import (
	model "github.com/alvinwijaya/pokemon-battle-royale/model/match_detail"
	"gorm.io/gorm"
)

func (repo *matchDetailRepository) BulkStore(db *gorm.DB, data *[]model.MatchDetail) error {
	for _, d := range *data {
		if err := db.Create(&d).Error; err != nil {
			return err
		}
	}

	return nil
}
