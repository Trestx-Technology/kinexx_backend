package notification_db

import (
	"kinexx_backend/pkg/services/notification_service/entity"
)

type NotificationService interface {
	UpdateNotifications(string) (string, error)
	GetNotifications(userID string) ([]entity.NotiFicationMessage, error)
}
