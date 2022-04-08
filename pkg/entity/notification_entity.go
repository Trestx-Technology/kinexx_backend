package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotiFicationMessage struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	ReceiverID     string             `bson:"receiver_id" json:"receiver_id"`
	SenderID       string             `bson:"sender_id" json:"sender_id"`
	Body           string             `bson:"body" json:"body"`
	Type           string             `bson:"type" json:"type"`
	TypeID         string             `bson:"type_id" json:"type_id"`
	Read           bool               `bson:"read" json:"read"`
	SentTime       time.Time          `bson:"sent_time" json:"sent_time"`
	NotificationID string             `json:"notification_id"`
}
