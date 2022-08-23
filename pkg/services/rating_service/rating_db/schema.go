package rating_db

import (
	entity2 "kinexx_backend/pkg/entity"
	entity3 "kinexx_backend/pkg/services/post_service/entity"
	"kinexx_backend/pkg/services/rating_service/entity"
)

type RatingService interface {
	AddRating(rating *entity.RatingDB) (string, error)
	UpdateRating(rating *entity.RatingDB) (string, error)
	GetAllRatingsByUserID(userID string) ([]entity.RatingDB, error)
	GetItemReviews(itemID string) ([]entity.RatingDB, float64, error)
	GetUserReviews(itemID string) ([]entity.RatingDB, float64, entity2.ProfileDB, []entity3.PostDB, error)
}
