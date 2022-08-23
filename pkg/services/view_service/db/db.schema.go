package viewDB

import (
	viewEntity "kinexx_backend/pkg/services/view_service/entity"
)

type Service interface {
	Add(view *viewEntity.ViewDB) (viewEntity.ViewDB, error)
	Update(view *viewEntity.ViewDB, campaignId string) (viewEntity.ViewDB, error)
	GetMy(userID string) ([]viewEntity.ViewDB, error)
	GetAll() ([]viewEntity.ViewDB, error)
	Get(viewID string, userID string) ([]viewEntity.ViewDB, error)
	Delete(viewID string) (string, error)
}
