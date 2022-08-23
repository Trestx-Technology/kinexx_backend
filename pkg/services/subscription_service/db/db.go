package db

import (
	"kinexx_backend/pkg/services/subscription_service/entity"
	"kinexx_backend/pkg/services/subscription_service/subscriptions"
	"kinexx_backend/pkg/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = subcrriptions.NewSubscriptionsRepository("subscriptions")
)

type subscriptionService struct{}

// AddSubscription implements SubscriptionService
func (*subscriptionService) AddSubscription(subscription *entity.SubscriptionType) (string, error) {
	subscription.ID = primitive.NewObjectID()
	subscription.CreatedTime = time.Now()
	return repo.InsertOne(subscription)
}

// DeleteSubscription implements SubscriptionService
func (*subscriptionService) DeleteSubscription(subscriptionID string) (string, error) {
	id, err := primitive.ObjectIDFromHex(subscriptionID)
	if err != nil {
		return "", err
	}
	err = repo.DeleteOne(bson.M{"_id": id})
	return "", err
}

// GetAllSubscriptions implements SubscriptionService
func (*subscriptionService) GetAllSubscriptions() ([]entity.SubscriptionType, error) {
	res, err := repo.Find(bson.M{}, bson.M{})
	if err != nil {
		return nil, err
	}
	for i := range res {
		res[i].Icon = utils.CreatePreSignedDownloadUrl(res[i].Icon)
	}
	return res, nil
}

// GetSubscription implements SubscriptionService
func (*subscriptionService) GetSubscription(subscriptionID string) (entity.SubscriptionType, error) {
	id, err := primitive.ObjectIDFromHex(subscriptionID)
	if err != nil {
		return entity.SubscriptionType{}, err
	}
	res, err := repo.FindOne(bson.M{"_id": id}, bson.M{})
	if err != nil {
		return entity.SubscriptionType{}, err
	}
	res.Icon = utils.CreatePreSignedDownloadUrl(res.Icon)
	return res, nil
}

func NewSubscriptionService(repository subcrriptions.SubscriptionsRepository) SubscriptionService {
	repo = repository
	return &subscriptionService{}
}
