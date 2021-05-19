package handlers

import (
	"StarWarsBackEnd/models"
	"StarWarsBackEnd/services"

	"net/http"

	"github.com/labstack/echo"
)

// HealthCheck...
func GetPlanetById(c echo.Context) error {
	var response models.ApiResponse
	id := c.QueryParam("id")

	if id == "" {
		response.Message = "error"
		response.Value = "parâmetro-inválido"
		return c.JSON(http.StatusInternalServerError, response)
	}

	planet, err := services.GetPlanetById(id)
	if err != nil {
		response.Message = "error"
		return c.JSON(http.StatusNotFound, response)
	}

	response.Message = "success"
	response.Value = planet
	return c.JSON(http.StatusOK, response)
}

func GetPlanetByName(c echo.Context) error {
	var response models.ApiResponse
	nome := c.QueryParam("nome")

	if nome == "" {
		response.Message = "error"
		response.Value = "parâmetro-inválido"
		return c.JSON(http.StatusInternalServerError, response)
	}

	planet, err := services.GetPlanetByName(nome)
	if err != nil {
		response.Message = "error"
		return c.JSON(http.StatusNotFound, response)
	}

	response.Message = "success"
	response.Value = planet
	return c.JSON(http.StatusOK, response)
}
