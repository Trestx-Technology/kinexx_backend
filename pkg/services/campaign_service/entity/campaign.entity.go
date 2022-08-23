package campaignEntity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Campaign struct {
	ID                   primitive.ObjectID `bson:"_id" json:"id"`
	CampaignID           string             `bson:"campaign_id" json:"campaign_id"`
	UserID               string             `bson:"user_id" json:"user_id"`
	Name                 string             `bson:"name" json:"name"`
	CampaignManagerEmail string             `bson:"campaign_manager_email" json:"campaign_manager_email"`
	CampaignManagerName  string             `bson:"campaign_manager_name" json:"campaign_manager_name"`
	CampaignStartDate    string             `bson:"campaign_start_date" json:"campaign_start_date"`
	CampaignEndDate      string             `bson:"campaign_end_date" json:"campaign_end_date"`
	ContentURL           string             `bson:"content_url" json:"content_url"`
	ContentTitle         string             `bson:"content_title" json:"content_title"`
	PlacementStartDate   string             `bson:"placement_start_date" json:"placement_start_date"`
	PlacementEndDate     string             `bson:"placement_end_date" json:"placement_end_date"`
	ProductName          string             `bson:"product_name" json:"product_name"`
	ManufacturerName     string             `bson:"manufacturer_name" json:"manufacturer_name"`
	DistributorName      string             `bson:"distributor_name" json:"distributor_name"`
	ProductType          string             `bson:"product_type" json:"product_type"`
	ProductSize          string             `bson:"product_size" json:"product_size"`
	ProductQuantity      int64              `bson:"product_quantity" json:"product_quantity"`
	DiscountRate         float64            `bson:"discount_rate" json:"discount_rate"`
	DiscountAvailability string             `bson:"discount_availability" json:"discount_availability"`
	ProductImageURL      string             `bson:"product_image_url" json:"product_image_url"`
	ProductDescription   string             `bson:"product_description" json:"product_description"`
	ProductPrice         float64            `bson:"product_price" json:"product_price"`
	CharityID            string             `bson:"charity_id" json:"charity_id"`
	DonationAmount       float64            `bson:"donation_amount" json:"donation_amount"`
	VideoURL             string             `bson:"video_url" json:"video_url"`
	VideoTitle           string             `bson:"video_title" json:"video_title"`
	VideoDescription     string             `bson:"video_description" json:"video_description"`
	VideoRelatedTags     string             `bson:"video_related_tags" json:"video_related_tags"`
	PlaylistName         string             `bson:"playlist_name" json:"playlist_name"`
	CreatedDate          time.Time          `bson:"created_date" json:"created_date"`
	ModifiedDate         time.Time          `bson:"modified_date" json:"modified_date" `
	Status               int                `bson:"status" json:"status"`
}
