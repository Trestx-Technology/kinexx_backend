package groupUserdb

import "kinexx_backend/pkg/entity"

// import "kinexx_backend/pkg/entity"

type GroupUserService interface {
	AddUserToGroup(groupID, userID, status string) (string, error)
	RemoveUserFromGroup(userID, groupID string) error
	GetGroupsForUser(UserID string) ([]entity.GroupDB, error)
	GetUsersInGroup(groupID string) ([]entity.ProfileDB, error)
}
