package db

import "kinexx_backend/pkg/entity"

type ProductService interface {
	AddProduct(product *entity.ProductDB) (string, error)
	GetAllProducts(storeID string, spotID, proType string) ([]entity.ProductDB, error)
	GetProduct(productID string) (entity.ProductDB, error)
	DeleteProduct(productID string) (string, error)
}
