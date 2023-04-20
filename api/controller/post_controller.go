package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hongdangcseiu/go-back-end/bootstrap"
	"github.com/hongdangcseiu/go-back-end/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
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

	timeString := bootstrap.GetTimeNow()

	post := domain.Post{
		ID:         primitive.NewObjectID(),
		Title:      request.Title,
		Content:    request.Content,
		DateCreate: timeString,
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

func (pc *PostController) Edit(c *gin.Context) {
	userID := c.GetString("x-user-id")
	postID := c.Param("id")

	var post domain.Post
	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	postToEdit, err := pc.PostUsercase.GetPostByID(c, postID)
	log.Print("get post by id: ", postID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if postToEdit.UserID.Hex() != userID {
		log.Print(userID, " edit: ", postToEdit.UserID.Hex())
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Not Authorized!!! "})
		return
	}

	if err := pc.PostUsercase.Edit(c, postID, &post); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

func (u *PostController) GetPostByUserId(c *gin.Context) {
	userID := c.Param("id")

	posts, err := u.PostUsercase.GetPostByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (u *PostController) GetPostById(c *gin.Context) {
	/*
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

	*/
	postID := c.Param("id")

	post, err := u.PostUsercase.GetPostByID(c, postID)
	log.Print("get post by id: ", postID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
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

func (u *PostController) GetPostByCategory(c *gin.Context) {
	category := c.Param("category")

	posts, err := u.PostUsercase.GetPostByCategory(c, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}
