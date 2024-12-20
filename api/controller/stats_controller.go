package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gq-leon/sport-backend/domain"
	"github.com/gq-leon/sport-backend/internal/util"
)

type StatsController struct {
	StatsUseCase domain.StatsUseCase
}

func (sc *StatsController) Profile(c *gin.Context) {
	var (
		userId = util.GetUserId(c)
	)

	stats, err := sc.StatsUseCase.GetProfileStats(c, userId)
	if err != nil {
		domain.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	domain.SuccessResponse(c, stats)
}
