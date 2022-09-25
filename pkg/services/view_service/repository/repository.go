package viewRepository

import (
	viewEntity "kinexx_backend/pkg/services/view_service/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type Repository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (viewEntity.ViewDB, error)
	Find(filter, projection bson.M) ([]viewEntity.ViewDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
	Aggregate(pipeline bson.A) ([]viewEntity.ViewDB, error)
	Count(filter bson.M) (int64, error)
}
