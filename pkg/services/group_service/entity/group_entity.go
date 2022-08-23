package entity

import (
	"kinexx_backend/pkg/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GroupDB struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	Name           string             `bson:"name" json:"name"`
	Description    string             `bson:"description" json:"description"`
	Public         bool               `bson:"public" json:"public"`
	Visible        string             `bson:"visible" json:"visible"`
	CreatedDate    time.Time          `bson:"created_date" json:"created_date"`
	Banner         string             `bson:"banner" json:"banner"`
	Logo           string             `bson:"logo" json:"logo"`
	PromoVideos    []string           `bson:"promo_videos" json:"promo_videos"`
	CreatorUserID  string             `bson:"creator_user_id" json:"creator_user_id"`
	CreatorDetails entity.ProfileDB   `json:"creator_details"`
	Status         string             `bson:"status" json:"status"`
	UserIDList     []string           `json:"user_id_list"`
	GoalIDList     []string           `json:"goal_id_list"`
}

type GroupUserDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	GroupID     string             `bson:"group_id" json:"group_id"`
	UserID      string             `bson:"user_id" json:"user_id"`
	AccessType  string             `bson:"access_type" json:"access_type"`
	Subscribed  string             `bson:"subscribed" json:"subscribed"`
	Status      string             `bson:"status" json:"status"`
	CreatedDate time.Time          `bson:"created_date" json:"created_date"`
}
