package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CharityDB struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	Icon          string             `json:"icon"`
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	Banner        string             `json:"banner"`
	CreatedDate   time.Time          `json:"created_date"`
	CreatorUserID string             `json:"creator_user_id"`
}

type CharityGroupDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	GroupID     string             `bson:"group_id" json:"group_id"`
	CharityID   string             `bson:"charity_id" json:"charity_id"`
	Type        string             `bson:"type" json:"type"`
	CreatedDate time.Time          `bson:"created_date" json:"created_date"`
}
