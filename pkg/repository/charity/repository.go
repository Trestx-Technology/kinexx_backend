package charities

import (
	"kinexx_backend/pkg/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type CharityRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.CharityDB, error)
	Find(filter, projection bson.M) ([]entity.CharityDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
