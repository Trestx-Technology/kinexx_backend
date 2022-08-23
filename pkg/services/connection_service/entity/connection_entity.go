package entity

import (
	"kinexx_backend/pkg/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Connection represents the model for an Connection
type ConnectionDB struct {
	ID                    primitive.ObjectID `bson:"_id" json:"id"`
	UserID                string             `bson:"user_id" json:"user_id"`
	ReceiverID            string             `bson:"receiver_id" json:"receiver_id"`
	Type                  string             `bson:"type" json:"type"`
	Message               string             `bson:"message" json:"message"`
	ContentType           string             `bson:"content_type" json:"content_type"`
	ContentURL            string             `bson:"content_url" json:"content_url"`
	MessageByReceiver     string             `bson:"message_by_receiver" json:"message_by_receiver"`
	ContentTypeByReceiver string             `bson:"content_type_by_receiver" json:"content_type_by_receiver"`
	ContentURLByReceiver  string             `bson:"content_url_by_receiver" json:"content_url_by_receiver"`
	Status                string             `bson:"status" json:"status"`
	CreatedTime           time.Time          `bson:"created_time" json:"created_time"`
	UpdatedTime           time.Time          `bson:"updated_time" json:"updated_time"`
	User                  []entity.ProfileDB `json:"user"`
	Receiver              entity.ProfileDB   `json:"receiver"`
	ConnectionID          string             `json:"connection_id"`
}
