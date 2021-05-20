package services

import (
	"StarWarsBackEnd/models"
	"StarWarsBackEnd/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetPlanets...
func GetPlanets() models.ApiResponse {
	var response models.ApiResponse

	planets, err := repository.GetPlanets()
	if err != nil {
		response.Message = "error"
		return response
	}

	response.Message = "success"
	response.Value = planets

	return response
}

// GetPlanetById...
func GetPlanetById(id string) models.ApiResponse {
	var response models.ApiResponse

	planet, err := repository.GetPlanetById(id)
	if err != nil {
		response.Message = "error"
		response.Value = err.Error()
		return response
	}

	response.Message = "success"
	response.Value = planet
	return response
}

// GetPlanetByName...
func GetPlanetByName(nome string) models.ApiResponse {
	var response models.ApiResponse

	planet, err := repository.GetPlanetByName(nome)
	if err != nil {
		response.Message = "error"
		response.Value = err.Error()
		return response
	}

	response.Message = "success"
	response.Value = planet
	return response
}

// AddPlanet...
func AddPlanet(planet models.Planeta) models.ApiResponse {
	var response models.ApiResponse

	_, err := repository.GetPlanetByName(planet.Nome)
	if err == nil {
		response.Message = "error"
		response.Value = "planeta-já-criado"
		return response
	}

	planet.ID = primitive.NewObjectID()
	planet.Deletado = false
	planet.CreatedAt = time.Now()
	planet.UpdatedAt = time.Now()

	err = repository.AddPlanet(planet)
	if err != nil {
		response.Message = "error"
		response.Value = err.Error()
		return response
	}

	response.Message = "success"
	response.Value = "planeta-criado"
	return response
}

// DeletePlanetById...
func DeletePlanetById(id string) models.ApiResponse {
	var response models.ApiResponse

	planet, err := repository.GetPlanetById(id)
	if err != nil {
		response.Message = "error"
		response.Value = "planeta-não-criado"
		return response
	}

	err = repository.DeletePlanet(planet)
	if err != nil {
		response.Message = "error"
		response.Value = err.Error()
		return response
	}

	response.Message = "success"
	response.Value = "planeta-deletado"
	return response
}
