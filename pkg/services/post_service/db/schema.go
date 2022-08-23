package db

import (
	entity2 "kinexx_backend/pkg/entity"
	"kinexx_backend/pkg/services/post_service/entity"
)

type PostService interface {
	AddPost(post *entity.PostDB) (string, error)
	UpdatePost(post *entity.PostDB) (string, error)
	GetPost(postType string) ([]entity.PostDB, error)
	GetPostByID(userID, postType string) ([]entity.PostDB, error)
	GetPostByPostID(postID string) (entity.PostDB, error)
	LikePost(post, userID string) (string, error)
	DeletePost(post string) error
	LikedPost(post string) ([]entity2.ProfileDB, error)
	SharedPost(post string) ([]entity2.ProfileDB, error)
	DisLikePost(post, userID string) (string, error)
	GetUserData(user string) (entity2.ProfileDB, []entity.PostDB, error)
}
