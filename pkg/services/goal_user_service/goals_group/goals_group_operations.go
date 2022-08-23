package goal_group

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"kinexx_backend/pkg/entity"

	trestCommon "github.com/Trestx-technology/trestx-common-go-lib"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type repo struct {
	CollectionName string
}

func NewGoalGroupRepository(collectionName string) GoalGroupRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

func (r *repo) InsertOne(document interface{}) (string, error) {
	group_user, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert follow",
			err,
			logrus.Fields{
				"document":        document,
				"collection name": r.CollectionName,
			})
		return "", err
	}
	group_user_id := group_user.InsertedID.(primitive.ObjectID).Hex()
	return group_user_id, nil
}

//used by update follow ,login and email verifcation
func (r *repo) UpdateOne(filter, update bson.M) (string, error) {
	result, err := trestCommon.UpdateOne(filter, update, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"update follow",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("follow not found(404)")
		trestCommon.ECLog3(
			"update follow",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})
		return "", err
	}
	return "updated successfully", nil
}

//used by get follow ,login and email verification
func (r *repo) FindOne(filter, projection bson.M) (entity.GoalGroupDB, error) {
	var group_user entity.GoalGroupDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&group_user)
	if err != nil {
		trestCommon.ECLog3(
			"Find follow",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return group_user, err
	}
	return group_user, err
}

//not used may use in future for gettin list of follows
func (r *repo) Find(filter, projection bson.M) ([]entity.GoalGroupDB, error) {
	var group_users []entity.GoalGroupDB
	cursor, err := trestCommon.FindSort(filter, projection, bson.M{"_id": -1}, 1000000, 0, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find follows",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, context.Background())
	for cursor.Next(context.TODO()) {
		var group_user entity.GoalGroupDB
		if err = cursor.Decode(&group_user); err != nil {
			trestCommon.ECLog3(
				"Find follows",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return group_users, nil
		}
		group_users = append(group_users, group_user)
	}
	return group_users, nil
}

//not using
func (r *repo) DeleteOne(filter bson.M) error {
	deleteResult, err := trestCommon.DeleteOne(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"delete follow",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("follow not found(404)")
		trestCommon.ECLog3(
			"delete follow",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}
