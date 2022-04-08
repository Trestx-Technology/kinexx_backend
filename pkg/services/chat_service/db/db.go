package db

import (
	"kinexx_backend/pkg/entity"
	chat "kinexx_backend/pkg/repository/chats"
	"kinexx_backend/pkg/services/profile_service/db"
	"kinexx_backend/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = chat.NewChatRepository("chats")
)

type chatService struct{}

func NewChatService(repository chat.ChatRepository) ChatService {
	repo = repository
	return &chatService{}
}

func (chat *chatService) SendMessage(message *entity.Message) ([]entity.Message, error) {
	message.ID = primitive.NewObjectID()
	_, err := repo.InsertOne(message)
	if err != nil {
		return []entity.Message{}, err
	}

	return chat.GetChats(message.SenderID, message.ReceiverID)
}

func (*chatService) GetChats(user string, receiver string) ([]entity.Message, error) {
	chats, err := repo.Find(bson.M{"$or": bson.A{bson.M{"sender_id": user, "receiver_id": receiver}, bson.M{"sender_id": receiver, "receiver_id": user}}}, bson.M{})
	if err != nil {
		return []entity.Message{}, nil
	}
	users, _ := db.GetProfilesForIDs([]string{user, receiver})
	for i := range chats {
		if chats[i].SenderID == users[0].ID.Hex() {
			chats[i].Sender = users[0]
		} else if chats[i].ReceiverID == users[0].ID.Hex() {
			chats[i].Receiver = users[0]
		}
		if chats[i].SenderID == users[1].ID.Hex() {
			chats[i].Sender = users[1]
		} else if chats[i].ReceiverID == users[1].ID.Hex() {
			chats[i].Receiver = users[1]
		}
		chats[i].ContentURL = utils.CreatePreSignedDownloadUrl(chats[i].ContentURL)
	}
	return chats, nil
}
func (*chatService) GetAllChatsStarted(user string) ([]entity.Message, error) {
	chats, err := repo.Find(bson.M{"sender_id": user}, bson.M{})
	if err != nil {
		return []entity.Message{}, nil
	}
	var userid = make(map[string]bool)
	var newChats []entity.Message
	var userArray []string
	for i := range chats {
		if _, ok := userid[chats[i].ReceiverID]; !ok {
			newChats = append(newChats, chats[i])
			userid[chats[i].ReceiverID] = true
			userArray = append(userArray, chats[i].ReceiverID)
			chats[i].ContentURL = utils.CreatePreSignedDownloadUrl(chats[i].ContentURL)
		}
	}
	users, _ := db.GetProfilesForIDs(userArray)
	for i := range newChats {
		for _, user := range users {
			if user.ID.Hex() == newChats[i].ReceiverID {
				newChats[i].Receiver = user
			}
		}
	}
	return newChats, nil
}
func (chat *chatService) DeleteChat(messageID, userID, receiverID string) ([]entity.Message, error) {
	id, err := primitive.ObjectIDFromHex(messageID)
	if err != nil {
		return chat.GetChats(userID, receiverID)
	}
	repo.DeleteOne(bson.M{"_id": id, "sender_id": userID})
	return chat.GetChats(userID, receiverID)
}
