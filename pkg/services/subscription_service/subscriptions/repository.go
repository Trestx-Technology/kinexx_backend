package subcrriptions

import (
	"kinexx_backend/pkg/services/subscription_service/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type SubscriptionsRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.SubscriptionType, error)
	Find(filter, projection bson.M) ([]entity.SubscriptionType, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
