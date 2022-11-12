package repository

import (
	model "github.com/alvinwijaya/pokemon-battle-royale/model/pokemon_average_score"
	"gorm.io/gorm"
)

func (repo *pokemonAverageScoreRepository) Store(db *gorm.DB, data model.PokemonAverageScore) (*model.PokemonAverageScore, error) {
	if err := db.Create(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
