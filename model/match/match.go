package match

import (
	"time"

	"github.com/alvinwijaya/pokemon-battle-royale/model/match_detail"
)

type Match struct {
	ID uint64 `gorm:"primary_key"`

	CreatedAt *time.Time
	UpdatedAt *time.Time

	MatchDetails []match_detail.MatchDetail `gorm:"association_autocreate:false;association_autoupdate:false"`
}

func (Match) ModelName() string {
	return "Match"
}

func (Match) TableName() string {
	return "matches"
}

func (data Match) ToResponseDTO() MatchResponseDTO {
	result := MatchResponseDTO{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	matchDetailsResponseDTO := []match_detail.MatchDetailResponseDTO{}
	for _, matchDetail := range data.MatchDetails {
		matchDetailsResponseDTO = append(matchDetailsResponseDTO, matchDetail.ToResponseDTO())
	}

	result.MatchDetails = matchDetailsResponseDTO

	return result
}
