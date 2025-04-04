package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET opening",
			})
		})
		v1.POST("/create", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "POST opening",
			})
		})
		v1.DELETE("/delete", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "DELETE opening",
			})
		})
		v1.PUT("/update", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "PUT opening",
			})
		})
		v1.GET("/list", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET openings",
			})
		})
	}
}
