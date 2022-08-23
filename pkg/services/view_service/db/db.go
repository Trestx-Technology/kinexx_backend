package viewDB

import (
	viewEntity "kinexx_backend/pkg/services/view_service/entity"
	viewRepository "kinexx_backend/pkg/services/view_service/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var repo = viewRepository.NewRepository("views")

type service struct{}

func (s service) Add(view *viewEntity.ViewDB) (viewEntity.ViewDB, error) {
	view.ID = primitive.NewObjectID()
	view.ViewID = view.ID.Hex()
	view.Status = 1
	view.CreatedTime = time.Now()
	view.UpdatedTime = time.Now()
	_, err := repo.InsertOne(view)
	if err != nil {
		return viewEntity.ViewDB{}, err
	}
	return repo.FindOne(bson.M{"_id": view.ID}, bson.M{})

}

func (s service) Update(view *viewEntity.ViewDB, campaignId string) (viewEntity.ViewDB, error) {
	id, err := primitive.ObjectIDFromHex(campaignId)
	if err != nil {
		return viewEntity.ViewDB{}, err
	}
	filter := bson.M{"_id": id}
	set := bson.M{}
	if view.Status != 0 {
		set["status"] = view.Status
	}
	if view.Banner != "" {
		set["banner"] = view.Banner
	}
	if view.Cover != "" {
		set["cover"] = view.Cover
	}
	if view.Trailer != "" {
		set["trailer"] = view.Trailer
	}
	if view.Tags != "" {
		set["tags"] = view.Tags
	}
	set["update_time"] = time.Now()
	_, err = repo.UpdateOne(filter, bson.M{"$set": set})
	if err != nil {
		return viewEntity.ViewDB{}, err
	}
	return repo.FindOne(filter, bson.M{})

}

func (s service) GetMy(userID string) ([]viewEntity.ViewDB, error) {
	return repo.Find(bson.M{"user_id": userID}, bson.M{})

}

func (s service) GetAll() ([]viewEntity.ViewDB, error) {
	return repo.Find(bson.M{}, bson.M{})

}

func (s service) Get(viewID string, userID string) ([]viewEntity.ViewDB, error) {
	_, err := repo.UpdateOne(bson.M{"view_id": viewID}, bson.M{"$inc": bson.M{"clicked": 1}, "$addToSet": bson.M{"clicked_by": userID}})
	if err != nil {
	}
	var pipeline = bson.A{
		bson.M{"$match": bson.M{"view_id": viewID}},
		bson.M{"$lookup": bson.M{
			"from":         "view_contents",
			"localField":   "view_id",
			"foreignField": "view_id",
			"as":           "content",
		},
		},
	}
	return repo.Aggregate(pipeline)
}

func (s service) Delete(viewID string) (string, error) {
	id, err := primitive.ObjectIDFromHex(viewID)
	if err != nil {
		return "unable To Delete", err
	}
	err = repo.DeleteOne(bson.M{"_id": id})
	if err != nil {
		return "unable To Delete", err
	}
	return "Deleted Successfully", nil
}

func NewService(repository viewRepository.Repository) Service {
	repo = repository
	return &service{}
}
