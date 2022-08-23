package dbInternal

import (
	"errors"
	goal_group "kinexx_backend/pkg/services/goal_user_service/goals_group"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = goal_group.NewGoalGroupRepository("goal_groups")
)

func AddUserToGoalInternal(goalID, groupID string) (string, error) {

	filter := bson.M{"goal_id": goalID, "group_id": groupID}
	goalGroup, err := repo.FindOne(filter, bson.M{})
	if err != nil {
		// goal_group not present
		id := primitive.NewObjectID()
		goalGroup.ID = id
		goalGroup.CreatedDate = time.Now()
		goalGroup.GoalID = goalID
		goalGroup.GroupID = groupID
		// how to use set
		return repo.InsertOne(goalGroup)
		// not inserting goalID, groupID when inserting
	}
	// goal_group present
	return "", errors.New("document already there")

}
