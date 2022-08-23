package viewEntity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kinexx_backend/pkg/entity"
	viewContentEntity "kinexx_backend/pkg/services/view_content_service/entity"
	"time"
)

type ViewDB struct {
	ID             primitive.ObjectID                `bson:"_id" json:"id"`
	ViewID         string                            `bson:"view_id" json:"view_id"`
	UserID         string                            `bson:"user_id" json:"user_id"`
	Status         int                               `bson:"status" json:"status"`
	CreatedTime    time.Time                         `bson:"created_time" json:"created_time"`
	UpdatedTime    time.Time                         `bson:"updated_time" json:"updated_time"`
	Banner         string                            `bson:"banner" json:"banner"`
	Cover          string                            `bson:"cover" json:"cover"`
	Trailer        string                            `bson:"trailer" json:"trailer"`
	Tags           string                            `bson:"tags" json:"tags"`
	Description    string                            `bson:"description" json:"description"`
	Likes          int                               `bson:"likes" json:"likes"`
	LikedBy        []string                          `bson:"liked_by" json:"liked_by"`
	LikedByUsers   []entity.ProfileDB                `bson:"liked_by_users" json:"liked_by_users"`
	SharedByUsers  []entity.ProfileDB                `bson:"shared_by_users" json:"shared_by_users"`
	User           entity.ProfileDB                  `json:"user"`
	Comment        []entity.CommentDB                `json:"comment"`
	Content        []viewContentEntity.ViewContentDB `json:"content"`
	PostID         string                            `json:"post_id"`
	SharedInstance bool                              `json:"shared_instance"`
	ShareCount     int                               `json:"shared_count"`
	Clicked        int                               `bson:"clicked" json:"clicked"`
	ClickedBy      []string                          `bson:"clicked_by" json:"clicked_by"`
}
