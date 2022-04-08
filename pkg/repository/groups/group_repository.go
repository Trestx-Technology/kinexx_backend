package groups

import (
	"kinexx_backend/pkg/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type GroupRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.GroupDB, error)
	Find(filter, projection bson.M) ([]entity.GroupDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}

