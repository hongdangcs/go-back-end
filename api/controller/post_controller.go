package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hongdangcseiu/go-back-end/bootstrap"
	"github.com/hongdangcseiu/go-back-end/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type PostController struct {
	PostUsercase domain.PostUsecase
	Env          *bootstrap.Env
}

func (pc *PostController) Create(c *gin.Context) {
	var request domain.Post
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("x-user-id")

	post := domain.Post{
		ID:      primitive.NewObjectID(),
		Title:   request.Title,
		Content: request.Content,
	}
	post.UserID, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = pc.PostUsercase.Create(c, &post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Post created successfully",
	})
}

func (u *PostController) GetPostByUserId(c *gin.Context) {
	userID := c.GetString("x-user-id")

	posts, err := u.PostUsercase.GetPostByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (u *PostController) GetPostById(c *gin.Context) {
	var request domain.Post

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	post, err := u.PostUsercase.GetPostByID(c, request.ID.Hex())
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Post not found!!!"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (u *PostController) GetPost(c *gin.Context) {

	posts, err := u.PostUsercase.GetPost(c)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Post not found!!!"})
		return
	}

	c.JSON(http.StatusOK, posts)
}
