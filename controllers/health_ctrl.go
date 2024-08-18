package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ashutoshpandey/gogin/services"
)

func RegisterHealthRoutes(router *gin.Engine) {
	healthService := services.NewHealthService()

	healthRoutes := router.Group("/health")
	{
		healthRoutes.GET("/", func(c *gin.Context) {
			ReturnHealth(c, healthService)
		})
	}
}

func ReturnHealth(c *gin.Context, healthService services.HealthService) {
	message := healthService.GetServerHealth()
	c.JSON(http.StatusOK, gin.H{"message": message})
}
