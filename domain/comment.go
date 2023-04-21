package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionComments = "comments"
)

type Comment struct {
	ID         primitive.ObjectID `bson:"_id" json:"-"`
	UserID     primitive.ObjectID `bson:"userID" json:"userID"`
	PostID     primitive.ObjectID `bson:"postID" json:"postID"`
	Content    string             `bson:"content" form:"content" binding:"required" json:"content"`
	DateCreate string             `bson:"date_create" form:"date_create" json:"date_create"`
}

type CommentRepository interface {
	Create(c context.Context, comment *Comment) error
	Edit(c context.Context, commentID string, comment *Comment) error
	GetCommentByPostID(c context.Context, postID string) ([]Comment, error)
	Delete(c context.Context, comment *Comment) error
}

type CommentUsecase interface {
	Create(c context.Context, comment *Comment) error
	Edit(c context.Context, commentID string, comment *Comment) error
	GetCommentByPostID(c context.Context, postID string) ([]Comment, error)
	Delete(c context.Context, comment *Comment) error
}
