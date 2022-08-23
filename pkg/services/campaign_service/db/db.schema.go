package campaignDB

import (
	campaignEntity "kinexx_backend/pkg/services/campaign_service/entity"
)

type CampaignService interface {
	AddCampaign(brand *campaignEntity.Campaign) (campaignEntity.Campaign, error)
	UpdateCampaign(brand *campaignEntity.Campaign, campaignId string) (campaignEntity.Campaign, error)
	GetMyCampaign(userID string) ([]campaignEntity.Campaign, error)
	GetAllCampaign() ([]campaignEntity.Campaign, error)
	SearchCampaign(search string) ([]campaignEntity.Campaign, error)
	GetCampaign(brandID string) (campaignEntity.Campaign, error)
	DeleteCampaign(brandID string) (string, error)
}
