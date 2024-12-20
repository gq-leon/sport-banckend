package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gq-leon/sport-backend/domain"
	"github.com/gq-leon/sport-backend/internal/util"
)

type AttendanceController struct {
	AttendanceUseCase domain.AttendanceUseCase
}

func (ac *AttendanceController) Create(c *gin.Context) {
	var request domain.Attendance

	if err := c.ShouldBindJSON(&request); err != nil {
		domain.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	userId := util.GetUserId(c)
	userID, _ := primitive.ObjectIDFromHex(userId)
	request.ID = primitive.NewObjectID()
	request.UserID = userID

	if err := ac.AttendanceUseCase.Create(c, &request); err != nil {
		domain.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if request.Location == domain.AttendanceGym {
		_ = ac.AttendanceUseCase.CheckIn(c, userId)
	}

	domain.SuccessResponse(c, nil)
}

func (ac *AttendanceController) List(c *gin.Context) {
	userId := util.GetUserId(c)
	attendances, err := ac.AttendanceUseCase.Fetch(c, userId)
	if err != nil {
		domain.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	domain.SuccessResponse(c, attendances)
}

func (ac *AttendanceController) BackDateCheckIn(c *gin.Context) {
	var request domain.CheckInRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		domain.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if len(request.Date) < 1 {
		domain.ErrorResponse(c, http.StatusBadRequest, errors.New("date is required"))
		return
	}

	userId := util.GetUserId(c)
	if err := ac.AttendanceUseCase.BackDateCheckIn(c, userId, request.Date); err != nil {
		domain.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	domain.SuccessResponse(c, nil)
}
