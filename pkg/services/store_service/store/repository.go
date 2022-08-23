package stores

import (
	"kinexx_backend/pkg/services/store_service/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type StoreRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.ShopDB, error)
	Find(filter, projection bson.M) ([]entity.ShopDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
