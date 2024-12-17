package util

import "github.com/gin-gonic/gin"

func GetUserId(c *gin.Context) string {
	return c.GetString("x-user-id")
}
