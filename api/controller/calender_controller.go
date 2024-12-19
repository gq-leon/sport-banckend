package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gq-leon/sport-backend/domain"
	"github.com/gq-leon/sport-backend/internal/util"
)

type CalenderController struct {
	CalenderUseCase domain.CalenderUseCase
}

func (cc *CalenderController) Workouts(c *gin.Context) {
	var request domain.MonthWorkoutsRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		domain.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	userId := util.GetUserId(c)
	workouts, err := cc.CalenderUseCase.GetDayWorkouts(c, userId, request.Date)
	if err != nil {
		domain.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	var (
		result = domain.WorkoutsResponse{
			Date: request.Date,
		}
	)

	for _, workout := range workouts {
		result.Exercises = append(result.Exercises, domain.Exercise{
			ID:        workout.GetID(),
			Name:      workout.Name,
			Category:  workout.Category,
			Reps:      workout.Reps,
			Weight:    workout.Weight,
			Distance:  workout.Distance,
			Duration:  workout.Duration,
			Sets:      workout.Sets,
			Completed: workout.Completed,
			Type:      workout.GetType(),
		})
	}

	domain.SuccessResponse(c, result)
}

func (cc *CalenderController) MonthWorkouts(c *gin.Context) {
	var request domain.MonthWorkoutsRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		domain.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	userId := util.GetUserId(c)
	workouts, err := cc.CalenderUseCase.GetMonthWorkouts(c, userId, request.Date)
	if err != nil {
		domain.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	var (
		data   = make(map[string][]domain.Exercise)
		result []domain.WorkoutsResponse
	)

	for _, workout := range workouts {
		exercises := data[workout.Date]
		data[workout.Date] = append(exercises, domain.Exercise{
			ID:        workout.GetID(),
			Name:      workout.Name,
			Category:  workout.Category,
			Reps:      workout.Reps,
			Weight:    workout.Weight,
			Distance:  workout.Distance,
			Duration:  workout.Duration,
			Sets:      workout.Sets,
			Completed: workout.Completed,
			Type:      workout.GetType(),
		})
	}

	for date, datum := range data {
		result = append(result, domain.WorkoutsResponse{
			Date:      date,
			Exercises: datum,
		})
	}

	domain.SuccessResponse(c, result)
}
