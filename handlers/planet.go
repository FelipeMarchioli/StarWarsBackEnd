package handlers

import (
	"StarWarsBackEnd/models"
	"StarWarsBackEnd/services"
	"time"

	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetPlanets...
func GetPlanets(c echo.Context) error {
	var response models.ApiResponse

	planets, err := services.GetPlanets()
	if err != nil {
		response.Message = "error"
		return c.JSON(http.StatusNotFound, response)
	}

	response.Message = "success"
	response.Value = planets
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

	planet, err := services.GetPlanetById(id)
	if err != nil {
		response.Message = "error"
		response.Value = err.Error()
		return c.JSON(http.StatusNotFound, response)
	}

	response.Message = "success"
	response.Value = planet
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

	planet, err := services.GetPlanetByName(nome)
	if err != nil {
		response.Message = "error"
		response.Value = err.Error()
		return c.JSON(http.StatusNotFound, response)
	}

	response.Message = "success"
	response.Value = planet
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

	_, err = services.GetPlanetByName(planet.Nome)
	if err == nil {
		response.Message = "error"
		response.Value = "planeta-já-criado"
		return c.JSON(http.StatusNotFound, response)
	}

	planet.ID = primitive.NewObjectID()
	planet.Deletado = false
	planet.CreatedAt = time.Now()
	planet.UpdatedAt = time.Now()

	err = services.AddPlanet(planet)
	if err != nil {
		response.Message = "error"
		response.Value = err.Error()
		return c.JSON(http.StatusNotFound, response)
	}

	response.Message = "success"
	response.Value = "planeta-criado"
	return c.JSON(http.StatusOK, response)
}

// DeletePlanetById...
func DeletePlanetById(c echo.Context) error {
	var response models.ApiResponse
	id := c.QueryParam("id")

	defer c.Request().Body.Close()

	planet, err := services.GetPlanetById(id)
	if err != nil {
		response.Message = "error"
		response.Value = "planeta-não-criado"
		return c.JSON(http.StatusNotFound, response)
	}

	err = services.DeletePlanet(planet)
	if err != nil {
		response.Message = "error"
		response.Value = err.Error()
		return c.JSON(http.StatusNotFound, response)
	}

	response.Message = "success"
	response.Value = "planeta-deletado"
	return c.JSON(http.StatusOK, response)
}
