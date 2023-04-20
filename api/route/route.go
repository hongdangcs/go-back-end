package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hongdangcseiu/go-back-end/api/middleware"
	"github.com/hongdangcseiu/go-back-end/bootstrap"
	"github.com/hongdangcseiu/go-back-end/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	//my edit
	GetPostRouter(env, timeout, db, publicRouter)
	GetUserRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProfileRouter(env, timeout, db, protectedRouter)
	NewTaskRouter(env, timeout, db, protectedRouter)

	//my edit
	NewPostRouter(env, timeout, db, protectedRouter)
	EditUserRouter(env, timeout, db, protectedRouter)
	EditPostRouter(env, timeout, db, protectedRouter)
}
