package subcrriptions

import (
	"context"
	"errors"
	"kinexx_backend/pkg/services/subscription_service/entity"

	trestCommon "github.com/Trestx-technology/trestx-common-go-lib"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type repo struct {
	CollectionName string
}

//NewFirestoreRepository creates a new repo
func NewSubscriptionsRepository(collectionName string) SubscriptionsRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

//used by signup
func (r *repo) InsertOne(document interface{}) (string, error) {
	user, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert Subscriptions",
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

//used by update Subscriptions ,login and email verifcation
func (r *repo) UpdateOne(filter, update bson.M) (string, error) {
	result, err := trestCommon.UpdateOne(filter, update, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"update Subscriptions",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("Subscriptions not found(404)")
		trestCommon.ECLog3(
			"update Subscriptions",
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

//used by get Subscriptions ,login and email verification
func (r *repo) FindOne(filter, projection bson.M) (entity.SubscriptionType, error) {
	var Subscriptions entity.SubscriptionType
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&Subscriptions)
	if err != nil {
		trestCommon.ECLog3(
			"Find Subscriptions",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return Subscriptions, err
	}
	return Subscriptions, err
}

//not used may use in future for gettin list of Subscriptionss
func (r *repo) Find(filter, projection bson.M) ([]entity.SubscriptionType, error) {
	var Subscriptionss []entity.SubscriptionType
	cursor, err := trestCommon.FindSort(filter, projection, bson.M{"_id": -1}, 1000000, 0, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find Subscriptionss",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var Subscriptions entity.SubscriptionType
		if err = cursor.Decode(&Subscriptions); err != nil {
			trestCommon.ECLog3(
				"Find Subscriptionss",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return Subscriptionss, nil
		}
		Subscriptionss = append(Subscriptionss, Subscriptions)
	}
	return Subscriptionss, nil
}

//not using
func (r *repo) DeleteOne(filter bson.M) error {
	deleteResult, err := trestCommon.DeleteOne(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"delete Subscriptions",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("Subscriptions not found(404)")
		trestCommon.ECLog3(
			"delete Subscriptions",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}
