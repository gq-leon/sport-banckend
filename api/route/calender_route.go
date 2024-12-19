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

func NewCalenderRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	cr := repository.NewCalenderRepository(db, domain.CollectionTrainPlan)
	cc := &controller.CalenderController{
		CalenderUseCase: usecase.NewCalenderUseCase(cr, timeout),
	}
	{
		group.POST("/calender/workouts", cc.Workouts)
		group.POST("/calender/month-workouts", cc.MonthWorkouts)
	}
}
