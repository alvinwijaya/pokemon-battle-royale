package pokemon_average_score

import "time"

type PokemonAverageScoreResponseDTO struct {
	ID           uint64  `json:"id"`
	PokemonID    uint64  `json:"pokemonId"`
	AverageScore float64 `json:"averageScore"`
	MatchCount   uint64  `json:"matchCount"`

	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
