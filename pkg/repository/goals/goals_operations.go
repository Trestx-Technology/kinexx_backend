package goals

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

func NewGoalRepository(collectionName string) GoalRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

func (r *repo) InsertOne(document interface{}) (string, error) {
	goal, err := trestCommon.InsertOne(document, r.CollectionName)
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
	goalid := goal.InsertedID.(primitive.ObjectID).Hex()
	return goalid, nil
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
func (r *repo) FindOne(filter, projection bson.M) (entity.GoalDB, error) {
	var goal entity.GoalDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&goal)
	if err != nil {
		trestCommon.ECLog3(
			"Find follow",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return goal, err
	}
	return goal, err
}

//not used may use in future for gettin list of follows
func (r *repo) Find(filter, projection bson.M) ([]entity.GoalDB, error) {
	var goals []entity.GoalDB
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
		err := cursor.Close(ctx)
		if err != nil {
		}
	}(cursor, context.Background())
	for cursor.Next(context.TODO()) {
		var goal entity.GoalDB
		if err = cursor.Decode(&goal); err != nil {
			trestCommon.ECLog3(
				"Find follows",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return goals, nil
		}
		goals = append(goals, goal)
	}
	return goals, nil
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
