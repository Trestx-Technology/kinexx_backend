package groupDB

import (
	"kinexx_backend/pkg/services/group_service/entity"
)

type GroupService interface {
	MakeGroup(group *entity.GroupDB) (string, error)
	GetGroup(groupID string) (entity.GroupDB, error)
	GetAllGroup() ([]entity.GroupDB, error)

	GetGroupCreatedByMe(userID string) ([]entity.GroupDB, error)
	DeleteGroup(groupID string) error
	PauseGroup(status, groupID string) (string, error)
	EditGroup(group entity.GroupDB, groupID string) (string, error)
}
