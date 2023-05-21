package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": true,
			})
		})
		v1.GET("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "GET OPENING",
			})
		})
		v1.POST("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "CREATE OPENING",
			})
		})
		v1.DELETE("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "DELETE OPENING",
			})
		})
		v1.PUT("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "PUT OPENING",
			})
		})
		v1.GET("/openings", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "GET OPENINGS",
			})
		})
	}
}
