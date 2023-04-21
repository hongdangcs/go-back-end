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

func GetCommentByPostID(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	cr := repository.NewCommentRepository(db, domain.CollectionComments)
	cc := &controller.CommentController{
		CommentUsercase: usecase.NewCommentUsecase(cr, timeout),
		Env:             env,
	}
	group.GET("/comment/post/:id", cc.GetCommentByPostID)

}

func EditComment(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	cr := repository.NewCommentRepository(db, domain.CollectionComments)
	cc := &controller.CommentController{
		CommentUsercase: usecase.NewCommentUsecase(cr, timeout),
		Env:             env,
	}
	group.POST("/comment/:postID", cc.Create)
	group.DELETE("/comment/:id", cc.Delete)
	group.PUT("/comment/:id", cc.Edit)

}
