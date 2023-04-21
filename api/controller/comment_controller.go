package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hongdangcseiu/go-back-end/bootstrap"
	"github.com/hongdangcseiu/go-back-end/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type CommentController struct {
	CommentUsercase domain.CommentUsecase
	Env             *bootstrap.Env
}

func (cc *CommentController) Create(c *gin.Context) {
	var request domain.Comment
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("x-user-id")
	postID := c.Param("postID")

	timeString := bootstrap.GetTimeNow()

	comment := domain.Comment{
		ID:         primitive.NewObjectID(),
		Content:    request.Content,
		DateCreate: timeString,
	}
	comment.UserID, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	comment.PostID, err = primitive.ObjectIDFromHex(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = cc.CommentUsercase.Create(c, &comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Comment created successfully",
	})
}

func (cc *CommentController) GetCommentByPostID(c *gin.Context) {
	postID := c.Param("id")

	comments, err := cc.CommentUsercase.GetCommentByPostID(c, postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}
func (cc *CommentController) Delete(c *gin.Context) {
	userID := c.GetString("x-user-id")
	commentID := c.Param("id")

	commentToDelete, err := cc.CommentUsercase.GetCommentByID(c, commentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	if commentToDelete.UserID.Hex() != userID {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Not Authorized!!! "})
		return
	}

	if err := cc.CommentUsercase.Delete(c, &commentToDelete); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post Delete successfully"})
}

func (cc *CommentController) Edit(c *gin.Context) {
	userID := c.GetString("x-user-id")
	commentID := c.Param("id")

	var comment domain.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	commentToEdit, err := cc.CommentUsercase.GetCommentByID(c, commentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	if commentToEdit.UserID.Hex() != userID {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Not Authorized!!! "})
		return
	}

	if err := cc.CommentUsercase.Edit(c, commentID, &comment); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}
