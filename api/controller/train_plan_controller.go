package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gq-leon/sport-backend/bootstrap"
	"github.com/gq-leon/sport-backend/domain"
	"github.com/gq-leon/sport-backend/internal/util"
)

type TrainPlanController struct {
	Env              *bootstrap.Env
	TrainPlanUseCase domain.TrainPlanUseCase
}

func (tpc *TrainPlanController) Create(c *gin.Context) {
	var request domain.AddPlanRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		domain.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	var (
		userId       = util.GetUserId(c)
		idFromHex, _ = primitive.ObjectIDFromHex(userId)
		date         = time.Now().Format(domain.TrainPlanFormat)
	)

	if err := tpc.TrainPlanUseCase.Create(c, domain.TrainPlan{
		ID:        primitive.NewObjectID(),
		UserID:    idFromHex,
		Date:      date,
		Name:      request.Name,
		Category:  request.Category,
		Weight:    request.Weight,
		Reps:      request.Reps,
		Duration:  request.Duration,
		Distance:  request.Distance,
		Sets:      request.Sets,
		Completed: request.Completed,
	}); err != nil {
		domain.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	plans, err := tpc.TrainPlanUseCase.GetPlansByDate(c, userId, date)
	if err != nil {
		domain.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	domain.SuccessResponse(c, domain.TodayWorkoutResponse{
		Date:      date,
		Exercises: plans,
	})
}

func (tpc *TrainPlanController) Delete(c *gin.Context) {
	var request domain.DelPlanRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		domain.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := tpc.TrainPlanUseCase.Delete(c, request.Id); err != nil {
		domain.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	domain.SuccessResponse(c, nil)
}

func (tpc *TrainPlanController) TodayWorkout(c *gin.Context) {

	var (
		userId = util.GetUserId(c)
		date   = time.Now().Format(domain.TrainPlanFormat)
	)

	plans, err := tpc.TrainPlanUseCase.GetPlansByDate(c, userId, date)
	if err != nil {
		domain.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	domain.SuccessResponse(c, domain.TodayWorkoutResponse{
		Date:      date,
		Exercises: plans,
	})
}

func (tpc *TrainPlanController) UpdateCompletion(c *gin.Context) {
	var request domain.UpdateProgressPlanRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		domain.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	trainPlan, err := tpc.TrainPlanUseCase.GetPlanByID(c, request.Id)
	if err != nil {
		domain.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if request.IsCompleted() {
		trainPlan.Completion()
	} else {
		trainPlan.InCompletion()
	}

	if err = tpc.TrainPlanUseCase.UpdateByID(c, request.Id, &trainPlan); err != nil {
		domain.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	domain.SuccessResponse(c, nil)
}
