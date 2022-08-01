package dtos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Product struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	UUID      string             `bson:"uuid" json:"uuid"`
	Name      string             `bson:"name" json:"name"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type ProductDto struct {
	Name string `json:"name" validate:"required"`
}
