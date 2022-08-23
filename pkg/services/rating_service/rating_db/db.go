package rating_db

import (
	entity2 "kinexx_backend/pkg/entity"
	postDB "kinexx_backend/pkg/services/post_service/db"
	entity3 "kinexx_backend/pkg/services/post_service/entity"
	"kinexx_backend/pkg/services/profile_service/db"
	"kinexx_backend/pkg/services/rating_service/entity"
	"kinexx_backend/pkg/services/rating_service/ratings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = ratings.NewRatingRepository("ratings")
)

type ratingService struct{}

func NewRatingService(repository ratings.RatingRepository) RatingService {
	repo = repository
	return &ratingService{}
}

// AddRating implements RatingService
func (*ratingService) AddRating(rating *entity.RatingDB) (string, error) {
	rating.CreatedDate = time.Now()
	rating.ID = primitive.NewObjectID()
	return repo.InsertOne(rating)
}

// GetAllRatingsByUserID implements RatingService
func (*ratingService) GetAllRatingsByUserID(userID string) ([]entity.RatingDB, error) {
	filter := bson.M{"user_id": userID}
	res, err := repo.Find(filter, bson.M{})
	if err != nil {
		return []entity.RatingDB{}, err
	}
	var userIDs []string
	for i := range res {
		if res[i].ItemType == "PERSON" {
			userIDs = append(userIDs, res[i].ItemID)
		}
	}
	profiles, _ := db.GetProfilesForIDs(userIDs)
	for i := range res {
		for _, profile := range profiles {
			if profile.ID.Hex() == res[i].ItemID {
				res[i].ItemUser = profile
			}
		}
	}
	return res, nil
}

// UpdateRating implements RatingService
func (*ratingService) UpdateRating(*entity.RatingDB) (string, error) {
	panic("unimplemented")
}
func (*ratingService) GetUserReviews(itemID string) ([]entity.RatingDB, float64, entity2.ProfileDB, []entity3.PostDB, error) {
	profile, posts, err := postDB.GetUserDataInternal(itemID)
	if err != nil {
		return []entity.RatingDB{}, 0, profile, posts, err
	}
	filter := bson.M{"item_id": itemID}
	res, err := repo.Find(filter, bson.M{})
	if err != nil {
		return []entity.RatingDB{}, 0, profile, posts, err
	}
	var userIDs []string
	averageRating := float64(0)
	if len(res) > 0 {
		for i := range res {
			userIDs = append(userIDs, res[i].UserID)
			averageRating += float64(res[i].Rating)
		}
		averageRating /= float64(len(res))
	}
	profiles, _ := db.GetProfilesForIDs(userIDs)
	for i := range res {
		for _, profile := range profiles {
			if profile.ID.Hex() == res[i].UserID {
				res[i].ItemUser = profile
			}
		}
	}
	return res, averageRating, profile, posts, nil
}

func (*ratingService) GetItemReviews(itemID string) ([]entity.RatingDB, float64, error) {
	filter := bson.M{"item_id": itemID}
	res, err := repo.Find(filter, bson.M{})
	if err != nil {
		return []entity.RatingDB{}, 0, err
	}
	var userIDs []string
	averageRating := float64(0)
	for i := range res {
		userIDs = append(userIDs, res[i].UserID)
		averageRating += float64(res[i].Rating)
	}
	averageRating /= float64(len(res))
	profiles, _ := db.GetProfilesForIDs(userIDs)
	for i := range res {
		for _, profile := range profiles {
			if profile.ID.Hex() == res[i].UserID {
				res[i].ItemUser = profile
			}
		}
	}
	return res, averageRating, nil
}
