package handlers

import (
	"StarWarsBackEnd/models"

	"net/http"

	"github.com/labstack/echo"
)

// HealthCheck...
func GetPlanets(c echo.Context) error {

	resp := models.HealthCheckResponse{
		Message: "Auth test!",
	}
	return c.JSON(http.StatusOK, resp)
}
