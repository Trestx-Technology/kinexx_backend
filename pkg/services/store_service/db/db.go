package db

import (
	"kinexx_backend/pkg/entity"
	"kinexx_backend/pkg/services/store_service/store"
	"kinexx_backend/pkg/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = stores.NewStoreRepository("stores")
)

type storeService struct{}

// AddShop implements StoreService
func (*storeService) AddShop(shop *entity.ShopDB) (string, error) {
	shop.ID = primitive.NewObjectID()
	shop.CreatedTime = time.Now()
	return repo.InsertOne(shop)
}

// DeleteShop implements StoreService
func (*storeService) DeleteShop(shopID string) (string, error) {
	id, err := primitive.ObjectIDFromHex(shopID)
	if err != nil {
		return "", err
	}
	err = repo.DeleteOne(bson.M{"_id": id})
	return "", err
}

// GetAllShops implements StoreService
func (*storeService) GetAllShops(spotID, groupID, brandID, creatorID string) ([]entity.ShopDB, error) {
	filter := bson.M{"spot_id": spotID}
	if groupID != "" {
		filter = bson.M{
			"groups": bson.M{"$in": bson.A{groupID}},
		}
	} else if brandID != "" {
		filter = bson.M{
			"brand_id": bson.M{"$in": bson.A{brandID}},
		}
	} else if creatorID != "" {
		filter = bson.M{
			"creator_id": creatorID,
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

// GetShop implements StoreService
func (*storeService) GetShop(shopID string) (entity.ShopDB, error) {
	id, err := primitive.ObjectIDFromHex(shopID)
	if err != nil {
		return entity.ShopDB{}, err
	}
	res, err := repo.FindOne(bson.M{"_id": id}, bson.M{})
	if err != nil {
		return entity.ShopDB{}, err
	}
	res.Banner = utils.CreatePreSignedDownloadUrl(res.Banner)
	res.Logo = utils.CreatePreSignedDownloadUrl(res.Logo)
	for j := range res.PromoVideos {
		res.PromoVideos[j] = utils.CreatePreSignedDownloadUrl(res.PromoVideos[j])
	}

	return res, nil
}

func NewStoreService(repository stores.StoreRepository) StoreService {
	repo = repository
	return &storeService{}
}
