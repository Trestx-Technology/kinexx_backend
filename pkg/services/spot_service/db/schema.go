package db

import "kinexx_backend/pkg/entity"

type SpotService interface {
	AddSpot(spot *entity.SpotDB) (string, error)
	GetAllSpots(groupID, creatorID string) ([]entity.SpotDB, error)
	GetSpot(spotID string) (entity.SpotDB, error)
	DeleteSpot(spotID string) (string, error)
}
