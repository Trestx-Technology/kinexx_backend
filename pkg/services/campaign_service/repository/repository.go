package repository

import (
	campaignEntity "kinexx_backend/pkg/services/campaign_service/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type CampaignRepository interface {
	InsertOne(document interface{}) (string, error)
	FindOne(filter, projection bson.M) (campaignEntity.Campaign, error)
	Find(filter, projection bson.M) ([]campaignEntity.Campaign, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
	Count(filter bson.M) (int64, error)
}
