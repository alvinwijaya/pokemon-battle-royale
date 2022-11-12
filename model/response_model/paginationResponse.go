package response_model

import "github.com/alvinwijaya/pokemon-battle-royale/model/pagination_model"

type PaginationResponse struct {
	Data pagination_model.PaginationData `json:"data"`
}
