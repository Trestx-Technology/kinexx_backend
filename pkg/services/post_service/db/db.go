package db

import (
	"kinexx_backend/pkg/entity"
	post "kinexx_backend/pkg/repository/posts"
	comment_db "kinexx_backend/pkg/services/comment_service/db"
	"kinexx_backend/pkg/services/profile_service/db"
	share "kinexx_backend/pkg/services/share_service/db"

	"strings"

	"errors"
	"time"

	"github.com/aekam27/trestCommon"

	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = post.NewPostRepository("posts")
)

type postService struct{}

func NewPostService(repository post.PostRepository) PostService {
	repo = repository
	return &postService{}
}

func (*postService) AddPost(post *entity.PostDB) (string, error) {
	if post.UserID == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"inserted post section",
			err,
		)
		return "", err
	}
	post.ID = primitive.NewObjectID()
	post.CreatedTime = time.Now()
	post.Status = "created"
	post.LikedBy = make([]string, 0)
	result, err := repo.InsertOne(post)
	if err != nil {
		trestCommon.ECLog3(
			"inserted post section",
			err,
			logrus.Fields{
				"post": post,
			})
		return "", err
	}

	return result, nil
}

func (*postService) GetPost(postType string) ([]entity.PostDB, error) {
	filter := bson.M{}
	if postType != "" {
		filter["post_type"] = postType
	}
	post, err := repo.Find(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetPost section",
			err,
		)
		return post, err
	}
	var userID []string
	var postID []string
	var posts []entity.PostDB
	for id := range post {
		newUrl := createPreSignedDownloadUrl(post[id].ContentURL)
		post[id].ContentURL = newUrl
		userID = append(userID, post[id].UserID)
		postID = append(postID, post[id].ID.Hex())
		if len(post[id].LikedBy) > 0 {
			likedByUsers, _ := db.GetProfilesForIDs(post[id].LikedBy)
			post[id].LikedByUsers = append(post[id].LikedByUsers, likedByUsers...)
		}
		data, _ := db.GetProfilesForIDs([]string{post[id].UserID})
		post[id].User = data[0]

		sharedByUsersIDs, _ := share.GetShareByPostID(post[id].ID.Hex())
		post[id].ShareCount = len(sharedByUsersIDs)
		posts = append(posts, post[id])
		if len(sharedByUsersIDs) > 0 {
			sharedByUsers, _ := db.GetProfilesForIDs(sharedByUsersIDs)
			for _, use := range sharedByUsers {
				post[id].SharedByUsers = []entity.ProfileDB{use}
				post[id].SharedInstance = true
				post[id].ShareCount = len(sharedByUsersIDs)
				posts = append(posts, post[id])
			}
		}

	}
	data, _ := db.GetProfilesForIDs(userID)
	comments, _ := comment_db.GetCommentForPosts(postID)
	for id := range posts {
		for _, profile := range data {
			if profile.ID.Hex() == posts[id].UserID {
				posts[id].User = profile
			}
		}
		for _, comment := range comments {
			if comment.PostID == posts[id].ID.Hex() {
				posts[id].Comment = append(posts[id].Comment, comment)
			}
		}
	}
	return posts, nil
}
func (*postService) LikedPost(postID string) ([]entity.ProfileDB, error) {
	id, _ := primitive.ObjectIDFromHex(postID)
	filter := bson.M{"_id": id}

	post, err := repo.FindOne(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetPost section",
			err,
		)
		return []entity.ProfileDB{}, err
	}
	return db.GetProfilesForIDs(post.LikedBy)

}
func (*postService) SharedPost(post string) ([]entity.ProfileDB, error) {
	sharedByUsersIDs, _ := share.GetShareByPostID(post)
	return db.GetProfilesForIDs(sharedByUsersIDs)

}
func (*postService) GetPostByPostID(postI string) (entity.PostDB, error) {
	id, _ := primitive.ObjectIDFromHex(postI)
	post, err := repo.FindOne(bson.M{"_id": id}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetPost section",
			err,
		)
		return post, err
	}
	var userID []string
	var postID []string
	newUrl := createPreSignedDownloadUrl(post.ContentURL)
	post.ContentURL = newUrl
	userID = append(userID, post.UserID)
	postID = append(postID, post.ID.Hex())
	likedByUsers, _ := db.GetProfilesForIDs(post.LikedBy)
	post.LikedByUsers = append(post.LikedByUsers, likedByUsers...)
	sharedByUsersIDs, _ := share.GetShareByPostID(post.ID.Hex())
	sharedByUsers, _ := db.GetProfilesForIDs(sharedByUsersIDs)
	post.SharedByUsers = append(post.SharedByUsers, sharedByUsers...)
	data, _ := db.GetProfilesForIDs([]string{post.UserID})
	post.User = data[0]
	datas, _ := db.GetProfilesForIDs(userID)
	comments, _ := comment_db.GetCommentForPosts(postID)

	for _, profile := range datas {
		if profile.ID.Hex() == post.UserID {
			post.User = profile
		}
	}
	for _, comment := range comments {
		if comment.PostID == post.ID.Hex() {
			post.Comment = append(post.Comment, comment)
		}
	}

	return post, nil
}
func (*postService) GetPostByID(user string, postType string) ([]entity.PostDB, error) {
	filter := bson.M{"user_id": user}
	if postType != "" {
		filter["post_type"] = postType
	}
	post, err := repo.Find(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetPost section",
			err,
		)
		return post, err
	}
	var userID []string
	var postID []string
	var posts []entity.PostDB
	for id := range post {
		newUrl := createPreSignedDownloadUrl(post[id].ContentURL)
		post[id].ContentURL = newUrl
		userID = append(userID, post[id].UserID)
		postID = append(postID, post[id].ID.Hex())
		if len(post[id].LikedBy) > 0 {
			likedByUsers, _ := db.GetProfilesForIDs(post[id].LikedBy)
			post[id].LikedByUsers = append(post[id].LikedByUsers, likedByUsers...)
		}
		data, _ := db.GetProfilesForIDs([]string{post[id].UserID})
		post[id].User = data[0]
		posts = append(posts, post[id])
		sharedByUsersIDs, _ := share.GetShareByPostID(post[id].ID.Hex())
		if len(sharedByUsersIDs) > 0 {
			sharedByUsers, _ := db.GetProfilesForIDs(sharedByUsersIDs)
			for _, use := range sharedByUsers {
				post[id].SharedByUsers = []entity.ProfileDB{use}
				post[id].SharedInstance = true
				posts = append(posts, post[id])
			}
		}

	}
	data, _ := db.GetProfilesForIDs(userID)
	comments, _ := comment_db.GetCommentForPosts(postID)
	for id := range posts {
		for _, profile := range data {
			if profile.ID.Hex() == posts[id].UserID {
				posts[id].User = profile
			}
		}
		for _, comment := range comments {
			if comment.PostID == posts[id].ID.Hex() {
				posts[id].Comment = append(posts[id].Comment, comment)
			}
		}
	}
	return posts, nil
}

func (*postService) GetUserData(user string) (entity.ProfileDB, []entity.PostDB, error) {
	users, err := db.GetProfilesForIDs([]string{user})
	if err != nil {
		trestCommon.ECLog2(
			"GetPost section",
			err,
		)
		return entity.ProfileDB{}, []entity.PostDB{}, err
	}
	users[0].Password = ""
	filter := bson.M{"user_id": user}

	post, err := repo.Find(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetPost section",
			err,
		)
		return users[0], post, err
	}
	var userID []string
	var postID []string
	var posts []entity.PostDB
	for id := range post {
		newUrl := createPreSignedDownloadUrl(post[id].ContentURL)
		post[id].ContentURL = newUrl
		userID = append(userID, post[id].UserID)
		postID = append(postID, post[id].ID.Hex())
		if len(post[id].LikedBy) > 0 {
			likedByUsers, _ := db.GetProfilesForIDs(post[id].LikedBy)
			post[id].LikedByUsers = append(post[id].LikedByUsers, likedByUsers...)
		}
		data, _ := db.GetProfilesForIDs([]string{post[id].UserID})
		post[id].User = data[0]
		posts = append(posts, post[id])
		sharedByUsersIDs, _ := share.GetShareByPostID(post[id].ID.Hex())
		if len(sharedByUsersIDs) > 0 {
			sharedByUsers, _ := db.GetProfilesForIDs(sharedByUsersIDs)
			for _, use := range sharedByUsers {
				post[id].SharedByUsers = []entity.ProfileDB{use}
				post[id].SharedInstance = true
				posts = append(posts, post[id])
			}
		}

	}
	data, _ := db.GetProfilesForIDs(userID)
	comments, _ := comment_db.GetCommentForPosts(postID)
	for id := range posts {
		for _, profile := range data {
			if profile.ID.Hex() == posts[id].UserID {
				posts[id].User = profile
			}
		}
		for _, comment := range comments {
			if comment.PostID == posts[id].ID.Hex() {
				posts[id].Comment = append(posts[id].Comment, comment)
			}
		}
	}
	return users[0], posts, nil
}
func GetUserDataInternal(user string) (entity.ProfileDB, []entity.PostDB, error) {
	users, err := db.GetProfilesForIDs([]string{user})
	if err != nil {
		trestCommon.ECLog2(
			"GetPost section",
			err,
		)
		return entity.ProfileDB{}, []entity.PostDB{}, err
	}
	users[0].Password = ""
	filter := bson.M{"user_id": user}

	post, err := repo.Find(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetPost section",
			err,
		)
		return users[0], post, err
	}
	var userID []string
	var postID []string
	var posts []entity.PostDB
	for id := range post {
		newUrl := createPreSignedDownloadUrl(post[id].ContentURL)
		post[id].ContentURL = newUrl
		userID = append(userID, post[id].UserID)
		postID = append(postID, post[id].ID.Hex())
		if len(post[id].LikedBy) > 0 {
			likedByUsers, _ := db.GetProfilesForIDs(post[id].LikedBy)
			post[id].LikedByUsers = append(post[id].LikedByUsers, likedByUsers...)
		}
		data, _ := db.GetProfilesForIDs([]string{post[id].UserID})
		post[id].User = data[0]
		posts = append(posts, post[id])
		sharedByUsersIDs, _ := share.GetShareByPostID(post[id].ID.Hex())
		if len(sharedByUsersIDs) > 0 {
			sharedByUsers, _ := db.GetProfilesForIDs(sharedByUsersIDs)
			for _, use := range sharedByUsers {
				post[id].SharedByUsers = []entity.ProfileDB{use}
				post[id].SharedInstance = true
				posts = append(posts, post[id])
			}
		}

	}
	data, _ := db.GetProfilesForIDs(userID)
	comments, _ := comment_db.GetCommentForPosts(postID)
	for id := range posts {
		for _, profile := range data {
			if profile.ID.Hex() == posts[id].UserID {
				posts[id].User = profile
			}
		}
		for _, comment := range comments {
			if comment.PostID == posts[id].ID.Hex() {
				posts[id].Comment = append(posts[id].Comment, comment)
			}
		}
	}
	return users[0], posts, nil
}
func GetPostID(postId string) (entity.PostDB, error) {
	id, _ := primitive.ObjectIDFromHex(postId)

	filter := bson.M{"_id": id}
	post, err := repo.FindOne(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetPost section",
			err,
		)
		return post, err
	}
	var userID []string
	var postID []string

	newUrl := createPreSignedDownloadUrl(post.ContentURL)
	post.ContentURL = newUrl
	userID = append(userID, post.UserID)
	postID = append(postID, post.ID.Hex())
	data, _ := db.GetProfilesForIDs(userID)
	comments, _ := comment_db.GetCommentForPosts(postID)
	post.User = data[0]
	post.Comment = append(post.Comment, comments...)

	return post, nil
}

func (*postService) UpdatePost(post *entity.PostDB) (string, error) {
	id, err := primitive.ObjectIDFromHex(post.PostID)
	if err != nil {
		return "", err
	}
	filter := bson.M{"_id": id}
	set := bson.M{}
	if post.Body != "" {
		set["body"] = post.Body
	}
	if post.Status != "" {
		set["status"] = post.Status
	}
	if post.Tags != "" {
		set["tags"] = post.Tags
	}
	if post.ContentType != "" {
		set["content_type"] = post.ContentType
	}
	if post.ContentURL != "" {
		set["content_url"] = post.ContentURL
	}
	set["edited"] = true
	set["updated_time"] = time.Now()

	return repo.UpdateOne(filter, bson.M{"$set": set})
}
func (*postService) LikePost(post, userID string) (string, error) {
	id, err := primitive.ObjectIDFromHex(post)
	if err != nil {
		return "", err
	}
	filter := bson.M{"_id": id}
	set := bson.M{"$inc": bson.M{"likes": 1}}
	set["$addToSet"] = bson.M{"liked_by": userID}
	set["$set"] = bson.M{"updated_time": time.Now()}

	return repo.UpdateOne(filter, set)
}
func (*postService) DeletePost(post string) error {
	id, err := primitive.ObjectIDFromHex(post)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}

	return repo.DeleteOne(filter)
}
func (*postService) DisLikePost(post, userID string) (string, error) {
	id, err := primitive.ObjectIDFromHex(post)
	if err != nil {
		return "", err
	}
	filter := bson.M{"_id": id}
	set := bson.M{"$inc": bson.M{"likes": -1}}
	set["$pull"] = bson.M{"liked_by": userID}
	set["$set"] = bson.M{"updated_time": time.Now()}

	return repo.UpdateOne(filter, set)
}
func createPreSignedDownloadUrl(url string) string {
	s := strings.Split(url, "?")
	if len(s) > 0 {
		o := strings.Split(s[0], "/")
		if len(o) > 3 {
			fileName := o[4]
			path := o[3]
			downUrl, _ := trestCommon.PreSignedDownloadUrlAWS(fileName, path)
			return downUrl
		}
	}
	return ""
}
