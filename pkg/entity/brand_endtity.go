package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BrandDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Logo        string             `bson:"logo" json:"logo"`
	Banner      string             `bson:"banner" json:"banner"`
	BrandName   string             `bson:"brand_name" json:"brand_name"`
	Price       float64            `bson:"price" json:"price"`
	PromoVideos []string           `bson:"promo_videos" json:"promo_videos"`
	CreatorID   string             `bson:"creator_id" json:"creator_id"`
	CreatedTime time.Time          `bson:"created_time" json:"created_time"`
	Status      string             `bson:"status" json:"status"`
}
