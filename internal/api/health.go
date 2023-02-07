package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthCheck method    Checks if the API is running
func (a App) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "healthy")
}
