package usecase

import (
	"context"
	"github.com/hongdangcseiu/go-back-end/domain"
	"time"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (uu userUsecase) GetUserByUserName(c context.Context, userID string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.GetUserByUserName(ctx, userID)
}

func (uu *userUsecase) EditUser(c context.Context, userID string, name string, bio string, profilePic string, socialMedia string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	user, err := uu.userRepository.GetUserByID(c, userID)
	if err != nil {
		return err
	}

	if name != "" {
		user.Name = name
	}
	if bio != "" {
		user.Bio = bio
	}
	if profilePic != "" {
		user.ProfilePic = profilePic
	}
	if socialMedia != "" {
		user.SocialMedia = socialMedia
	}

	err = uu.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
