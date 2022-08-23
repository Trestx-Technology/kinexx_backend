package db

import (
	"kinexx_backend/pkg/services/product_service/entity"
	"kinexx_backend/pkg/services/product_service/product"
	"kinexx_backend/pkg/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = products.NewProductRepository("products")
)

type productService struct{}

// UpdateProduct implements ProductService
func (*productService) UpdateProduct(productID string, product *entity.ProductDB) (string, error) {
	id, _ := primitive.ObjectIDFromHex(productID)
	filter := bson.M{"_id": id}
	data, err := repo.FindOne(filter, bson.M{})
	if err != nil {
		return "", err
	}
	set := bson.M{}

	if product.Name != "" && product.Name != data.Name {
		set["name"] = product.Name
	}
	if product.Description != "" && product.Description != data.Description {
		set["description"] = product.Description
	}
	if product.Banner != "" && product.Banner != data.Banner {
		set["banner"] = product.Banner
	}
	if product.Logo != "" && product.Logo != data.Logo {
		set["logo"] = product.Logo
	}
	if product.PromoVideos != nil && len(product.PromoVideos) > 0 && len(product.PromoVideos) != len(data.PromoVideos) {
		set["promo_videos"] = product.PromoVideos
	}
	if product.Images != nil && len(product.Images) > 0 && len(product.Images) != len(data.Images) {
		set["images"] = product.Images
	}
	if product.Quantity > 0 && product.Quantity != data.Quantity {
		set["quantity"] = product.Quantity
	}
	if product.QRCode != "" {
		set["qr_code"] = product.QRCode
	}
	if product.QRImage != "" {
		set["qr_image"] = product.QRImage
	}
	if product.Price != "" && product.Price != data.Price {
		set["price"] = product.Price
	}
	if product.DiscountedPrice != "" && product.DiscountedPrice != data.DiscountedPrice {
		set["discounted_price"] = product.DiscountedPrice
	}
	if product.Deals != "" && product.Deals != data.Deals {
		set["deals"] = product.Deals
	}
	if product.Tags != nil && len(product.Tags) > 0 && len(product.Tags) != len(data.Tags) {
		set["tags"] = product.Tags
	}
	if product.AgeGroup != "" && product.AgeGroup != data.AgeGroup {
		set["age_group"] = product.AgeGroup
	}
	if product.Status != "" && product.Status != data.Status {
		set["status"] = product.Status
	}
	if product.SpotID != "" {
		set["spot_id"] = product.SpotID
	}
	if product.StoreID != "" {
		set["store_id"] = product.StoreID
	}
	if product.BrandID != "" {
		set["brand_id"] = product.BrandID
	}
	if product.Address.Address != "" && product.Address.Address != data.Address.Address {
		set["address.address"] = product.Address.Address
	}
	if product.Address.GeoLocation != nil {
		set["address.geo_location"] = product.Address.GeoLocation
	}
	if product.Address.City != "" && product.Address.City != data.Address.City {
		set["address.city"] = product.Address.City
	}
	if product.Address.State != "" && product.Address.State != data.Address.State {
		set["address.state"] = product.Address.State
	}
	if product.Address.Country != "" && product.Address.Country != data.Address.Country {
		set["address.country"] = product.Address.Country
	}
	return repo.UpdateOne(filter, bson.M{"$set": set})
}

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

func NewProductService(repository products.ProductRepository) ProductService {
	repo = repository
	return &productService{}
}
