package db

import (
	"kinexx_backend/pkg/entity"
	"kinexx_backend/pkg/repository/comments"
	"kinexx_backend/pkg/services/profile_service/db"
	"kinexx_backend/pkg/utils"

	"errors"
	"time"

	"github.com/aekam27/trestCommon"

	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = comments.NewCommentRepository("comments")
)

type commentService struct{}

func NewCommentService(repository comments.CommentRepository) CommentService {
	repo = repository
	return &commentService{}
}

func (*commentService) AddComment(comment *entity.CommentDB) (string, error) {
	if comment.UserID == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"inserted comment section",
			err,
		)
		return "", err
	}
	comment.ID = primitive.NewObjectID()
	comment.CreatedTime = time.Now()
	comment.Status = "created"
	comment.LikedBy = make([]string, 0)

	result, err := repo.InsertOne(comment)
	if err != nil {
		trestCommon.ECLog3(
			"inserted comment section",
			err,
			logrus.Fields{
				"comment": comment,
			})
		return "", err
	}

	return result, nil
}

func (*commentService) GetComment(postID string) ([]entity.CommentDB, error) {
	if postID == "" {
		err := errors.New("post id missing")
		trestCommon.ECLog2(
			"GetComment section",
			err,
		)
		return []entity.CommentDB{}, err
	}
	return GetInternalComment(postID)
}
func GetInternalComment(postID string) ([]entity.CommentDB, error) {
	comment, err := repo.Find(bson.M{"post_id": postID}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetComment section",
			err,
		)
		return comment, err
	}
	var userID []string
	for _, com := range comment {
		userID = append(userID, com.UserID)
	}
	data, _ := db.GetProfilesForIDs(userID)
	for com := range comment {
		for _, profile := range data {
			if profile.ID.Hex() == comment[com].UserID {
				comment[com].User = profile
			}
			comment[com].ContentURL = utils.CreatePreSignedDownloadUrl(comment[com].ContentURL)
		}

		likedByUsers, _ := db.GetProfilesForIDs(comment[com].LikedBy)
		comment[com].LikedByUsers = append(comment[com].LikedByUsers, likedByUsers...)
	}
	return comment, nil
}

func GetCommentForPosts(postIDs []string) ([]entity.CommentDB, error) {
	filter := bson.A{}
	for _, post := range postIDs {
		filter = append(filter, bson.M{"post_id": post})
	}
	comment, err := repo.Find(bson.M{"$or": filter}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetComment section",
			err,
		)
		return comment, err
	}
	var userID []string
	for _, com := range comment {
		userID = append(userID, com.UserID)
	}
	data, _ := db.GetProfilesForIDs(userID)
	for com := range comment {
		for _, profile := range data {
			if profile.ID.Hex() == comment[com].UserID {
				comment[com].User = profile
			}
		}
		comment[com].ContentURL = utils.CreatePreSignedDownloadUrl(comment[com].ContentURL)
		likedByUsers, _ := db.GetProfilesForIDs(comment[com].LikedBy)
		comment[com].LikedByUsers = append(comment[com].LikedByUsers, likedByUsers...)
	}
	return comment, nil
}

func (*commentService) UpdateComment(comment *entity.CommentDB) (string, error) {
	id, err := primitive.ObjectIDFromHex(comment.CommentID)
	if err != nil {
		return "", err
	}
	filter := bson.M{"_id": id}
	set := bson.M{}
	if comment.Body != "" {
		set["body"] = comment.Body
	}
	if comment.Status != "" {
		set["status"] = comment.Status
	}
	if comment.Tags != "" {
		set["tags"] = comment.Tags
	}
	if comment.ContentType != "" {
		set["content_type"] = comment.ContentType
	}
	if comment.ContentURL != "" {
		set["content_url"] = comment.ContentURL
	}
	set["edited"] = true
	set["updated_time"] = time.Now()
	return repo.UpdateOne(filter, bson.M{"$set": set})
}

func (*commentService) DeleteComment(comment string) error {
	id, err := primitive.ObjectIDFromHex(comment)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	return repo.DeleteOne(filter)
}
func (*commentService) LikeComment(comment, userID string) (string, error) {
	id, err := primitive.ObjectIDFromHex(comment)
	if err != nil {
		return "", err
	}
	filter := bson.M{"_id": id}
	set := bson.M{"$inc": bson.M{"likes": 1}}
	set["$addToSet"] = bson.M{"liked_by": userID}
	set["$set"] = bson.M{"updated_time": time.Now()}

	return repo.UpdateOne(filter, set)
}

func (*commentService) DisLikeComment(comment, userID string) (string, error) {
	id, err := primitive.ObjectIDFromHex(comment)
	if err != nil {
		return "", err
	}
	filter := bson.M{"_id": id}
	set := bson.M{"$inc": bson.M{"likes": -1}}
	set["$pull"] = bson.M{"liked_by": userID}
	set["$set"] = bson.M{"updated_time": time.Now()}

	return repo.UpdateOne(filter, set)
}
