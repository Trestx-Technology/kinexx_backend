package db

import (
	"kinexx_backend/pkg/entity"
	"kinexx_backend/pkg/services/chat_service/chats"
	groupDB "kinexx_backend/pkg/services/group_service/db"
	"kinexx_backend/pkg/services/profile_service/db"
	"kinexx_backend/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = chats.NewChatRepository("chats")
)

type chatService struct{}

// GetSpotsChat implements ChatService
func (*chatService) GetSpotsChat(spotID string) ([]entity.Message, error) {
	chats, err := repo.Find(bson.M{"receiver_id": spotID}, bson.M{})
	if err != nil {
		return []entity.Message{}, nil
	}
	var userArray []string
	for i := range chats {
		userArray = append(userArray, chats[i].SenderID)
	}
	users, _ := db.GetProfilesForIDs(userArray)

	for i := range chats {
		for _, user := range users {
			if chats[i].SenderID == user.ID.Hex() {
				chats[i].Sender = user
			}
		}
		chats[i].ContentURL = utils.CreatePreSignedDownloadUrl(chats[i].ContentURL)
	}
	return chats, nil
}

func NewChatService(repository chats.ChatRepository) ChatService {
	repo = repository
	return &chatService{}
}

func (chat *chatService) SendMessage(message *entity.Message) ([]entity.Message, error) {
	message.ID = primitive.NewObjectID()
	_, err := repo.InsertOne(message)
	if err != nil {
		return []entity.Message{}, err
	}
	if message.Group {
		return chat.GetSpotsChat(message.ReceiverID)

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
	chats, err := repo.Find(bson.M{"$or": bson.A{bson.M{"sender_id": user, "hide_for_receiver": false}, bson.M{"receiver_id": user, "hide_for_sender": false}}}, bson.M{})
	if err != nil {
		return []entity.Message{}, nil
	}
	var userid = make(map[string]bool)
	var newChats []entity.Message
	var userArray []string
	var spotArray []string
	userid[user] = true
	for i := range chats {
		if _, ok := userid[chats[i].ReceiverID]; !ok {
			newChats = append(newChats, chats[i])
			userid[chats[i].ReceiverID] = true
			if chats[i].Group {
				spotArray = append(spotArray, chats[i].ReceiverID)
			} else {
				userArray = append(userArray, chats[i].ReceiverID)
			}
		}
		if _, ok := userid[chats[i].SenderID]; !ok {
			newChats = append(newChats, chats[i])
			userid[chats[i].SenderID] = true
			userArray = append(userArray, chats[i].SenderID)
		}
		chats[i].ContentURL = utils.CreatePreSignedDownloadUrl(chats[i].ContentURL)

	}
	users, _ := db.GetProfilesForIDs(userArray)
	spots, _ := groupDB.GetGroupByIDs(spotArray)
	for i := range newChats {
		for _, user := range users {
			if user.ID.Hex() == newChats[i].ReceiverID {
				newChats[i].Receiver = user
			}
			if user.ID.Hex() == newChats[i].SenderID {
				newChats[i].Receiver = user
			}
		}
		for _, spot := range spots {
			if spot.ID.Hex() == newChats[i].ReceiverID {
				newChats[i].Squad = spot
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
func (chat *chatService) HideChat(userID, otherID string) ([]entity.Message, error) {
	repo.UpdateOne(bson.M{"sender_id": userID, "receiver_id": otherID}, bson.M{"$set": bson.M{"hide_for_receiver": true}})
	repo.UpdateOne(bson.M{"receiver_id": userID, "sender_id": otherID}, bson.M{"$set": bson.M{"hide_for_sender": true}})
	return chat.GetAllChatsStarted(userID)
}
