package groupDB

import (
	"kinexx_backend/pkg/entity"
	"kinexx_backend/pkg/repository/groups"
	groupUserInternal "kinexx_backend/pkg/services/group_user_service/db_internal"
	"kinexx_backend/pkg/services/profile_service/db"
	"kinexx_backend/pkg/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = groups.NewGroupRepository("groups")
)

type groupService struct{}

func (*groupService) MakeGroup(group *entity.GroupDB) (string, error) {

	group.ID = primitive.NewObjectID()
	group.CreatedDate = time.Now()

	_, err := repo.InsertOne(group)
	if err != nil {
		return "", err
	}
	if len(group.GoalIDList) > 0 {
		utils.AddGroupToGoal(group.GoalIDList, group.ID.Hex())
	}
	if len(group.UserIDList) > 0 {
		for _, id := range group.UserIDList {
			groupUserInternal.AddUserToGroupInternal(group.ID.Hex(), id)
		}
	}
	return "Created Successfully", nil
}
func (*groupService) EditGroup(group entity.GroupDB, groupID string) (string, error) {
	id, err := primitive.ObjectIDFromHex(groupID)
	if err != nil {
		return "", err
	}
	filter := bson.M{"_id": id}
	set := bson.M{}
	if group.Name != "" {
		set["name"] = group.Name
	}
	if group.Description != "" {
		set["description"] = group.Description
	}
	if group.Visible != "" {
		set["visible"] = group.Visible
	}
	if group.Banner != "" {
		set["banner"] = group.Banner
	}
	if group.Logo != "" {
		set["logo"] = group.Logo
	}
	if len(group.PromoVideos) > 0 {
		set["promo_videos"] = group.PromoVideos
	}
	if group.Status != "" {
		set["status"] = group.Status
	}
	return repo.UpdateOne(filter, bson.M{"$set": set})
}

// GetGroup implements GroupService
func (*groupService) GetGroup(groupID string) (entity.GroupDB, error) {
	id, err := primitive.ObjectIDFromHex(groupID)
	if err != nil {
		return entity.GroupDB{}, nil
	}
	filter := bson.M{"_id": id}
	groups, err := repo.FindOne(filter, bson.M{})
	if err != nil {
		return entity.GroupDB{}, err
	}
	groups.Banner = utils.CreatePreSignedDownloadUrl(groups.Banner)
	groups.Logo = utils.CreatePreSignedDownloadUrl(groups.Logo)
	for j := range groups.PromoVideos {
		groups.PromoVideos[j] = utils.CreatePreSignedDownloadUrl(groups.PromoVideos[j])
	}
	userDetails, _ := db.GetProfilesForIDs([]string{groups.CreatorUserID})
	if len(userDetails) > 0 {
		groups.CreatorDetails = userDetails[0]
	}
	return groups, err
}

// GetGroupCreatedByMe implements GroupService
func (*groupService) GetGroupCreatedByMe(userID string) ([]entity.GroupDB, error) {
	groups, err := repo.Find(bson.M{"creator_user_id": userID}, bson.M{})
	if err != nil {
		return []entity.GroupDB{}, err
	}
	for i := range groups {
		groups[i].Banner = utils.CreatePreSignedDownloadUrl(groups[i].Banner)
		groups[i].Logo = utils.CreatePreSignedDownloadUrl(groups[i].Logo)
		for j := range groups[i].PromoVideos {
			groups[i].PromoVideos[j] = utils.CreatePreSignedDownloadUrl(groups[i].PromoVideos[j])
		}
	}
	return groups, err
}

func GetGroupsByIDs(groupIDs []string) ([]entity.GroupDB, error) {
	groupBsonArray := bson.A{}
	for _, groupID := range groupIDs {
		id, err := primitive.ObjectIDFromHex(groupID)
		if err == nil {
			groupBsonArray = append(groupBsonArray, bson.M{"_id": id})
		}
	}
	groups, err := repo.Find(bson.M{"$or": groupBsonArray}, bson.M{})
	if err != nil {
		return []entity.GroupDB{}, err
	}
	for i := range groups {
		groups[i].Banner = utils.CreatePreSignedDownloadUrl(groups[i].Banner)
		groups[i].Logo = utils.CreatePreSignedDownloadUrl(groups[i].Logo)
		for j := range groups[i].PromoVideos {
			groups[i].PromoVideos[j] = utils.CreatePreSignedDownloadUrl(groups[i].PromoVideos[j])
		}
	}
	return groups, err
}
func NewGroupService(repository groups.GroupRepository) GroupService {
	repo = repository
	return &groupService{}
}

func (*groupService) DeleteGroup(groupID string) error {
	id, err := primitive.ObjectIDFromHex(groupID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	return repo.DeleteOne(filter)
}

func (*groupService) PauseGroup(status, groupID string) (string, error) {
	id, err := primitive.ObjectIDFromHex(groupID)
	if err != nil {
		return "", err
	}
	filter := bson.M{"_id": id}
	set := bson.M{"$set": bson.M{"status": status}}

	return repo.UpdateOne(filter, set)
}
