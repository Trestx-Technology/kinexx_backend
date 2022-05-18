package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductDB struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`
	Logo            string             `bson:"logo" json:"logo"`
	Banner          string             `bson:"banner" json:"banner"`
	Name            string             `bson:"name" json:"name"`
	SpotID          string             `bson:"spot_id" json:"spot_id"`
	StoreID         string             `bson:"store_id" json:"store_id"`
	BrandID         string             `bson:"brand_id" json:"brand_id"`
	PromoVideos     []string           `bson:"promo_videos" json:"promo_videos"`
	Images          []string           `bson:"images" json:"images"`
	Quantity        int                `bson:"quantity" json:"quantity"`
	QRCode          string             `bson:"qr_code" json:"qr_code"`
	QRImage         string             `bson:"qr_image" json:"qr_image"`
	Type            string             `bson:"type" json:"type"`
	Price           string             `bson:"price" json:"price"`
	DiscountedPrice string             `bson:"discounted_price" json:"discounted_price"`
	Deals           string             `bson:"deals" json:"deals"`
	Description     string             `bson:"description" json:"description"`
	Tags            []string           `bson:"tags" json:"tags"`
	AgeGroup        string             `bson:"age_group" json:"age_group"`
	CreatorID       string             `bson:"creator_id" json:"creator_id"`
	CreatedTime     time.Time          `bson:"created_time" json:"created_time"`
	Status          string             `bson:"status" json:"status"`
	Address         AddressDB          `bson:"address" json:"address"`
}
