package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Share represents the model for an Share
type ShareDB struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	UserID        string             `bson:"user_id" json:"user_id"`
	SharedItemID  string             `bson:"shared_item_id" json:"shared_item_id"`
	ShredItemType string             `bson:"shred_item_type" json:"shred_item_type"`
	ReceiverID    string             `bson:"receiver_id" json:"receiver_id"`
	Type          string             `bson:"type" json:"type"`
	Status        string             `bson:"status" json:"status"`
	Message       string             `bson:"message" json:"message"`
	CreatedTime   time.Time          `bson:"created_time" json:"created_time"`
	UpdatedTime   time.Time          `bson:"updated_time" json:"updated_time"`
	Receivers     []interface{}      `json:"receivers"`
	Receiver      interface{}        `json:"receiver"`
	Content       interface{}        `json:"content"`
	ShareID       string             `json:"Share_id"`
}
