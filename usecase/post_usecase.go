package usecase

import (
	"context"
	"github.com/hongdangcseiu/go-back-end/domain"
	"time"
)

type postUsecase struct {
	postRepository domain.PostRepository
	contextTimeout time.Duration
}

func NewPostUsecase(postRepository domain.PostRepository, timeout time.Duration) domain.PostUsecase {
	return &postUsecase{
		postRepository: postRepository,
		contextTimeout: timeout,
	}
}

func (pu *postUsecase) Create(c context.Context, post *domain.Post) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.postRepository.Create(ctx, post)
}

func (pu *postUsecase) Edit(c context.Context, postId string, post *domain.Post) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.postRepository.Edit(ctx, postId, post)
}

func (pu *postUsecase) Search(c context.Context, query string) ([]domain.Post, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.postRepository.Search(ctx, query)
}

func (pu *postUsecase) GetPostByUserID(c context.Context, userID string) ([]domain.Post, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.postRepository.GetPostByUserID(ctx, userID)
}

func (pu *postUsecase) GetPost(c context.Context) ([]domain.Post, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.postRepository.GetPost(ctx)
}

func (pu *postUsecase) GetPostByID(c context.Context, postId string) (domain.Post, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.postRepository.GetPostByID(ctx, postId)
}

func (pu *postUsecase) GetPostByCategory(c context.Context, category string) ([]domain.Post, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.postRepository.GetPostByCategory(ctx, category)
}

func (pu *postUsecase) Delete(c context.Context, post *domain.Post) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.postRepository.Delete(ctx, post)
}
