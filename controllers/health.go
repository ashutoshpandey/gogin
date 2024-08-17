package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Server is running fine"})
}

func RegisterHealthRoutes(router *gin.Engine) {
	userRoutes := router.Group("/health")
	{
		userRoutes.GET("/", ReturnHealth)
	}
}
