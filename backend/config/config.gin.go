package config

import "github.com/gin-gonic/gin"

func Response(ctx *gin.Context, httpCode int, data interface{}) {
	ctx.JSON(httpCode, data)
}

func ResponseWithMessage(ctx *gin.Context, httpCode int, message string) {
	ctx.JSON(httpCode, gin.H{
		"message": message,
	})
}

func ResponseWithMessageAndData(ctx *gin.Context, httpCode int, message string, data interface{}) {
	ctx.JSON(httpCode, gin.H{
		"message": message,
		"data":    data,
	})
}

func ResponseWithError(ctx *gin.Context, httpCode int, err error) {
	ctx.JSON(httpCode, gin.H{
		"error": err.Error(),
	})
}

func ResponseWithErrorMessage(ctx *gin.Context, httpCode int, message string, err error) {
	ctx.JSON(httpCode, gin.H{
		"message": message,
		"error":   err.Error(),
	})
}

func ResponseWithErrorMessageAndData(ctx *gin.Context, httpCode int, message string, err error, data interface{}) {
	ctx.JSON(httpCode, gin.H{
		"message": message,
		"error":   err.Error(),
		"data":    data,
	})
}
