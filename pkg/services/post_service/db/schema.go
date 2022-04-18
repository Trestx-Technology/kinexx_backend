package db

import "kinexx_backend/pkg/entity"

type PostService interface {
	AddPost(post *entity.PostDB) (string, error)
	UpdatePost(post *entity.PostDB) (string, error)
	GetPost(postType string) ([]entity.PostDB, error)
	GetPostByID(userID, postType string) ([]entity.PostDB, error)
	GetPostByPostID(postID string) (entity.PostDB, error)
	LikePost(post, userID string) (string, error)
	DeletePost(post string) error
	LikedPost(post string) ([]entity.ProfileDB, error)
	SharedPost(post string) ([]entity.ProfileDB, error)
	DisLikePost(post, userID string) (string, error)
	GetUserData(user string) (entity.ProfileDB, []entity.PostDB, error)
}
