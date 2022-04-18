package chats

import (
	"kinexx_backend/pkg/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type ChatRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.Message, error)
	Find(filter, projection bson.M) ([]entity.Message, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
