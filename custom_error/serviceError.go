package custom_error

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type ServiceError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (serviceErr *ServiceError) HttpError() *echo.HTTPError {
	httpCode := serviceErr.Code

	if httpCode == 0 {
		httpCode = http.StatusInternalServerError
	}

	fmt.Println(serviceErr)

	return echo.NewHTTPError(httpCode, *serviceErr)
}
