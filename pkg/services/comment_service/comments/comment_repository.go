package comments

import (
	"go.mongodb.org/mongo-driver/bson"
	"kinexx_backend/pkg/entity"
)

type CommentRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.CommentDB, error)
	Find(filter, projection bson.M) ([]entity.CommentDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
