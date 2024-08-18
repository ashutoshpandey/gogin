package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ashutoshpandey/gogin/config"
	"github.com/ashutoshpandey/gogin/controllers"
)

func main() {
	cfg := config.LoadServerConfig()

	r := gin.Default()

	// Setup all controller routes
	registerRoutes(r)

	// Start the server using the loaded configuration
	setupServer(cfg, r)
}

func setupServer(cfg *config.ServerConfig, r *gin.Engine) {
	// Apply CORS middleware globally
	r.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(cfg.ALLOWED_ORIGINS, ","),             // Allow specific domains
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allow specific HTTP methods
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"}, // Allow specific headers
		ExposeHeaders:    []string{"Content-Length"},                          // Expose specific headers
		AllowCredentials: true,                                                // Allow credentials (cookies, authorization headers)
		MaxAge:           12 * time.Hour,                                      // Preflight request cache duration
	}))

	port := cfg.PORT
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting server on port %s\n", port)
	r.Run(":" + port)
}

// Initialize all controllers
func registerRoutes(r *gin.Engine) {
	controllers.RegisterUserRoutes(r)
	controllers.RegisterHealthRoutes(r)
}
