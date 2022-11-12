package pokemon

import "github.com/alvinwijaya/pokemon-battle-royale/custom_error"

type Service interface {
	GetCount() (uint64, *custom_error.ServiceError)
}
