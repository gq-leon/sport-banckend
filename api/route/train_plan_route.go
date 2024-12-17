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

func NewTrainPlanRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tpr := repository.NewTrainPlanRepository(db, domain.CollectionTrainPlan)
	tpc := &controller.TrainPlanController{
		Env:              env,
		TrainPlanUseCase: usecase.NewTrainPlanUseCase(tpr, timeout),
	}

	{
		group.GET("/train-plan/today-workout", tpc.TodayWorkout)
		group.POST("/train-plan", tpc.Create)
		group.POST("/train-plan/update", tpc.Update)
		group.POST("/train-plan/del", tpc.Delete)
		group.POST("/train-plan/complete", tpc.UpdateCompletion)
	}
}
