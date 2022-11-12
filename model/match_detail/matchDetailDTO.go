package match_detail

import "time"

type StoreMatchDetailDTO struct {
	PokemonID uint64 `json:"pokemonId"`
	Score     uint64 `json:"score"`
}

func (data StoreMatchDetailDTO) ToStore(matchID uint64) MatchDetail {
	return MatchDetail{
		MatchID:   matchID,
		PokemonID: data.PokemonID,
		Score:     data.Score,
	}
}

type MatchDetailResponseDTO struct {
	ID        uint64 `json:"id"`
	MatchID   uint64 `json:"matchId"`
	PokemonID uint64 `json:"pokemonId"`
	Score     uint64 `json:"score"`

	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
