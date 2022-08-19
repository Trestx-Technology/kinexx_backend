package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CourseDB struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	UserID         string             `bson:"user_id" json:"user_id"`
	Status         string             `bson:"status" json:"status"`
	CreatedTime    time.Time          `bson:"created_time" json:"created_time"`
	UpdatedTime    time.Time          `bson:"updated_time" json:"updated_time"`
	Banner         string             `bson:"banner" json:"banner"`
	Cover          string             `bson:"cover" json:"cover"`
	ContentType    string             `bson:"content_type" json:"content_type"`
	ContentURL     string             `bson:"content_url" json:"content_url"`
	Tags           string             `bson:"tags" json:"tags"`
	Likes          int                `bson:"likes" json:"likes"`
	LikedBy        []string           `bson:"liked_by" json:"liked_by"`
	LikedByUsers   []ProfileDB        `bson:"liked_by_users" json:"liked_by_users"`
	SharedByUsers  []ProfileDB        `bson:"shared_by_users" json:"shared_by_users"`
	User           ProfileDB          `json:"user"`
	Comment        []CommentDB        `json:"comment"`
	PostID         string             `json:"post_id"`
	SharedInstance bool               `json:"shared_instance"`
	ShareCount     int                `json:"shared_count"`
}
