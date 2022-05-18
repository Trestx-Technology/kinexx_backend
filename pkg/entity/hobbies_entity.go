package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HobbiesDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Image       string             `bson:"image" json:"image"`
	Description string             `bson:"description" json:"description"`
}
