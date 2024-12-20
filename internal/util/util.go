package util

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUserId(c *gin.Context) string {
	return c.GetString("x-user-id")
}

func GenerateCheckInKey(userID string, date ...time.Time) string {
	now := time.Now()
	if len(date) != 0 {
		now = date[0]
	}

	return fmt.Sprintf("checkin_%d-%s", now.Year(), userID)
}
