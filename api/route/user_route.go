package route

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gq-leon/sport-backend/api/controller"
	"github.com/gq-leon/sport-backend/bootstrap"
	"github.com/gq-leon/sport-backend/domain"
	"github.com/gq-leon/sport-backend/pkg/mongo"
	"github.com/gq-leon/sport-backend/repository"
	"github.com/gq-leon/sport-backend/usecase"
)

func NewUserRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	uc := &controller.UserController{
		Env:         env,
		UserUseCase: usecase.NewUserUseCase(ur, timeout),
	}

	group.POST("/signup", uc.Signup)
	group.POST("/login", uc.Login)
}
