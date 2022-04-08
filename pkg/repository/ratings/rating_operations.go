package ratings

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

func NewRatingRepository(collectionName string) RatingRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

func (r *repo) InsertOne(document interface{}) (string, error) {
	user, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert rating",
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

func (r *repo) UpdateOne(filter, update bson.M) (string, error) {
	result, err := trestCommon.UpdateOne(filter, update, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"update rating",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("rating not found(404)")
		trestCommon.ECLog3(
			"update rating",
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

func (r *repo) FindOne(filter, projection bson.M) (entity.RatingDB, error) {
	var rating entity.RatingDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&rating)
	if err != nil {
		trestCommon.ECLog3(
			"Find rating",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return rating, err
	}
	return rating, err
}

func (r *repo) Find(filter, projection bson.M) ([]entity.RatingDB, error) {
	var ratings []entity.RatingDB
	cursor, err := trestCommon.FindSort(filter, projection, bson.M{"_id": -1}, 1000000, 0, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find ratings",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var rating entity.RatingDB
		if err = cursor.Decode(&rating); err != nil {
			trestCommon.ECLog3(
				"Find ratings",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return ratings, nil
		}
		ratings = append(ratings, rating)
	}
	return ratings, nil
}

func (r *repo) DeleteOne(filter bson.M) error {
	deleteResult, err := trestCommon.DeleteOne(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"delete rating",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("rating not found(404)")
		trestCommon.ECLog3(
			"delete rating",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}
