package repository

import (
	model "github.com/alvinwijaya/pokemon-battle-royale/model/pokemon_average_score"
	"gorm.io/gorm"
)

func (repo *pokemonAverageScoreRepository) GetByPokemonIDs(db *gorm.DB, pokemonIDs []uint64) (*[]model.PokemonAverageScore, error) {
	var result []model.PokemonAverageScore

	if err := db.Where("pokemon_id IN (?)", pokemonIDs).Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
