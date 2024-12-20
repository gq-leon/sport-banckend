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

func NewAttendanceRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ar := repository.NewAttendanceRepository(db, domain.CollectionAttendance)
	ac := &controller.AttendanceController{
		AttendanceUseCase: usecase.NewAttendanceUseCase(ar, timeout),
	}
	{
		group.POST("/attendance", ac.Create)
		group.GET("/attendance", ac.List)
		group.POST("/attendance/back-check-in", ac.BackDateCheckIn)
	}
}
