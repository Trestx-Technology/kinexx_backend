package spots

import (
	"kinexx_backend/pkg/services/spot_service/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type SpotRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.SpotDB, error)
	Find(filter, projection bson.M) ([]entity.SpotDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
