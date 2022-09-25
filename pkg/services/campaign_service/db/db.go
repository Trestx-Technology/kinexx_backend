package campaignDB

import (
	campaignEntity "kinexx_backend/pkg/services/campaign_service/entity"
	"kinexx_backend/pkg/services/campaign_service/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = repository.NewCampaignRepository("campaigns")
)

type campaignService struct{}

func (c campaignService) AddCampaign(brand *campaignEntity.Campaign) (campaignEntity.Campaign, error) {
	brand.ID = primitive.NewObjectID()
	brand.CampaignID = brand.ID.Hex()
	brand.CreatedDate = time.Now()
	brand.ModifiedDate = time.Now()
	brand.Status = 1
	_, err := repo.InsertOne(brand)
	if err != nil {
		return campaignEntity.Campaign{}, err
	}
	return repo.FindOne(bson.M{"_id": brand.ID}, bson.M{})
}

func (c campaignService) UpdateCampaign(brand *campaignEntity.Campaign, campaignId string) (campaignEntity.Campaign, error) {
	id, err := primitive.ObjectIDFromHex(campaignId)
	if err != nil {
		return campaignEntity.Campaign{}, err
	}
	set := bson.M{}
	if brand.Name != "" {
		set["name"] = brand.Name
	}
	if brand.CampaignManagerEmail != "" {
		set["campaign_manager_email"] = brand.CampaignManagerEmail
	}
	if brand.CampaignManagerName != "" {
		set["campaign_manager_name"] = brand.CampaignManagerName
	}
	if brand.CampaignStartDate != "" {
		set["campaign_start_date"] = brand.CampaignStartDate
	}
	if brand.CampaignEndDate != "" {
		set["campaign_end_date"] = brand.CampaignEndDate
	}
	if brand.ContentURL != "" {
		set["content_url"] = brand.ContentURL
	}
	if brand.ContentTitle != "" {
		set["content_title"] = brand.ContentTitle
	}
	if brand.PlacementStartDate != "" {
		set["placement_start_date"] = brand.PlacementStartDate
	}
	if brand.PlacementEndDate != "" {
		set["placement_end_date"] = brand.PlacementEndDate
	}
	if brand.ProductName != "" {
		set["product_name"] = brand.ProductName
	}
	if brand.ManufacturerName != "" {
		set["manufacturer_name"] = brand.ManufacturerName
	}
	if brand.DistributorName != "" {
		set["distributor_name"] = brand.DistributorName
	}
	if brand.ProductType != "" {
		set["product_type"] = brand.ProductType
	}
	if brand.ProductSize != "" {
		set["product_size"] = brand.ProductSize
	}
	if brand.ProductQuantity != 0 {
		set["product_quantity"] = brand.ProductQuantity
	}
	if brand.DiscountRate != 0 {
		set["discount_rate"] = brand.DiscountRate
	}
	if brand.DiscountAvailability != "" {
		set["discount_availability"] = brand.DiscountAvailability
	}
	if brand.ProductImageURL != "" {
		set["product_image_url"] = brand.ProductImageURL
	}
	if brand.ProductDescription != "" {
		set["product_description"] = brand.ProductDescription
	}
	if brand.ProductPrice != 0 {
		set["product_price"] = brand.ProductPrice
	}
	if brand.CharityID != "" {
		set["charity_id"] = brand.CharityID
	}
	if brand.DonationAmount != 0 {
		set["donation_amount"] = brand.DonationAmount
	}
	if brand.VideoURL != "" {
		set["video_url"] = brand.VideoURL
	}
	if brand.VideoTitle != "" {
		set["video_title"] = brand.VideoTitle
	}
	if brand.VideoDescription != "" {
		set["video_description"] = brand.VideoDescription
	}
	if brand.VideoRelatedTags != "" {
		set["video_related_tags"] = brand.VideoRelatedTags
	}
	if brand.PlaylistName != "" {
		set["playlist_name"] = brand.PlaylistName
	}
	if brand.Status != 0 {
		set["status"] = brand.Status
	}
	set["modified_data"] = time.Now()
	_, err = repo.UpdateOne(bson.M{"_id": id}, bson.M{"$set": set})
	if err != nil {
		return campaignEntity.Campaign{}, err
	}
	return repo.FindOne(bson.M{"_id": id}, bson.M{})
}

func (c campaignService) GetMyCampaign(userID string) ([]campaignEntity.Campaign, error) {
	return repo.Find(bson.M{"user_id": userID}, bson.M{})

}

func (c campaignService) GetAllCampaign() ([]campaignEntity.Campaign, error) {
	return repo.Find(bson.M{}, bson.M{})

}
func (c campaignService) Count() (int64, error) {
	return repo.Count(bson.M{})

}
func (c campaignService) SearchCampaign(search string) ([]campaignEntity.Campaign, error) {
	return repo.Find(bson.M{"name": search}, bson.M{})

}

func (c campaignService) GetCampaign(brandID string) (campaignEntity.Campaign, error) {
	id, err := primitive.ObjectIDFromHex(brandID)
	if err != nil {
		return campaignEntity.Campaign{}, err
	}
	return repo.FindOne(bson.M{"_id": id}, bson.M{})
}

func (c campaignService) DeleteCampaign(brandID string) (string, error) {
	id, err := primitive.ObjectIDFromHex(brandID)
	if err != nil {
		return "unable To Delete", err
	}
	err = repo.DeleteOne(bson.M{"_id": id})
	if err != nil {
		return "unable To Delete", err
	}
	return "Deleted Successfully", nil
}

func NewCampaignService(repository repository.CampaignRepository) CampaignService {
	repo = repository
	return &campaignService{}
}
