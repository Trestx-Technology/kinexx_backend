package charities

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

func NewCharityRepository(collectionName string) CharityRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

func (r *repo) InsertOne(document interface{}) (string, error) {
	charity, err := trestCommon.InsertOne(document, r.CollectionName)
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
	charityid := charity.InsertedID.(primitive.ObjectID).Hex()
	return charityid, nil
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
func (r *repo) FindOne(filter, projection bson.M) (entity.CharityDB, error) {
	var charity entity.CharityDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&charity)
	if err != nil {
		trestCommon.ECLog3(
			"Find follow",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return charity, err
	}
	return charity, err
}

//not used may use in future for gettin list of follows
func (r *repo) Find(filter, projection bson.M) ([]entity.CharityDB, error) {
	var charitys []entity.CharityDB
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
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var charity entity.CharityDB
		if err = cursor.Decode(&charity); err != nil {
			trestCommon.ECLog3(
				"Find follows",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return charitys, nil
		}
		charitys = append(charitys, charity)
	}
	return charitys, nil
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
