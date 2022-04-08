package goalDB

import (
	"kinexx_backend/pkg/entity"
	"kinexx_backend/pkg/repository/goals"
	"kinexx_backend/pkg/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = goals.NewGoalRepository("goals")
)

type goalService struct{}

func (*goalService) MakeGoal(goal *entity.GoalDB) (string, error) {

	goal.ID = primitive.NewObjectID()
	goal.CreatedDate = time.Now()

	_, err := repo.InsertOne(goal)
	if err != nil {
		return "", err
	}
	if len(goal.GroupIDList) > 0 {
		utils.AddGoalToGroup(goal.ID.Hex(), goal.GroupIDList)
	}
	return "Created Successfully", nil
}
func (*goalService) EditGoal(goal entity.GoalDB, goalID string) (string, error) {
	id, err := primitive.ObjectIDFromHex(goalID)
	if err != nil {
		return "", err
	}
	filter := bson.M{"_id": id}
	set := bson.M{}
	if goal.Name != "" {
		set["name"] = goal.Name
	}
	if goal.Description != "" {
		set["description"] = goal.Description
	}

	if goal.Banner != "" {
		set["banner"] = goal.Banner
	}

	return repo.UpdateOne(filter, bson.M{"$set": set})
}

// GetGoal implements GoalService
func (*goalService) GetGoal(goalID string) (entity.GoalDB, error) {
	id, err := primitive.ObjectIDFromHex(goalID)
	if err != nil {
		return entity.GoalDB{}, nil
	}
	filter := bson.M{"_id": id}
	goals, err := repo.FindOne(filter, bson.M{})
	if err != nil {
		return entity.GoalDB{}, err
	}
	goals.Banner = utils.CreatePreSignedDownloadUrl(goals.Banner)

	return goals, err
}

// GetGoalCreatedByMe implements GoalService
func (*goalService) GetGoalCreatedByMe(userID string) ([]entity.GoalDB, error) {
	goals, err := repo.Find(bson.M{"creator_user_id": userID}, bson.M{})
	if err != nil {
		return []entity.GoalDB{}, err
	}
	for i := range goals {
		goals[i].Banner = utils.CreatePreSignedDownloadUrl(goals[i].Banner)

	}
	return goals, err
}

func GetGoalsByIDs(goalIDs []string) ([]entity.GoalDB, error) {
	goalBsonArray := bson.A{}
	for _, goalID := range goalIDs {
		id, err := primitive.ObjectIDFromHex(goalID)
		if err == nil {
			goalBsonArray = append(goalBsonArray, bson.M{"_id": id})
		}
	}
	goals, err := repo.Find(bson.M{"$or": goalBsonArray}, bson.M{})
	if err != nil {
		return []entity.GoalDB{}, err
	}
	for i := range goals {
		goals[i].Banner = utils.CreatePreSignedDownloadUrl(goals[i].Banner)

	}
	return goals, err
}
func NewGoalService(repository goals.GoalRepository) GoalService {
	repo = repository
	return &goalService{}
}

func (*goalService) DeleteGoal(goalID string) error {
	id, err := primitive.ObjectIDFromHex(goalID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	return repo.DeleteOne(filter)
}

func (*goalService) PauseGoal(status, goalID string) (string, error) {
	id, err := primitive.ObjectIDFromHex(goalID)
	if err != nil {
		return "", err
	}
	filter := bson.M{"_id": id}
	set := bson.M{"$set": bson.M{"status": status}}

	return repo.UpdateOne(filter, set)
}
