package match_detail

import (
	model "github.com/alvinwijaya/pokemon-battle-royale/model/match_detail"
	"gorm.io/gorm"
)

type Repository interface {
	BulkStore(db *gorm.DB, data *[]model.MatchDetail) error
	Delete(db *gorm.DB, data *model.MatchDetail) (*model.MatchDetail, error)
	GetOne(db *gorm.DB, q map[string]interface{}) (*model.MatchDetail, error)
}
