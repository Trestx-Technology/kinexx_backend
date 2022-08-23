package notifications

import (
	"kinexx_backend/pkg/services/notification_service/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type NotificationRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.NotiFicationMessage, error)
	Find(filter, projection bson.M) ([]entity.NotiFicationMessage, error)
	// what is filter, projection
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
