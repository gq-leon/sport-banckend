package controller

import (
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

	userId, _ := primitive.ObjectIDFromHex(util.GetUserId(c))
	request.ID = primitive.NewObjectID()
	request.UserID = userId

	if err := ac.AttendanceUseCase.Create(c, &request); err != nil {
		domain.ErrorResponse(c, http.StatusInternalServerError, err)
		return
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
