package match_detail

import "time"

type MatchDetail struct {
	ID        uint64 `gorm:"primary_key"`
	MatchID   uint64 `gorm:"not null"`
	PokemonID uint64 `gorm:"not null"`
	Score     uint64 `gorm:"not null"`

	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (MatchDetail) ModelName() string {
	return "Match Detail"
}

func (MatchDetail) TableName() string {
	return "match_details"
}

func (data MatchDetail) ToResponseDTO() MatchDetailResponseDTO {
	return MatchDetailResponseDTO{
		ID:        data.ID,
		MatchID:   data.MatchID,
		PokemonID: data.PokemonID,
		Score:     data.Score,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
