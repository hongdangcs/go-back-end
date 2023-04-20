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
