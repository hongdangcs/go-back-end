package main

import (
	"time"

	"github.com/gin-gonic/gin"
	route "github.com/hongdangcseiu/go-back-end/api/route"
	"github.com/hongdangcseiu/go-back-end/bootstrap"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)
}
