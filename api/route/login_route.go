package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hongdangcseiu/go-back-end/api/controller"
	"github.com/hongdangcseiu/go-back-end/bootstrap"
	"github.com/hongdangcseiu/go-back-end/domain"
	"github.com/hongdangcseiu/go-back-end/mongo"
	"github.com/hongdangcseiu/go-back-end/repository"
	"github.com/hongdangcseiu/go-back-end/usecase"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/api/login/", lc.Login)
}
