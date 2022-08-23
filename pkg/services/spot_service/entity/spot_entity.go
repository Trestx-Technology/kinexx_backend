package entity

import (
	"kinexx_backend/pkg/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SpotDB struct {
	ID                primitive.ObjectID `bson:"_id" json:"id"`
	Logo              string             `bson:"logo" json:"logo"`
	Banner            string             `bson:"banner" json:"banner"`
	Name              string             `bson:"name" json:"name"`
	BussinessName     string             `bson:"bussiness_name" json:"bussiness_name"`
	BussinessID       string             `bson:"bussiness_id" json:"bussiness_id"`
	CharityID         string             `bson:"charity_id" json:"charity_id"`
	PromoVideos       []string           `bson:"promo_videos" json:"promo_videos"`
	NumberOfLocations int                `bson:"number_of_locations" json:"number_of_locations"`
	Addresses         []entity.AddressDB `bson:"addresses" json:"addresses"`
	Package           string             `bson:"package" json:"package"`
	Payment           string             `bson:"payment" json:"payment"`
	CreatorID         string             `bson:"creator_id" json:"creator_id"`
	CreatedTime       time.Time          `bson:"created_time" json:"created_time"`
	BrandID           []string           `bson:"brand_id" json:"brand_id"`
	Groups            []string           `bson:"groups" json:"groups"`
	Status            string             `bson:"status" json:"status"`
}
