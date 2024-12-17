package domain

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, err error) {
	c.JSON(http.StatusBadRequest, Response{
		Code:    code,
		Message: err.Error(),
		Data:    nil,
	})
}
