package stores

import (
	"context"
	"errors"
	"kinexx_backend/pkg/services/store_service/entity"

	trestCommon "github.com/Trestx-technology/trestx-common-go-lib"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type repo struct {
	CollectionName string
}

//NewFirestoreRepository creates a new repo
func NewStoreRepository(collectionName string) StoreRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

//used by signup
func (r *repo) InsertOne(document interface{}) (string, error) {
	user, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert Store",
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

//used by update Store ,login and email verifcation
func (r *repo) UpdateOne(filter, update bson.M) (string, error) {
	result, err := trestCommon.UpdateOne(filter, update, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"update Store",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("Store not found(404)")
		trestCommon.ECLog3(
			"update Store",
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

//used by get Store ,login and email verification
func (r *repo) FindOne(filter, projection bson.M) (entity.ShopDB, error) {
	var Store entity.ShopDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&Store)
	if err != nil {
		trestCommon.ECLog3(
			"Find Store",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return Store, err
	}
	return Store, err
}

//not used may use in future for gettin list of stores
func (r *repo) Find(filter, projection bson.M) ([]entity.ShopDB, error) {
	var stores []entity.ShopDB
	cursor, err := trestCommon.FindSort(filter, projection, bson.M{"_id": -1}, 1000000, 0, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find stores",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var Store entity.ShopDB
		if err = cursor.Decode(&Store); err != nil {
			trestCommon.ECLog3(
				"Find stores",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return stores, nil
		}
		stores = append(stores, Store)
	}
	return stores, nil
}

//not using
func (r *repo) DeleteOne(filter bson.M) error {
	deleteResult, err := trestCommon.DeleteOne(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"delete Store",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("Store not found(404)")
		trestCommon.ECLog3(
			"delete Store",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}
func (r *repo) Count(filter bson.M) (int64, error) {
	return trestCommon.Count(filter, bson.M{}, r.CollectionName)
}
