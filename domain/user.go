package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Email       string             `bson:"email"`
	Password    string             `bson:"password"`
	UserName    string             `bson:"username"`
	ProfilePic  string             `bson:"profile_pic"`
	Bio         string             `bson:"bio"`
	SocialMedia string             `bson:"social_media"`
}

type UserNameRequest struct {
	UserName string `bson:"username"`
}

type EditUserRequest struct {
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	ProfilePic  string `json:"profile_pic"`
	SocialMedia string `json:"social_media"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetUserByEmail(c context.Context, email string) (User, error)
	GetUserByID(c context.Context, id string) (User, error)
	GetUserByUserName(c context.Context, username string) (User, error)
	UpdateUser(c context.Context, user User) error
}

type UserUsecase interface {
	GetUserByUserName(c context.Context, id string) (User, error)
	GetUserByUserId(c context.Context, id string) (User, error)
	EditUser(c context.Context, id string, name string, bio string, pic string, media string) error
}
