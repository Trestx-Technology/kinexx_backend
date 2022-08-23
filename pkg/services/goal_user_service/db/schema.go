package goalGroupdb

import (
	entity2 "kinexx_backend/pkg/services/goal_service/entity"
	"kinexx_backend/pkg/services/group_service/entity"
)

// import "kinexx_backend/pkg/entity"

type GoalUserService interface {
	AddUserToGoal(groupID, userID string) (string, error)
	RemoveUserFromGoal(userID, groupID string) error
	GetGoalsForUser(UserID string) ([]entity2.GoalDB, error)
	GetUsersInGoal(groupID string) ([]entity.GroupDB, error)
}
