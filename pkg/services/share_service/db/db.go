package db

import (
	"kinexx_backend/pkg/services/profile_service/db"
	"kinexx_backend/pkg/services/share_service/entity"
	"kinexx_backend/pkg/services/share_service/share"
	"kinexx_backend/pkg/utils"

	"strings"

	"errors"
	"time"

	trestCommon "github.com/Trestx-technology/trestx-common-go-lib"
	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = shares.NewShareRepository("shares")
)

type shareService struct{}

func NewShareService(repository shares.ShareRepository) ShareService {
	repo = repository
	return &shareService{}
}

func (*shareService) AddShare(share *entity.ShareDB) (string, error) {
	if share.UserID == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"inserted share section",
			err,
		)
		return "", err
	}
	share.ID = primitive.NewObjectID()
	share.CreatedTime = time.Now()
	if share.Type == "" {
		share.Type = "POST"
	}
	share.Status = "shared"

	result, err := repo.InsertOne(share)
	user, _ := db.GetProfilesForIDs([]string{share.UserID})
	body := "Shared " + share.Type + "with you"
	if len(user) > 0 {
		body = user[0].Name + " shared " + share.Type + "with you"
	}
	if strings.Contains(share.ReceiverID, ",") {
		rec_ids := strings.Split(share.ReceiverID, ",")
		for _, rec := range rec_ids {
			utils.SendNotification(rec, share.Type, share.SharedItemID, body, share.UserID)
		}
	} else if share.ReceiverID != "" {
		utils.SendNotification(share.ReceiverID, share.Type, share.SharedItemID, body, share.UserID)
	}
	if err != nil {
		trestCommon.ECLog3(
			"inserted share section",
			err,
			logrus.Fields{
				"share": share,
			})
		return "", err
	}

	return result, nil
}

func (*shareService) GetShareByID(user, shareType string) ([]entity.ShareDB, error) {
	filter := bson.M{}
	if shareType == "" {
		filter["$or"] = bson.A{bson.M{"type": ""}, bson.M{"type": "POST"}}
	} else {
		filter["type"] = shareType
	}
	filter["receiver_id"] = user
	shares, err := repo.Find(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetShare section",
			err,
		)
		return []entity.ShareDB{}, err
	}
	if shareType == "ECARD" {
		var userID []string
		for _, share := range shares {
			userID = append(userID, share.SharedItemID)
		}
		user, _ := db.GetProfilesForIDs(userID)
		for i := range shares {
			for _, use := range user {
				if use.ID.Hex() == shares[i].SharedItemID {
					shares[i].Receiver = use
				}
			}
		}
	}
	return shares, err
}
func (*shareService) GetMyShare(user, shareType string) ([]entity.ShareDB, error) {
	filter := bson.M{}
	if shareType == "" {
		filter["$or"] = bson.A{bson.M{"type": ""}, bson.M{"type": "POST"}}
	} else {
		filter["type"] = shareType
	}
	filter["user_id"] = user

	shares, err := repo.Find(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetShare section",
			err,
		)
		return []entity.ShareDB{}, err
	}
	if shareType == "ECARD" {
		var userID []string
		for _, share := range shares {
			userID = append(userID, share.SharedItemID)
		}
		user, _ := db.GetProfilesForIDs(userID)
		for i := range shares {
			for _, use := range user {
				if use.ID.Hex() == shares[i].SharedItemID {
					shares[i].Receiver = use
				}
			}
		}
	}
	return shares, err
}
func GetShareByPostID(post string) ([]string, error) {

	shares, err := repo.Find(bson.M{"shared_item_id": post, "$or": bson.A{bson.M{"type": ""}, bson.M{"type": "POST"}}}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetShare section",
			err,
		)
		return []string{}, err
	}
	var userID []string
	for _, share := range shares {
		userID = append(userID, share.UserID)
	}
	return userID, err
}

func (*shareService) UpdateShare(*entity.ShareDB) (string, error) {

	return "", nil
}
