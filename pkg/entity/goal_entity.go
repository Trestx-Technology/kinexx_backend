package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GoalDB struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	Icon          string             `bson:"icon" json:"icon"`
	Name          string             `bson:"name" json:"name"`
	Description   string             `bson:"description" json:"description"`
	Banner        string             `bson:"banner" json:"banner"`
	PromoVideos   []string           `bson:"promo_videos" json:"promo_videos"`
	CreatedDate   time.Time          `bson:"created_date" json:"created_date"`
	CreatorUserID string             `bson:"creator_user_id" json:"creator_user_id"`
	GroupIDList   []string           `json:"group_id_list"`
	TimePeriod    string             `bson:"time_period" json:"time_period"`
}

type GoalGroupDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	GroupID     string             `bson:"group_id" json:"group_id"`
	GoalID      string             `bson:"goal_id" json:"goal_id"`
	Type        string             `bson:"type" json:"type"`
	CreatedDate time.Time          `bson:"created_date" json:"created_date"`
}
