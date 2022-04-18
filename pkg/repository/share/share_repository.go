package shares

import (
	"kinexx_backend/pkg/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type ShareRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.ShareDB, error)
	Find(filter, projection bson.M) ([]entity.ShareDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
