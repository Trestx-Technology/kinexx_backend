package shares

import (
	"context"
	"errors"
	"kinexx_backend/pkg/entity"

	"github.com/aekam27/trestCommon"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type repo struct {
	CollectionName string
}

//NewFirestoreRepository creates a new repo
func NewShareRepository(collectionName string) ShareRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

//used by signup
func (r *repo) InsertOne(document interface{}) (string, error) {
	user, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert share",
			err,
			logrus.Fields{
				"document":        document,
				"collection name": r.CollectionName,
			})
		return "", err
	}
	userid := user.InsertedID.(primitive.ObjectID).Hex()
	return userid, nil
}

//used by update share ,login and email verifcation
func (r *repo) UpdateOne(filter, update bson.M) (string, error) {
	result, err := trestCommon.UpdateOne(filter, update, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"update share",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("share not found(404)")
		trestCommon.ECLog3(
			"update share",
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

//used by get share ,login and email verification
func (r *repo) FindOne(filter, projection bson.M) (entity.ShareDB, error) {
	var share entity.ShareDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&share)
	if err != nil {
		trestCommon.ECLog3(
			"Find share",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return share, err
	}
	return share, err
}

//not used may use in future for gettin list of shares
func (r *repo) Find(filter, projection bson.M) ([]entity.ShareDB, error) {
	var shares []entity.ShareDB
	cursor, err := trestCommon.FindSort(filter, projection, bson.M{"_id": -1}, 1000000, 0, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find shares",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var share entity.ShareDB
		if err = cursor.Decode(&share); err != nil {
			trestCommon.ECLog3(
				"Find shares",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return shares, nil
		}
		shares = append(shares, share)
	}
	return shares, nil
}

//not using
func (r *repo) DeleteOne(filter bson.M) error {
	deleteResult, err := trestCommon.DeleteOne(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"delete share",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("share not found(404)")
		trestCommon.ECLog3(
			"delete share",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}
