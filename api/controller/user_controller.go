package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hongdangcseiu/go-back-end/bootstrap"
	"github.com/hongdangcseiu/go-back-end/domain"
	"net/http"
)

type UserController struct {
	UserUsecase domain.UserUsecase
	Env         *bootstrap.Env
}

func (uc *UserController) GetUserByUserName(c *gin.Context) {
	/*
		var request domain.UserNameRequest

		err := c.ShouldBind(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
			return
		}

		// get username from get request here

		user, err := uc.UserUsecase.GetUserByUserName(c, request.UserName)
		log.Println("check username: ", request.UserName)
		if err == nil {
			log.Print("exists username")
			c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User name already exists!"})
			return
		}
		c.JSON(http.StatusOK, domain.User{Name: user.Name,
			Bio:         user.Bio,
			UserName:    user.UserName,
			ProfilePic:  user.ProfilePic,
			SocialMedia: user.SocialMedia})

	*/

	username := c.Param("username")

	user, err := uc.UserUsecase.GetUserByUserName(c, username)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found"})
		return
	}

	c.JSON(http.StatusOK, domain.User{Name: user.Name,
		Bio:         user.Bio,
		UserName:    user.UserName,
		ProfilePic:  user.ProfilePic,
		SocialMedia: user.SocialMedia})
}

func (uc *UserController) EditUser(c *gin.Context) {
	userID := c.GetString("x-user-id")

	var request domain.EditUserRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = uc.UserUsecase.EditUser(c, userID, request.Name, request.Bio, request.ProfilePic, request.SocialMedia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "User updated successfully"})
}
