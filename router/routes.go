package router

import (
	"github.com/fogtz/goVacations.git/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowVacancyHandler)
		v1.POST("/create", handler.CreateVacancyHandler)
		v1.DELETE("/delete", handler.DeleteVacancyHandler)
		v1.PUT("/update", handler.UpdateVacancyHandler)
		v1.GET("/list", handler.ListVacancyHandler)
	}
}
