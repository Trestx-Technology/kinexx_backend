package db

import (
	"kinexx_backend/pkg/services/spot_service/entity"
	"kinexx_backend/pkg/services/spot_service/spot"
	"kinexx_backend/pkg/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = spots.NewSpotRepository("spots")
)

type spotService struct{}

// AddSpot implements SpotService
func (*spotService) AddSpot(spot *entity.SpotDB) (string, error) {
	spot.ID = primitive.NewObjectID()
	spot.CreatedTime = time.Now()
	return repo.InsertOne(spot)
}

// DeleteSpot implements SpotService
func (*spotService) DeleteSpot(spotID string) (string, error) {
	id, err := primitive.ObjectIDFromHex(spotID)
	if err != nil {
		return "", err
	}
	err = repo.DeleteOne(bson.M{"_id": id})
	return "", err
}

// GetAllSpots implements SpotService
func (*spotService) GetAllSpots(groupID, creator_id string) ([]entity.SpotDB, error) {
	filter := bson.M{}
	if groupID != "" {
		filter = bson.M{
			"groups": bson.M{"$in": bson.A{groupID}},
		}
	} else if creator_id != "" {
		filter = bson.M{
			"creator_id": creator_id,
		}
	}
	res, err := repo.Find(filter, bson.M{})
	if err != nil {
		return nil, err
	}
	for i := range res {
		res[i].Banner = utils.CreatePreSignedDownloadUrl(res[i].Banner)
		res[i].Logo = utils.CreatePreSignedDownloadUrl(res[i].Logo)
		for j := range res[i].PromoVideos {
			res[i].PromoVideos[j] = utils.CreatePreSignedDownloadUrl(res[i].PromoVideos[j])
		}
	}
	return res, nil
}

// GetSpot implements SpotService
func (*spotService) GetSpot(spotID string) (entity.SpotDB, error) {
	id, err := primitive.ObjectIDFromHex(spotID)
	if err != nil {
		return entity.SpotDB{}, err
	}
	res, err := repo.FindOne(bson.M{"_id": id}, bson.M{})
	if err != nil {
		return entity.SpotDB{}, err
	}
	res.Banner = utils.CreatePreSignedDownloadUrl(res.Banner)
	res.Logo = utils.CreatePreSignedDownloadUrl(res.Logo)
	for j := range res.PromoVideos {
		res.PromoVideos[j] = utils.CreatePreSignedDownloadUrl(res.PromoVideos[j])
	}

	return res, nil
}

func NewSpotService(repository spots.SpotRepository) SpotService {
	repo = repository
	return &spotService{}
}
