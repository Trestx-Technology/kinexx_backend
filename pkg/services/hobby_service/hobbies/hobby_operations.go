package hobbies

import (
	"context"
	"errors"
	"kinexx_backend/pkg/services/hobby_service/entity"

	trestCommon "github.com/Trestx-technology/trestx-common-go-lib"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type repo struct {
	CollectionName string
}

//NewFirestoreRepository creates a new repo
func NewHobbiesRepository(collectionName string) HobbiesRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

//used by signup
func (r *repo) InsertOne(document interface{}) (string, error) {
	user, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert hobbies",
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

//used by update hobbies ,login and email verifcation
func (r *repo) UpdateOne(filter, update bson.M) (string, error) {
	result, err := trestCommon.UpdateOne(filter, update, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"update hobbies",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("hobbies not found(404)")
		trestCommon.ECLog3(
			"update hobbies",
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

//used by get hobbies ,login and email verification
func (r *repo) FindOne(filter, projection bson.M) (entity.HobbiesDB, error) {
	var hobbies entity.HobbiesDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&hobbies)
	if err != nil {
		trestCommon.ECLog3(
			"Find hobbies",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return hobbies, err
	}
	return hobbies, err
}

//not used may use in future for gettin list of hobbiess
func (r *repo) Find(filter, projection bson.M) ([]entity.HobbiesDB, error) {
	var hobbiess []entity.HobbiesDB
	cursor, err := trestCommon.FindSort(filter, projection, bson.M{"_id": -1}, 1000000, 0, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find hobbiess",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var hobbies entity.HobbiesDB
		if err = cursor.Decode(&hobbies); err != nil {
			trestCommon.ECLog3(
				"Find hobbiess",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return hobbiess, nil
		}
		hobbiess = append(hobbiess, hobbies)
	}
	return hobbiess, nil
}

//not using
func (r *repo) DeleteOne(filter bson.M) error {
	deleteResult, err := trestCommon.DeleteOne(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"delete hobbies",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("hobbies not found(404)")
		trestCommon.ECLog3(
			"delete hobbies",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}
