package entity

import (
	"kinexx_backend/pkg/entity"
	entity2 "kinexx_backend/pkg/services/group_service/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`
	Sender          entity.ProfileDB   `json:"sender"`
	Receiver        entity.ProfileDB   `json:"receiver"`
	Squad           entity2.GroupDB    `json:"squat"`
	Group           bool               `bson:"group" json:"group"`
	SenderID        string             `bson:"sender_id" json:"sender_id"`
	ReceiverID      string             `bson:"receiver_id" json:"receiver_id"`
	Body            string             `bson:"body" json:"body"`
	ContentType     string             `bson:"content_type" json:"content_type"`
	ContentURL      string             `bson:"content_url" json:"content_url"`
	Tags            string             `bson:"tags" json:"tags"`
	SentTime        time.Time          `bson:"sent_time" json:"sent_time"`
	HideForReceiver bool               `bson:"hide_for_receiver" json:"hide_for_receiver"`
	HideForSender   bool               `bson:"hide_for_sender" json:"hide_for_sender"`
}
