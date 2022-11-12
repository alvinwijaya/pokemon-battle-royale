package filter

import "github.com/labstack/echo"

type GetAllInPaginationFilter struct {
	StartDate string
	EndDate   string
}

func (f GetAllInPaginationFilter) FromContext(c echo.Context) GetAllInPaginationFilter {
	f.StartDate = c.QueryParam("startDate")
	f.EndDate = c.QueryParam("endDate")

	return f
}
