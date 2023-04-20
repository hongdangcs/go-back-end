package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hongdangcseiu/go-back-end/api/controller"
	"github.com/hongdangcseiu/go-back-end/bootstrap"
	"github.com/hongdangcseiu/go-back-end/domain"
	"github.com/hongdangcseiu/go-back-end/mongo"
	"github.com/hongdangcseiu/go-back-end/repository"
	"github.com/hongdangcseiu/go-back-end/usecase"
	"time"
)

func GetPostRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	pr := repository.NewPostRepository(db, domain.CollectionPost)
	pc := &controller.PostController{
		PostUsercase: usecase.NewPostUsecase(pr, timeout),
		Env:          env,
	}
	group.GET("/post/:id", pc.GetPostById)
	group.GET("/post/", pc.GetPost)
	group.GET("/post/user/:id", pc.GetPostByUserId)

}

func NewPostRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	pr := repository.NewPostRepository(db, domain.CollectionPost)
	pc := &controller.PostController{
		PostUsercase: usecase.NewPostUsecase(pr, timeout),
		Env:          env,
	}
	group.POST("/post/new", pc.Create)
}
