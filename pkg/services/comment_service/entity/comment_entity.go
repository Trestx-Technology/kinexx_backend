package entity

import (
	"kinexx_backend/pkg/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Comment represents the model for an comment
type CommentDB struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	PostID       string             `bson:"post_id" json:"post_id"`
	UserID       string             `bson:"user_id" json:"user_id"`
	Status       string             `bson:"status" json:"status"`
	CreatedTime  time.Time          `bson:"created_time" json:"created_time"`
	UpdatedTime  time.Time          `bson:"updated_time" json:"updated_time"`
	Body         string             `bson:"body" json:"body"`
	ContentType  string             `bson:"content_type" json:"content_type"`
	ContentURL   string             `bson:"content_url" json:"content_url"`
	Tags         string             `bson:"tags" json:"tags"`
	Likes        int                `bson:"likes" json:"likes"`
	LikedBy      []string           `bson:"liked_by" json:"liked_by"`
	LikedByUsers []entity.ProfileDB `bson:"liked_by_users" json:"liked_by_users"`
	User         entity.ProfileDB   `json:"user"`
	Comment      []CommentDB        `json:"comment"`
	CommentID    string             `json:"comment_id"`
}
