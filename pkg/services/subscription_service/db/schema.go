package db

import (
	"kinexx_backend/pkg/services/subscription_service/entity"
)

type SubscriptionService interface {
	AddSubscription(subscription *entity.SubscriptionType) (string, error)
	GetAllSubscriptions() ([]entity.SubscriptionType, error)
	GetSubscription(subscriptionID string) (entity.SubscriptionType, error)
	DeleteSubscription(subscriptionID string) (string, error)
}
