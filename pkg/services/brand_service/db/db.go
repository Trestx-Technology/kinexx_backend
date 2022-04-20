package db

import (
	"kinexx_backend/pkg/entity"
	brand "kinexx_backend/pkg/repository/brand"
	"kinexx_backend/pkg/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = brand.NewBrandRepository("brands")
)

type brandService struct{}

// AddBrand implements BrandService
func (*brandService) AddBrand(brand *entity.BrandDB) (string, error) {
	brand.ID = primitive.NewObjectID()
	brand.CreatedTime = time.Now()
	return repo.InsertOne(brand)

}

// DeleteBrand implements BrandService
func (*brandService) DeleteBrand(brandID string) (string, error) {
	id, err := primitive.ObjectIDFromHex(brandID)
	if err != nil {
		return "", err
	}
	err = repo.DeleteOne(bson.M{"_id": id})
	return "", err
}

// GetAllBrands implements BrandService
func (*brandService) GetAllBrands() ([]entity.BrandDB, error) {
	res, err := repo.Find(bson.M{}, bson.M{})
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

// GetBrand implements BrandService
func (*brandService) GetBrand(brandID string) (entity.BrandDB, error) {
	id, err := primitive.ObjectIDFromHex(brandID)
	if err != nil {
		return entity.BrandDB{}, err
	}
	res, err := repo.FindOne(bson.M{"_id": id}, bson.M{})
	if err != nil {
		return entity.BrandDB{}, err
	}
	res.Banner = utils.CreatePreSignedDownloadUrl(res.Banner)
	res.Logo = utils.CreatePreSignedDownloadUrl(res.Logo)
	for j := range res.PromoVideos {
		res.PromoVideos[j] = utils.CreatePreSignedDownloadUrl(res.PromoVideos[j])
	}

	return res, nil
}

func NewBrandService(repository brand.BrandRepository) BrandService {
	repo = repository
	return &brandService{}
}
