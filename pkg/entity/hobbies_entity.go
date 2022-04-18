package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HobbiesDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name,omitempty"`
	Image       string             `bson:"image" json:"image,omitempty"`
	Description string             `bson:"description" json:"description,omitempty"`
}
