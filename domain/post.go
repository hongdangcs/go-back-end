package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionPost = "posts"
)

type Post struct {
	ID                  primitive.ObjectID `bson:"_id" json:"-"`
	Title               string             `bson:"title" form:"title" binding:"required" json:"title"`
	UserID              primitive.ObjectID `bson:"userID" json:"userID"`
	Content             string             `bson:"content" form:"content" binding:"required" json:"content"`
	DateCreate          string             `bson:"date_create" form:"date_create" json:"date_create"`
	DateUpdate          string             `bson:"date_update" form:"date_update" json:"date_update"`
	Categories          []string           `bson:"categories" form:"categories" json:"categories"`
	ApprovedByModerator string             `bson:"approved" form:"approved" json:"approved"`
}

type PostRepository interface {
	Create(c context.Context, task *Post) error
	GetPost(c context.Context) ([]Post, error)
	GetPostByID(c context.Context, userID string) (Post, error)
	GetPostByUserID(c context.Context, userID string) ([]Post, error)
}

type PostUsecase interface {
	Create(c context.Context, task *Post) error
	GetPost(c context.Context) ([]Post, error)
	GetPostByID(c context.Context, userID string) (Post, error)
	GetPostByUserID(c context.Context, userID string) ([]Post, error)
}