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

func NewStatsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	sr := repository.NewStatsRepository(db, domain.CollectionTrainPlan)
	sc := &controller.StatsController{
		StatsUseCase: usecase.NewStatsUseCase(sr, timeout),
	}
	{
		group.GET("/stats/profile", sc.Profile)
	}
}
