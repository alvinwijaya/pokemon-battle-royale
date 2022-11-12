package pagination_model

import (
	"github.com/alvinwijaya/pokemon-battle-royale/constant"
	"github.com/alvinwijaya/pokemon-battle-royale/helper"
	"github.com/labstack/echo"
)

type PaginationResponse struct {
	Data interface{} `json:"data"`
}

type PaginationData struct {
	Count       uint64      `json:"count"`
	CurrentPage uint64      `json:"currentPage"`
	TotalPages  uint64      `json:"totalPages"`
	Params      Paging      `json:"params"`
	Items       interface{} `json:"items,omitempty"`
}

type Paging struct {
	Page  uint64 `json:"page"`
	Limit uint64 `json:"limit"`
}

func (p *Paging) FromContext(c echo.Context) *Paging {
	p.Page = helper.ConvertStringToUint64(c.QueryParam("page"))
	p.Limit = helper.ConvertStringToUint64(c.QueryParam("limit"))

	if p.Page == 0 {
		p.Page = constant.DEFAULT_PAGINATION_PAGE
	}
	if p.Limit == 0 {
		p.Limit = constant.DEFAULT_PAGINATION_LIMIT
	}

	return p
}
