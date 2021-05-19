package services

import (
	"StarWarsBackEnd/config"
	"StarWarsBackEnd/database"
	"StarWarsBackEnd/models"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPlanetById(id string) (models.Planeta, error) {
	var planeta models.Planeta

	collection := database.Connect().Collection(config.DB_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	oId, _ := primitive.ObjectIDFromHex(id)

	err := collection.FindOne(ctx, models.Planeta{ID: oId}).Decode(&planeta)

	return planeta, err
}

func GetPlanetByName(nome string) (models.Planeta, error) {
	var planeta models.Planeta

	collection := database.Connect().Collection(config.DB_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, models.Planeta{Nome: nome}).Decode(&planeta)

	return planeta, err
}
