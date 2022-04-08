package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CharityDB struct {
	ID            primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Icon          string             `json:"icon,omitempty"`
	Name          string             `json:"name,omitempty"`
	Description   string             `json:"description,omitempty"`
	Banner        string             `json:"banner,omitempty"`
	CreatedDate   time.Time          `json:"created_date,omitempty"`
	CreatorUserID string             `json:"creator_user_id,omitempty"`
}

type CharityGroupDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	GroupID     string             `bson:"group_id" json:"group_id"`
	CharityID   string             `bson:"charity_id" json:"charity_id"`
	Type        string             `bson:"type" json:"type"`
	CreatedDate time.Time          `bson:"created_date" json:"created_date"`
}
