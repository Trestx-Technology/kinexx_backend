package viewContentRepository

import (
	"go.mongodb.org/mongo-driver/bson"
	viewContentEntity "kinexx_backend/pkg/services/view_content_service/entity"
)

type Repository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (viewContentEntity.ViewContentDB, error)
	Find(filter, projection bson.M) ([]viewContentEntity.ViewContentDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
