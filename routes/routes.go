package routes

import (
	"net/http"

	"github.com/swaggo/swag"

	"github.com/labstack/echo/v4"
)

type Request struct {
	Protocol      string
	Host          string
	RemoteAddress string
	Method        string
	Path          string
	Headers       http.Header
}

// GetRequest example
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param   some_id      path   int     true  "Some ID"
// @Success 200 {string} string	"ok"
// @Router /request/{some_id} [get]
func GetRequest(c echo.Context) error {
	req := c.Request()
	data := Request{
		Protocol:      req.Proto,
		Host:          req.Host,
		RemoteAddress: req.RemoteAddr,
		Method:        req.Method,
		Path:          req.URL.Path,
		Headers:       req.Header,
	}
	return c.JSONPretty(http.StatusOK, data, "    ")
}

func GetSwagger(c echo.Context) error {
	data, err := swag.ReadDoc()
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, data)
}
