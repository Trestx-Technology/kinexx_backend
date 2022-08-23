package groupUserdb

import (
	"kinexx_backend/pkg/entity"
	"kinexx_backend/pkg/repository/groups/group_user"

	groupDB "kinexx_backend/pkg/services/group_service/db"
	"kinexx_backend/pkg/services/profile_service/db"

	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = group_user.NewGroupUserRepository("group_users")
)

type groupUserService struct{}

func NewGroupUserService(repository group_user.GroupUserRepository) GroupUserService {
	repo = repository
	return &groupUserService{}
}

func (*groupUserService) AddUserToGroup(groupID, userID, status string) (string, error) {

	filter := bson.M{"group_id": groupID, "user_id": userID}
	group_user, err := repo.FindOne(filter, bson.M{})
	if err != nil {
		// group_user not present
		id := primitive.NewObjectID()
		group_user.ID = id
		group_user.CreatedDate = time.Now()
		group_user.GroupID = groupID
		group_user.UserID = userID
		if status == "" {
			status = "ADDED"
		}
		group_user.Status = status
		// how to use set
		return repo.InsertOne(group_user)
		// not inserting groupID, userID when inserting
	}
	// group_user present
	return "", errors.New("document already there")

}

func (*groupUserService) RemoveUserFromGroup(userID, groupID string) error {
	filter := bson.M{"group_id": groupID, "user_id": userID}
	return repo.DeleteOne(filter)
}

func (*groupUserService) GetGroupsForUser(userID string) ([]entity.GroupDB, error) {
	var emptyslice []string
	// why not giving error when not satisfying return value

	filter := bson.M{"user_id": userID}
	groupIDs, err := repo.Find(filter, bson.M{})
	if err != nil {
		return []entity.GroupDB{}, err
	}
	for _, group := range groupIDs {
		emptyslice = append(emptyslice, group.GroupID)
	}

	return groupDB.GetGroupsByIDs(emptyslice)

}

func (*groupUserService) GetUsersInGroup(groupID string) ([]entity.ProfileDB, error) {
	var emptyslice []string

	filter := bson.M{"group_id": groupID}
	groupUserIDs, err := repo.Find(filter, bson.M{})
	if err != nil {
		return []entity.ProfileDB{}, err
	}
	for _, group := range groupUserIDs {
		emptyslice = append(emptyslice, group.UserID)
	}

	return db.GetProfilesForIDs(emptyslice)
}
