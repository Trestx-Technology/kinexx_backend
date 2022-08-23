package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Connection represents the model for an Connection
type FollowDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	UserID      string             `bson:"user_id" json:"user_id"`
	FollowedID  string             `bson:"followed_id" json:"followed_id"`
	Status      string             `bson:"status" json:"status"`
	CreatedTime time.Time          `bson:"created_time" json:"created_time"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time"`
	User        []ProfileDB        `json:"user"`
	FollowID    string             `json:"follow_id"`
}
