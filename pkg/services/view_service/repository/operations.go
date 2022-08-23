package viewRepository

import (
	"context"
	"errors"
	trestCommon "github.com/Trestx-technology/trestx-common-go-lib"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	viewEntity "kinexx_backend/pkg/services/view_service/entity"
)

type repo struct {
	CollectionName string
}

//NewFirestoreRepository creates a new repo
func NewRepository(collectionName string) Repository {
	return &repo{
		CollectionName: collectionName,
	}
}

//used by signup
func (r *repo) InsertOne(document interface{}) (string, error) {
	user, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert ViewContent",
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

//used by update ViewContent ,login and email verifcation
func (r *repo) UpdateOne(filter, update bson.M) (string, error) {
	result, err := trestCommon.UpdateOne(filter, update, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"update ViewContent",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("ViewContent not found(404)")
		trestCommon.ECLog3(
			"update ViewContent",
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

//used by get ViewContent ,login and email verification
func (r *repo) FindOne(filter, projection bson.M) (viewEntity.ViewDB, error) {
	var ViewContent viewEntity.ViewDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&ViewContent)
	if err != nil {
		trestCommon.ECLog3(
			"Find ViewContent",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return ViewContent, err
	}
	return ViewContent, err
}

//not used may use in future for gettin list of ViewContents
func (r *repo) Find(filter, projection bson.M) ([]viewEntity.ViewDB, error) {
	var viewContents []viewEntity.ViewDB
	cursor, err := trestCommon.FindSort(filter, projection, bson.M{"_id": -1}, 1000000, 0, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find ViewContents",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var viewContent viewEntity.ViewDB
		if err = cursor.Decode(&viewContent); err != nil {
			trestCommon.ECLog3(
				"Find ViewContents",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return viewContents, nil
		}
		viewContents = append(viewContents, viewContent)
	}
	return viewContents, nil
}

//not using
func (r *repo) DeleteOne(filter bson.M) error {
	deleteResult, err := trestCommon.DeleteOne(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"delete ViewContent",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("ViewContent not found(404)")
		trestCommon.ECLog3(
			"delete ViewContent",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}
func (r *repo) Aggregate(pipeline bson.A) ([]viewEntity.ViewDB, error) {
	var viewContents []viewEntity.ViewDB
	cursor, err := trestCommon.Aggregate(pipeline, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find ViewContents",
			err,
			logrus.Fields{
				"filter":          pipeline,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var viewContent viewEntity.ViewDB
		if err = cursor.Decode(&viewContent); err != nil {
			trestCommon.ECLog3(
				"Find ViewContents",
				err,
				logrus.Fields{
					"filter":          pipeline,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return viewContents, nil
		}
		viewContents = append(viewContents, viewContent)
	}
	return viewContents, nil
}
