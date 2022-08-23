package entity

import (
	"kinexx_backend/pkg/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShopDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Logo        string             `bson:"logo" json:"logo"`
	Banner      string             `bson:"banner" json:"banner"`
	StoreName   string             `bson:"store_name" json:"store_name"`
	SpotID      string             `bson:"spot_id" json:"spot_id"`
	PromoVideos []string           `bson:"promo_videos" json:"promo_videos"`
	Address     entity.AddressDB   `bson:"address" json:"address"`
	CreatorID   string             `bson:"creator_id" json:"creator_id"`
	CreatedTime time.Time          `bson:"created_time" json:"created_time"`
	BrandID     []string           `bson:"brand_id" json:"brand_id"`
	Groups      []string           `bson:"groups" json:"groups"`
	Status      string             `bson:"status" json:"status"`
}
