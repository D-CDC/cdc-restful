package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonResponse(c *gin.Context, httpCode, code int, msg string, data interface{}) {
	c.JSON(httpCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func InternalError(c *gin.Context, code int, msg string) {
	JsonResponse(c, http.StatusInternalServerError, code, msg, make(map[string]string))
}

func SuccessResponse(c *gin.Context, code int, msg string, data interface{}) {
	JsonResponse(c, http.StatusOK, code, msg, data)
}
