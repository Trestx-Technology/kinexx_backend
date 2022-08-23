package viewContentEntity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	entity2 "kinexx_backend/pkg/entity"
	"kinexx_backend/pkg/services/comment_service/entity"
	"time"
)

type ViewContentDB struct {
	ID             primitive.ObjectID  `bson:"_id" json:"id"`
	ViewContentID  string              `bson:"view_content_id" json:"view_content_id"`
	ViewID         string              `bson:"view_id" json:"view_id"`
	UserID         string              `bson:"user_id" json:"user_id"`
	Status         int                 `bson:"status" json:"status"`
	CreatedTime    time.Time           `bson:"created_time" json:"created_time"`
	UpdatedTime    time.Time           `bson:"updated_time" json:"updated_time"`
	Banner         string              `bson:"banner" json:"banner"`
	Cover          string              `bson:"cover" json:"cover"`
	VideoURL       string              `bson:"video_url" json:"video_url"`
	Description    string              `bson:"description" json:"description"`
	Tags           string              `bson:"tags" json:"tags"`
	Likes          int                 `bson:"likes" json:"likes"`
	LikedBy        []string            `bson:"liked_by" json:"liked_by"`
	LikedByUsers   []entity2.ProfileDB `bson:"liked_by_users" json:"liked_by_users"`
	SharedByUsers  []entity2.ProfileDB `bson:"shared_by_users" json:"shared_by_users"`
	User           entity2.ProfileDB   `json:"user"`
	Comment        []entity.CommentDB  `json:"comment"`
	PostID         string              `json:"post_id"`
	SharedInstance bool                `json:"shared_instance"`
	ShareCount     int                 `json:"shared_count"`
}
