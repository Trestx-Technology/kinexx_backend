package db

import (
	"kinexx_backend/pkg/services/store_service/entity"
)

type StoreService interface {
	AddShop(shop *entity.ShopDB) (string, error)
	GetAllShops(spotID, groupID, brandID, creatorID string) ([]entity.ShopDB, error)
	GetShop(shopID string) (entity.ShopDB, error)
	DeleteShop(shopID string) (string, error)
}
