package db

import "kinexx_backend/pkg/entity"

type ShareService interface {
	AddShare(post *entity.ShareDB) (string, error)
	UpdateShare(post *entity.ShareDB) (string, error)
	GetShareByID(userID, shareType string) ([]entity.ShareDB, error)
	GetMyShare(user, shareType string) ([]entity.ShareDB, error)
}
