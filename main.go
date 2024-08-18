package main

import (
	"fmt"

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
	startServer(cfg, r)
}

func startServer(cfg *config.ServerConfig, r *gin.Engine) {
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
