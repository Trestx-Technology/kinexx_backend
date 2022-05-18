package db

import (
	"kinexx_backend/pkg/entity"
	product "kinexx_backend/pkg/repository/product"
	"kinexx_backend/pkg/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = product.NewProductRepository("products")
)

type productService struct{}

// AddProduct implements ProductService
func (*productService) AddProduct(product *entity.ProductDB) (string, error) {
	product.ID = primitive.NewObjectID()
	product.CreatedTime = time.Now()
	return repo.InsertOne(product)
}

// DeleteProduct implements ProductService
func (*productService) DeleteProduct(productID string) (string, error) {
	id, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return "", err
	}
	err = repo.DeleteOne(bson.M{"_id": id})
	return "", err
}

// GetAllProducts implements ProductService
func (*productService) GetAllProducts(storeID, spotID, proType string) ([]entity.ProductDB, error) {
	filter := bson.M{"spot_id": spotID}
	if storeID != "" {
		filter = bson.M{"store_id": storeID}
	}
	if proType != "" {
		filter["type"] = proType
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
		for j := range res[i].Images {
			res[i].Images[j] = utils.CreatePreSignedDownloadUrl(res[i].Images[j])
		}
	}
	return res, nil
}

// GetProduct implements ProductService
func (*productService) GetProduct(productID string) (entity.ProductDB, error) {
	id, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return entity.ProductDB{}, err
	}
	res, err := repo.FindOne(bson.M{"_id": id}, bson.M{})
	if err != nil {
		return entity.ProductDB{}, err
	}
	res.Banner = utils.CreatePreSignedDownloadUrl(res.Banner)
	res.Logo = utils.CreatePreSignedDownloadUrl(res.Logo)
	for j := range res.PromoVideos {
		res.PromoVideos[j] = utils.CreatePreSignedDownloadUrl(res.PromoVideos[j])
	}
	for j := range res.Images {
		res.Images[j] = utils.CreatePreSignedDownloadUrl(res.Images[j])
	}

	return res, nil
}

func NewProductService(repository product.ProductRepository) ProductService {
	repo = repository
	return &productService{}
}
