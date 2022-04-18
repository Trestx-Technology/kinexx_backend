package groupUserInternal

import (
	"kinexx_backend/pkg/repository/groups/group_user"

	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = group_user.NewGroupUserRepository("group_users")
)

func AddUserToGroupInternal(groupID, userID string) (string, error) {

	filter := bson.M{"group_id": groupID, "user_id": userID}
	group_user, err := repo.FindOne(filter, bson.M{})
	if err != nil {
		// group_user not present
		id := primitive.NewObjectID()
		group_user.ID = id
		group_user.CreatedDate = time.Now()
		group_user.GroupID = groupID
		group_user.UserID = userID
		group_user.Status = "ADDED"
		// how to use set
		return repo.InsertOne(group_user)
		// not inserting groupID, userID when inserting
	}
	// group_user present
	return "", errors.New("document already there")

}
