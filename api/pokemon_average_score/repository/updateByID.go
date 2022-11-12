package repository

import (
	model "github.com/alvinwijaya/pokemon-battle-royale/model/pokemon_average_score"
	"gorm.io/gorm"
)

func (r *pokemonAverageScoreRepository) UpdateByID(db *gorm.DB, id uint64, p map[string]interface{}) (*model.PokemonAverageScore, error) {
	var res model.PokemonAverageScore

	err := db.First(&res, id).Updates(p).Error
	if err != nil {
		return nil, err
	}

	return &res, nil

}
