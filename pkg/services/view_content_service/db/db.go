package viewContentDB

import (
	viewContentEntity "kinexx_backend/pkg/services/view_content_service/entity"
	viewContentRepository "kinexx_backend/pkg/services/view_content_service/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = viewContentRepository.NewRepository("view_contents")
)

type service struct{}

func (s service) Add(viewContent *viewContentEntity.ViewContentDB) (viewContentEntity.ViewContentDB, error) {
	viewContent.ID = primitive.NewObjectID()
	viewContent.ViewContentID = viewContent.ID.Hex()
	viewContent.CreatedTime = time.Now()
	viewContent.UpdatedTime = time.Now()
	viewContent.Status = 1
	_, err := repo.InsertOne(viewContent)
	if err != nil {
		return viewContentEntity.ViewContentDB{}, err
	}
	return repo.FindOne(bson.M{"_id": viewContent.ID}, bson.M{})
}

func (s service) Update(viewContent *viewContentEntity.ViewContentDB, campaignId string) (viewContentEntity.ViewContentDB, error) {
	id, err := primitive.ObjectIDFromHex(campaignId)
	if err != nil {
		return viewContentEntity.ViewContentDB{}, err
	}
	filter := bson.M{"_id": id}
	set := bson.M{}
	if viewContent.Status != 0 {
		set["status"] = viewContent.Status
	}
	if viewContent.Banner != "" {
		set["banner"] = viewContent.Banner
	}
	if viewContent.Cover != "" {
		set["cover"] = viewContent.Cover
	}
	if viewContent.VideoURL != "" {
		set["video_url"] = viewContent.VideoURL
	}
	if viewContent.Tags != "" {
		set["tags"] = viewContent.Tags
	}
	set["update_time"] = time.Now()
	_, err = repo.UpdateOne(filter, bson.M{"$set": set})
	if err != nil {
		return viewContentEntity.ViewContentDB{}, err
	}
	return repo.FindOne(filter, bson.M{})
}

func (s service) GetForCampaign(userID string) ([]viewContentEntity.ViewContentDB, error) {
	return repo.Find(bson.M{"user_id": userID}, bson.M{})
}

func (s service) Delete(viewContentID string) (string, error) {
	id, err := primitive.ObjectIDFromHex(viewContentID)
	if err != nil {
		return "unable To Delete", err
	}
	err = repo.DeleteOne(bson.M{"_id": id})
	if err != nil {
		return "unable To Delete", err
	}
	return "Deleted Successfully", nil
}

func NewService(repository viewContentRepository.Repository) Service {
	repo = repository
	return &service{}
}
