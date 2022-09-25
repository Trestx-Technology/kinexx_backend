package db

import (
	"kinexx_backend/pkg/services/product_service/entity"
)

type ProductService interface {
	AddProduct(product *entity.ProductDB) (string, error)
	GetAllProducts(storeID string, spotID, proType string) ([]entity.ProductDB, error)
	GetProduct(productID string) (entity.ProductDB, error)
	DeleteProduct(productID string) (string, error)
	UpdateProduct(productID string, product *entity.ProductDB) (string, error)
	Count(proType string) (int64, error)
}
