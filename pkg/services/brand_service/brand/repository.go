package brands

import (
	"go.mongodb.org/mongo-driver/bson"
	"kinexx_backend/pkg/services/brand_service/entity"
)

type BrandRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.BrandDB, error)
	Find(filter, projection bson.M) ([]entity.BrandDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
	Count(filter bson.M) (int64, error)
}
