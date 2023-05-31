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
	group.GET("/api/post/:id/", pc.GetPostById)
	group.GET("/api/post/", pc.GetPost)
	group.GET("/api/post/user/:id/", pc.GetPostByUserId)
	group.GET("/api/post/category/:category/", pc.GetPostByCategory)
	group.GET("/api/post/search/:query/", pc.Search)

}

func NewPostRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	pr := repository.NewPostRepository(db, domain.CollectionPost)
	pc := &controller.PostController{
		PostUsercase: usecase.NewPostUsecase(pr, timeout),
		Env:          env,
	}
	group.POST("/api/post/new", pc.Create)
}

func EditPostRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	pr := repository.NewPostRepository(db, domain.CollectionPost)
	pc := &controller.PostController{
		PostUsercase: usecase.NewPostUsecase(pr, timeout),
		Env:          env,
	}
	group.PUT("/api/post/:id", pc.Edit)
	group.DELETE("/api/delete/:id", pc.Delete)
}
