package hobby_db

import (
	"kinexx_backend/pkg/entity"
	"kinexx_backend/pkg/services/hobby_service/hobbies"
	"kinexx_backend/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = hobbies.NewHobbiesRepository("hobbies")
)

func FindHobbies(query string) ([]entity.HobbiesDB, error) {
	hobies, err := repo.Find(bson.M{"name": bson.M{"$regex": query, "$options": "i"}}, bson.M{})
	if err == nil {
		for i := range hobies {
			hobies[i].Image = utils.CreatePreSignedDownloadUrl(hobies[i].Image)
		}
	}
	return hobies, err
}
func GetHobbies(ids []string) ([]entity.HobbiesDB, error) {
	if len(ids) > 0 {
		filter := bson.A{}
		for _, i := range ids {
			id, _ := primitive.ObjectIDFromHex(i)
			filter = append(filter, bson.M{"_id": id})
		}
		hobies, err := repo.Find(bson.M{"$or": filter}, bson.M{})
		if err == nil {
			for i := range hobies {
				hobies[i].Image = utils.CreatePreSignedDownloadUrl(hobies[i].Image)
			}
		}
		return hobies, err
	}
	return []entity.HobbiesDB{}, nil
}

func AddHobby(hobby *entity.HobbiesDB) (string, error) {
	hobby.ID = primitive.NewObjectID()
	return repo.InsertOne(hobby)
}
