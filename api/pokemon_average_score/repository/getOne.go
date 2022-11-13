package repository

import (
	model "github.com/alvinwijaya/pokemon-battle-royale/model/pokemon_average_score"
	"gorm.io/gorm"
)

func (repo *pokemonAverageScoreRepository) GetOne(db *gorm.DB, q map[string]interface{}) (*model.PokemonAverageScore, error) {
	var result model.PokemonAverageScore

	err := db.First(&result, q).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
