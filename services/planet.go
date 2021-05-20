package services

import (
	"StarWarsBackEnd/config"
	"StarWarsBackEnd/database"
	"StarWarsBackEnd/models"
	"fmt"

	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetPlanets...
func GetPlanets() ([]models.Planeta, error) {
	var planetas []models.Planeta
	var filter = bson.M{}

	collection := database.Connect().Collection(config.DB_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var planeta models.Planeta
		cursor.Decode(&planeta)
		planetas = append(planetas, planeta)
	}
	if err := cursor.Err(); err != nil {
		return planetas, err
	}

	return planetas, err
}

// GetPlanetById...
func GetPlanetById(id string) (models.Planeta, error) {
	var planeta models.Planeta

	collection := database.Connect().Collection(config.DB_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	oId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id":      oId,
		"deletado": false,
	}

	err := collection.FindOne(ctx, filter).Decode(&planeta)

	emptyId, _ := primitive.ObjectIDFromHex("")
	if planeta.ID == emptyId {
		err = errors.New("planeta-não-encontrado")
	}

	return planeta, err
}

// GetPlanetByName...
func GetPlanetByName(nome string) (models.Planeta, error) {
	var planeta models.Planeta

	collection := database.Connect().Collection(config.DB_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"nome":     nome,
		"deletado": false,
	}

	err := collection.FindOne(ctx, filter).Decode(&planeta)

	emptyId, _ := primitive.ObjectIDFromHex("")
	if planeta.ID == emptyId {
		err = errors.New("planeta-não-encontrado")
	}

	return planeta, err
}

// AddPlanet...
func AddPlanet(planet models.Planeta) error {
	type Planetas struct {
		Count   int                        `json:"count"`
		Next    string                     `json:"next"`
		Previus string                     `json:"previous"`
		Results []models.PlanetaSWResponse `json:"results"`
	}
	var planetasSW Planetas

	collection := database.Connect().Collection(config.DB_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := http.Get(config.URL_SW_PLANETS)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = json.NewDecoder(res.Body).Decode(&planetasSW)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, planeta := range planetasSW.Results {
		if planeta.Name == planet.Nome {
			planet.Aparicoes = len(planeta.FilmURLs)

			_, err := collection.InsertOne(ctx, planet)
			return err
		} else {
			err = errors.New("planeta-não-existe")
		}
	}

	return err
}

// DeletePlanet...
func DeletePlanet(planet models.Planeta) error {
	collection := database.Connect().Collection(config.DB_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": planet.ID}
	update := bson.M{"$set": bson.M{
		"updatedAt": time.Now(),
		"deletado":  true,
	}}

	result, err := collection.UpdateOne(ctx, filter, update)

	fmt.Println(result)

	return err
}
