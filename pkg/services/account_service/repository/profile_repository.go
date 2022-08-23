package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"kinexx_backend/pkg/entity"
)

type ProfileRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.ProfileDB, error)
	Find(filter, projection bson.M) ([]entity.ProfileDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
