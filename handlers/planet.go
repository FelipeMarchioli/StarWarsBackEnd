package handlers

import (
	"StarWarsBackEnd/models"
	"StarWarsBackEnd/services"

	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// GetPlanets...
func GetPlanets(c echo.Context) error {
	response := services.GetPlanets()

	return c.JSON(http.StatusOK, response)
}

// GetPlanetById...
func GetPlanetById(c echo.Context) error {
	var response models.ApiResponse
	id := c.QueryParam("id")

	if id == "" {
		response.Message = "error"
		response.Value = "parâmetro-inválido"
		return c.JSON(http.StatusInternalServerError, response)
	}

	response = services.GetPlanetById(id)

	return c.JSON(http.StatusOK, response)
}

// GetPlanetByName...
func GetPlanetByName(c echo.Context) error {
	var response models.ApiResponse
	nome := c.QueryParam("nome")

	if nome == "" {
		response.Message = "error"
		response.Value = "parâmetro-inválido"
		return c.JSON(http.StatusInternalServerError, response)
	}

	response = services.GetPlanetByName(nome)

	return c.JSON(http.StatusOK, response)
}

// AddPlanet...
func AddPlanet(c echo.Context) error {
	var response models.ApiResponse
	var planet models.Planeta

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&planet)
	if err != nil {
		log.Fatal(err)
		response.Message = "error"
		response.Value = "invalid-entity-fields"
		return c.JSON(http.StatusNotFound, response)
	}

	response = services.AddPlanet(planet)

	return c.JSON(http.StatusOK, response)
}

// DeletePlanetById...
func DeletePlanetById(c echo.Context) error {
	var response models.ApiResponse
	id := c.QueryParam("id")

	if id == "" {
		response.Message = "error"
		response.Value = "parâmetro-inválido"
		return c.JSON(http.StatusInternalServerError, response)
	}

	response = services.DeletePlanetById(id)

	return c.JSON(http.StatusOK, response)
}
