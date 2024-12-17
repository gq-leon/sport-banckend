package main

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gq-leon/sport-backend/api/route"
	"github.com/gq-leon/sport-backend/bootstrap"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	timeout := time.Duration(env.ContextTimeout) * time.Second

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	serve := gin.Default()
	route.Setup(env, timeout, db, serve)
	serve.Run(env.ServerAddress)
}
