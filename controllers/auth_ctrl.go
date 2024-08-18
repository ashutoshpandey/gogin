package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ashutoshpandey/gogin/dtos"
	"github.com/ashutoshpandey/gogin/services"
)

func RegisterAuthRoutes(router *gin.Engine) {
	dbService := services.NewDBService()
	loginService := services.NewLoginService(dbService)

	loginRoutes := router.Group("/auth")
	{
		loginRoutes.POST("/login", func(c *gin.Context) {
			DoLogin(c, loginService)
		})
	}
}

func DoLogin(c *gin.Context, loginService *services.LoginService) {
	var loginDTO dtos.LoginDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Authenticate the user
	token, err := loginService.DoLogin(loginDTO)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Return the JWT
	c.JSON(http.StatusOK, gin.H{"token": token})
}
