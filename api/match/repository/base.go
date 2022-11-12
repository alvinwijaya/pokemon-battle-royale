package repository

import "github.com/alvinwijaya/pokemon-battle-royale/api/match"

type matchRepository struct {
}

func NewRepository() match.Repository {
	return &matchRepository{}
}
