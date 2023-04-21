package usecase

import (
	"context"
	"github.com/hongdangcseiu/go-back-end/domain"
	"time"
)

type commentUsecase struct {
	commentRepository domain.CommentRepository
	contextTimeout    time.Duration
}

func NewCommentUsecase(commentRepository domain.CommentRepository, timeout time.Duration) domain.CommentUsecase {
	return &commentUsecase{
		commentRepository: commentRepository,
		contextTimeout:    timeout,
	}
}

func (c2 *commentUsecase) Create(c context.Context, comment *domain.Comment) error {
	ctx, cancel := context.WithTimeout(c, c2.contextTimeout)
	defer cancel()
	return c2.commentRepository.Create(ctx, comment)
}

func (c2 *commentUsecase) Edit(c context.Context, commentID string, comment *domain.Comment) error {
	ctx, cancel := context.WithTimeout(c, c2.contextTimeout)
	defer cancel()
	return c2.commentRepository.Edit(ctx, commentID, comment)
}

func (c2 *commentUsecase) GetCommentByPostID(c context.Context, postID string) ([]domain.Comment, error) {
	ctx, cancel := context.WithTimeout(c, c2.contextTimeout)
	defer cancel()
	return c2.commentRepository.GetCommentByPostID(ctx, postID)
}

func (c2 *commentUsecase) Delete(c context.Context, comment *domain.Comment) error {
	ctx, cancel := context.WithTimeout(c, c2.contextTimeout)
	defer cancel()
	return c2.commentRepository.Delete(ctx, comment)
}

func (c2 *commentUsecase) GetCommentByID(c context.Context, commentID string) (domain.Comment, error) {
	ctx, cancel := context.WithTimeout(c, c2.contextTimeout)
	defer cancel()
	return c2.commentRepository.GetCommentByID(ctx, commentID)
}
