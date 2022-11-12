package pokemon_average_score

import "time"

type PokemonAverageScore struct {
	ID           uint64  `gorm:"primary_key"`
	PokemonID    uint64  `gorm:"not null"`
	AverageScore float64 `gorm:"not null"`
	MatchCount   uint64  `gorm:"not null"`

	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (PokemonAverageScore) ModelName() string {
	return "Pokemon Average Score"
}

func (PokemonAverageScore) TableName() string {
	return "pokemon_average_scores"
}

func (data PokemonAverageScore) ToUpdateEventScore(score uint64) map[string]interface{} {
	newAverageScore := ((float64(data.AverageScore) * float64(data.MatchCount)) + float64(score)) / float64(data.MatchCount+1)

	return map[string]interface{}{
		"average_score": newAverageScore,
		"match_count":   data.MatchCount + 1,
	}
}

func (data PokemonAverageScore) ToUpdateEventAnnul(score uint64) map[string]interface{} {
	newAverageScore := ((float64(data.AverageScore) * float64(data.MatchCount)) - float64(score)) / float64(data.MatchCount-1)

	return map[string]interface{}{
		"average_score": newAverageScore,
		"match_count":   data.MatchCount - 1,
	}
}

func (data PokemonAverageScore) ToResponseDTO() PokemonAverageScoreResponseDTO {
	return PokemonAverageScoreResponseDTO(data)
}
