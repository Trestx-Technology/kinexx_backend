package hobbies

import (
	"kinexx_backend/pkg/services/hobby_service/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type HobbiesRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.HobbiesDB, error)
	Find(filter, projection bson.M) ([]entity.HobbiesDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
