package groupUserdb

import (
	entity2 "kinexx_backend/pkg/entity"
	"kinexx_backend/pkg/services/group_service/entity"
)

// import "kinexx_backend/pkg/entity"

type GroupUserService interface {
	AddUserToGroup(groupID, userID, status string) (string, error)
	RemoveUserFromGroup(userID, groupID string) error
	GetGroupsForUser(UserID string) ([]entity.GroupDB, error)
	GetUsersInGroup(groupID string) ([]entity2.ProfileDB, error)
}
