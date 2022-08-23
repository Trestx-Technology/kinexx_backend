package notification_db

import (
	"kinexx_backend/pkg/entity"
	"kinexx_backend/pkg/services/notification_service/notifications"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = notifications.NewNotificationRepository("notifications")
)

type notificationService struct{}

func NewNotificationService(repository notifications.NotificationRepository) NotificationService {
	repo = repository
	return &notificationService{}
}

// GetNotifications
// return repo.Find(bson.M{"receiver_id":userID}, bson.M{})
func (*notificationService) GetNotifications(ReceiverID string) ([]entity.NotiFicationMessage, error) {
	// getting from database
	return repo.Find(bson.M{"receiver_id": ReceiverID}, bson.M{})
}

func (*notificationService) UpdateNotifications(notificationID string) (string, error) {
	id, err := primitive.ObjectIDFromHex(notificationID)
	if err != nil {
		return "", err
	}
	filter := bson.M{"_id": id}
	set := bson.M{"read": true}
	// read, sentTime
	set["updated_time"] = time.Now()

	return repo.UpdateOne(filter, bson.M{"$set": set})
}

func AddNotification(notification *entity.NotiFicationMessage) (string, error) {
	notification.ID = primitive.NewObjectID()
	notification.SentTime = time.Now()
	return repo.InsertOne(notification)
}

// update
