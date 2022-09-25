package db

import (
	"kinexx_backend/pkg/services/spot_service/entity"
)

type SpotService interface {
	AddSpot(spot *entity.SpotDB) (string, error)
	GetAllSpots(groupID, creatorID string) ([]entity.SpotDB, error)
	GetSpot(spotID string) (entity.SpotDB, error)
	DeleteSpot(spotID string) (string, error)
	Count() (int64, error)
}
