package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ashutoshpandey/gogin/services"
)

func ReturnHealth(c *gin.Context, healthService services.HealthService) {
	message := healthService.GetServerHealth()
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func RegisterHealthRoutes(router *gin.Engine) {
	healthService := services.NewUserService()

	healthRoutes := router.Group("/health")
	{
		healthRoutes.GET("/", func(c *gin.Context) {
			ReturnHealth(c, healthService)
		})
	}
}
