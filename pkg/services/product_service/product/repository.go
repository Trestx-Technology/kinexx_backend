package products

import (
	"kinexx_backend/pkg/services/product_service/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type ProductRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.ProductDB, error)
	Find(filter, projection bson.M) ([]entity.ProductDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
