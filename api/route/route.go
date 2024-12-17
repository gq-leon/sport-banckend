package route

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gq-leon/sport-backend/api/middleware"
	"github.com/gq-leon/sport-backend/bootstrap"
	"github.com/gq-leon/sport-backend/pkg/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, serve *gin.Engine) {
	// CORS Middleware
	serve.Use(middleware.CorsMiddleware())

	publicRouter := serve.Group("api/v1")
	NewUserRouter(env, timeout, db, publicRouter)

	protectedRouter := serve.Group("api/v1")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	NewTrainPlanRouter(env, timeout, db, protectedRouter)
	NewAttendanceRouter(env, timeout, db, protectedRouter)
}
