package db

import "kinexx_backend/pkg/entity"

type ChatService interface {
	SendMessage(message *entity.Message) ([]entity.Message, error)
	GetAllChatsStarted(userID string) ([]entity.Message, error)
	GetChats(userID string, sender string) ([]entity.Message, error)
	GetSpotsChat(spotID string) ([]entity.Message, error)
	DeleteChat(messageID, user, receiver string) ([]entity.Message, error)
	HideChat(userID, otherID string) ([]entity.Message, error)
}
