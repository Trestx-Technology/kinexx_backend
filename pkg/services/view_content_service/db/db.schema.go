package viewContentDB

import (
	viewContentEntity "kinexx_backend/pkg/services/view_content_service/entity"
)

type Service interface {
	Add(brand *viewContentEntity.ViewContentDB) (viewContentEntity.ViewContentDB, error)
	Update(brand *viewContentEntity.ViewContentDB, campaignId string) (viewContentEntity.ViewContentDB, error)
	GetForCampaign(campaignID string) ([]viewContentEntity.ViewContentDB, error)
	Delete(brandID string) (string, error)
}
