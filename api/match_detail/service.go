package match_detail

import "github.com/alvinwijaya/pokemon-battle-royale/custom_error"

type Service interface {
	Delete(id uint64) (string, *custom_error.ServiceError)
}
