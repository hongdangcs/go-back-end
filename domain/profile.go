package domain

import "context"

type Profile struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"username"`
	Bio      string `json:"bio"`
}

type ProfileUsecase interface {
	GetProfileByID(c context.Context, userID string) (*Profile, error)
}
