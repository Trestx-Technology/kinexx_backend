package db

import (
	"kinexx_backend/pkg/entity"
	connection "kinexx_backend/pkg/repository/connections"
	"kinexx_backend/pkg/services/profile_service/db"
	"kinexx_backend/pkg/utils"
	"strings"

	"errors"
	"time"

	"github.com/aekam27/trestCommon"

	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = connection.NewConnectionRepository("connections")
)

type connectionService struct{}

func NewConnectionService(repository connection.ConnectionRepository) ConnectionService {
	repo = repository
	return &connectionService{}
}

func (*connectionService) AddConnection(connection *entity.ConnectionDB) (string, error) {
	if connection.UserID == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"inserted connection section",
			err,
		)
		return "", err
	}
	connection.ID = primitive.NewObjectID()
	connection.CreatedTime = time.Now()
	connection.Status = "initiated"
	result, err := repo.InsertOne(connection)
	if err != nil {
		trestCommon.ECLog3(
			"inserted connection section",
			err,
			logrus.Fields{
				"connection": connection,
			})
		return "", err
	}
	users, err := db.GetProfilesForIDs([]string{connection.UserID})
	if err == nil && len(users) > 0 {
		utils.SendNotification(connection.ReceiverID, "CONNECTION_REQUEST", connection.ID.Hex(), users[0].Name+" sent you a connection request", users[0].ID.Hex())
	}
	return result, nil
}

func (*connectionService) GetConnectionByID(user string) ([]entity.ConnectionDB, error) {

	connection, err := repo.Find(bson.M{"$or": bson.A{bson.M{"user_id": user}, bson.M{"receiver_id": user}}, "type": "friend"}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetConnection section",
			err,
		)
		return []entity.ConnectionDB{}, err
	}
	var userIDs []string
	for connect := range connection {
		if connection[connect].UserID != user {
			userIDs = append(userIDs, connection[connect].UserID)
		} else {
			userIDs = append(userIDs, connection[connect].ReceiverID)
		}
	}
	users, err := db.GetProfilesForIDs(userIDs)
	for _, us := range users {
		for connect := range connection {
			newUrl := createPreSignedDownloadUrl(connection[connect].ContentURL)
			connection[connect].ContentURL = newUrl
			if connection[connect].UserID == us.ID.Hex() || connection[connect].ReceiverID == us.ID.Hex() {
				connection[connect].User = append(connection[connect].User, us)
			}
		}
	}
	return connection, err
}

func (*connectionService) GetConnectionCountByID(user string) (int, error) {

	connection, err := repo.Find(bson.M{"$or": bson.A{bson.M{"user_id": user}, bson.M{"receiver_id": user}}, "type": "friend", "status": "accepted"}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetConnection section",
			err,
		)
		return 0, err
	}
	return len(connection), err
}
func (*connectionService) GetOnlineConnectionByID(user string) ([]entity.ConnectionDB, error) {

	connection, err := repo.Find(bson.M{"$or": bson.A{bson.M{"user_id": user}, bson.M{"receiver_id": user}}, "type": "friend"}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetConnection section",
			err,
		)
		return []entity.ConnectionDB{}, err
	}
	var userIDs []string
	for connect := range connection {
		if connection[connect].UserID != user {
			userIDs = append(userIDs, connection[connect].UserID)
		} else {
			userIDs = append(userIDs, connection[connect].ReceiverID)
		}
	}
	users, err := db.GetProfilesForIDs(userIDs)
	var onLineConnection []entity.ConnectionDB
	for _, us := range users {
		for connect := range connection {
			newUrl := createPreSignedDownloadUrl(connection[connect].ContentURL)
			connection[connect].ContentURL = newUrl
			if us.Online && connection[connect].UserID == us.ID.Hex() || connection[connect].ReceiverID == us.ID.Hex() {
				connection[connect].User = append(connection[connect].User, us)
				onLineConnection = append(onLineConnection, connection[connect])
			}
		}
	}
	return onLineConnection, err
}
func (*connectionService) UpdateConnection(connection *entity.ConnectionDB) (string, error) {
	id, err := primitive.ObjectIDFromHex(connection.ConnectionID)
	if err != nil {
		return "", err
	}
	filter := bson.M{"_id": id}
	con, err := repo.FindOne(filter, bson.M{})
	if err != nil {
		return "", err
	}
	set := bson.M{}
	set["message_by_receiver"] = connection.MessageByReceiver
	set["content_type_by_receiver"] = connection.ContentTypeByReceiver
	set["content_url_by_receiver"] = connection.ContentURLByReceiver
	set["status"] = connection.Status

	set["updated_time"] = time.Now()
	data, err := repo.UpdateOne(filter, bson.M{"$set": set})
	if err != nil {
		return data, err
	}
	users, err := db.GetProfilesForIDs([]string{con.UserID})
	if err == nil {
		utils.SendNotification(con.ReceiverID, "CONNECTION_REQUEST", connection.ID.Hex(), users[0].Name+" "+connection.Status+" your  connection request", users[0].ID.Hex())
	}
	return data, nil
}

func createPreSignedDownloadUrl(url string) string {
	s := strings.Split(url, "?")
	if len(s) > 0 {
		o := strings.Split(s[0], "/")
		if len(o) > 3 {
			fileName := o[4]
			path := o[3]
			downUrl, _ := trestCommon.PreSignedDownloadUrlAWS(fileName, path)
			return downUrl
		}
	}
	return ""
}
