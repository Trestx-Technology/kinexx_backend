package rating_db

import "kinexx_backend/pkg/entity"

type RatingService interface {
	AddRating(rating *entity.RatingDB) (string, error)
	UpdateRating(rating *entity.RatingDB) (string, error)
	GetAllRatingsByUserID(userID string) ([]entity.RatingDB, error)
	GetItemReviews(itemID string) ([]entity.RatingDB, float64, error)
	GetUserReviews(itemID string) ([]entity.RatingDB, float64, entity.ProfileDB, []entity.PostDB, error)
}
