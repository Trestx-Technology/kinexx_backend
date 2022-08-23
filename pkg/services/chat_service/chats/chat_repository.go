package chats

import (
	"go.mongodb.org/mongo-driver/bson"
	"kinexx_backend/pkg/services/chat_service/entity"
)

type ChatRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.Message, error)
	Find(filter, projection bson.M) ([]entity.Message, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
