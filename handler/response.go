package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(code, gin.H{
		"message":  message,
		"erroCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(200, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfully", op),
		"data":    data,
	})
}
