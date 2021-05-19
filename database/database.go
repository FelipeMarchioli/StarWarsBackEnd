package database

import (
	"StarWarsBackEnd/config"

	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.DB_STRING_CONNECTION))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(config.DB_DATABASE)

	return db
}
