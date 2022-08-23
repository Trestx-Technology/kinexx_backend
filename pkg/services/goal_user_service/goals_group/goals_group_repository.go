package goal_group

import (
	"kinexx_backend/pkg/services/goal_service/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type GoalGroupRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (entity.GoalGroupDB, error)
	Find(filter, projection bson.M) ([]entity.GoalGroupDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
