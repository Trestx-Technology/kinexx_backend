package ratings

import (
	"kinexx_backend/pkg/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type RatingRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.RatingDB, error)
	Find(filter, projection bson.M) ([]entity.RatingDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
