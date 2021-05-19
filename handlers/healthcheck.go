package handlers

import (
	"StarWarsBackEnd/models"

	"net/http"

	"github.com/labstack/echo"
)

// HealthCheck...
func HealthCheck(c echo.Context) error {

	resp := models.HealthCheckResponse{
		Message: "Everything is good!",
	}
	return c.JSON(http.StatusOK, resp)
}
