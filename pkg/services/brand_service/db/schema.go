package db

import "kinexx_backend/pkg/entity"

type BrandService interface {
	AddBrand(brand *entity.BrandDB) (string, error)
	GetAllBrands() ([]entity.BrandDB, error)
	GetBrand(brandID string) (entity.BrandDB, error)
	DeleteBrand(brandID string) (string, error)
}
