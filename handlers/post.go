package handlers

import (
	"net/http"

	"github.com/bfv/httpecho-go/data"
	"github.com/labstack/echo/v4"
)

// GetHandler is the handler for /get requests
func PostHandler(c echo.Context) error {
	data := data.GetCallData(c)
	return c.JSON(http.StatusOK, data)
}
