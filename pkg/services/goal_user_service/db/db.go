package goalGroupdb

import (
	"kinexx_backend/pkg/entity"

	"errors"
	goal_group "kinexx_backend/pkg/repository/goals/goals_group"
	goalDB "kinexx_backend/pkg/services/goal_service/db"
	groupDB "kinexx_backend/pkg/services/group_service/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = goal_group.NewGoalGroupRepository("goal_groups")
)

type goalUserService struct{}

func NewGoalUserService(repository goal_group.GoalGroupRepository) GoalUserService {
	repo = repository
	return &goalUserService{}
}

func (*goalUserService) AddUserToGoal(goalID, groupID string) (string, error) {

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

func (*goalUserService) RemoveUserFromGoal(groupID, goalID string) error {
	filter := bson.M{"goal_id": goalID, "group_id": groupID}
	return repo.DeleteOne(filter)
}

func (*goalUserService) GetGoalsForUser(groupID string) ([]entity.GoalDB, error) {
	var emptyslice []string
	// why not giving error when not satisfying return value

	filter := bson.M{"group_id": groupID}
	goalIDs, err := repo.Find(filter, bson.M{})
	if err != nil {
		return []entity.GoalDB{}, err
	}
	for _, goal := range goalIDs {
		emptyslice = append(emptyslice, goal.GoalID)
	}

	return goalDB.GetGoalsByIDs(emptyslice)

}

func (*goalUserService) GetUsersInGoal(goalID string) ([]entity.GroupDB, error) {
	var emptyslice []string

	filter := bson.M{"goal_id": goalID}
	goalUserIDs, err := repo.Find(filter, bson.M{"group_id": 1})
	if err != nil {
		return []entity.GroupDB{}, err
	}
	for _, goal := range goalUserIDs {
		emptyslice = append(emptyslice, goal.GroupID)
	}

	return groupDB.GetGroupsByIDs(emptyslice)
}

// []string in GetGoalsForUser vs []entity.GroupDB in GetUsersInGoal
