package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Planeta...
type Planeta struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Nome      string             `json:"nome" bson:"nome"`
	Clima     string             `json:"clima" bson:"clima"`
	Terreno   string             `json:"terreno" bson:"terreno"`
	Aparicoes int                `json:"aparicoes" bson:"aparicoes"`
	Deletado  bool               `json:"deletado" bson:"deletado"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
