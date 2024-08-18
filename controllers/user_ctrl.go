package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ashutoshpandey/gogin/dtos"
	"github.com/ashutoshpandey/gogin/services"
	"github.com/ashutoshpandey/gogin/utils"
)

func RegisterUserRoutes(router *gin.Engine) {
	dbService := services.NewDBService()
	userService := services.NewUserService(dbService)

	healthRoutes := router.Group("/users")
	{
		healthRoutes.GET("/", func(c *gin.Context) {
			GetUsers(c, userService)
		})
		healthRoutes.POST("/", func(c *gin.Context) {
			CreateUser(c, userService)
		})
	}
}

func CreateUser(c *gin.Context, userService *services.UserService) {
	var userDTO dtos.CreateUserDTO
	user, err := userService.CreateUser(userDTO)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context, userService *services.UserService) {
	// Define the keys you want to extract from the query string
	keys := []string{"pageNumber", "pageSize"}

	// Use the utility function to parse the query parameters
	queryParams, err := utils.ParseQueryParams(c, keys)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract pageNumber and pageSize from the returned map
	pageSize := queryParams["pageSize"]
	pageNumber := queryParams["pageNumber"]

	users, total, err := userService.GetUsers(pageNumber, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"total": total, "users": users}})
}
