package match

import (
	"time"

	"github.com/alvinwijaya/pokemon-battle-royale/model/match_detail"
)

type StoreMatchDTO struct {
	MatchDetails []match_detail.StoreMatchDetailDTO `json:"matchDetails"`
}

func (data StoreMatchDTO) ToStore() Match {
	return Match{}
}

type MatchResponseDTO struct {
	ID uint64 `json:"id"`

	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`

	MatchDetails []match_detail.MatchDetailResponseDTO `json:"matchDetails"`
}
