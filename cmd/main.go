package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gq-leon/sport-backend/api/route"
	"github.com/gq-leon/sport-backend/bootstrap"
	"github.com/gq-leon/sport-backend/pkg/redis"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	timeout := time.Duration(env.ContextTimeout) * time.Second

	if err := redis.InitRedis(env); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	serve := gin.Default()
	route.Setup(env, timeout, db, serve)
	serve.Run(env.ServerAddress)
}
