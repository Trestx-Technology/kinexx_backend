package goalGroupdb

import "kinexx_backend/pkg/entity"

// import "kinexx_backend/pkg/entity"

type GoalUserService interface {
	AddUserToGoal(groupID, userID string) (string, error)
	RemoveUserFromGoal(userID, groupID string) error
	GetGoalsForUser(UserID string) ([]entity.GoalDB, error)
	GetUsersInGoal(groupID string) ([]entity.GroupDB, error)
}
