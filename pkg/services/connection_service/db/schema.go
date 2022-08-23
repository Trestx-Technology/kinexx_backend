package db

import (
	"kinexx_backend/pkg/services/connection_service/entity"
)

type ConnectionService interface {
	AddConnection(post *entity.ConnectionDB) (string, error)
	UpdateConnection(post *entity.ConnectionDB) (string, error)
	GetConnectionByID(userID string) ([]entity.ConnectionDB, error)
	GetConnectionCountByID(userID string) (int, error)
	GetOnlineConnectionByID(user string) ([]entity.ConnectionDB, error)
}
