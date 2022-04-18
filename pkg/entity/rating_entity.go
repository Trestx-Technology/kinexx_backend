package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	UserID      string             `bson:"user_id" json:"user_id"`
	ItemID      string             `bson:"item_id" json:"item_id"`
	Rating      int                `bson:"rating" json:"rating"`
	Review      string             `bson:"review" json:"review"`
	ItemType    string             `bson:"item_type" json:"item_type"`
	CreatedDate time.Time          `bson:"created_date" json:"created_date"`
	Edited      bool               `bson:"edited" json:"edited"`
	RatingID    string             `json:"rating_id"`
	ItemUser    ProfileDB          `json:"item_user"`
}
