package repository

import (
	"github.com/alvinwijaya/pokemon-battle-royale/api/match_detail"
)

type matchDetailRepository struct {
}

func NewRepository() match_detail.Repository {
	return &matchDetailRepository{}
}
