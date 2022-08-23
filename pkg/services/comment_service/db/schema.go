package db

import (
	"kinexx_backend/pkg/entity"
)

type CommentService interface {
	AddComment(comment *entity.CommentDB) (string, error)
	UpdateComment(comment *entity.CommentDB) (string, error)
	GetComment(postID string) ([]entity.CommentDB, error)
	LikeComment(comment, userID string) (string, error)
	DisLikeComment(comment, userID string) (string, error)
	DeleteComment(comment string) error
}
