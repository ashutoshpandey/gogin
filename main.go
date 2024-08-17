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

	// Start the server using the loaded configuration
	registerRoutes(r)
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

func registerRoutes(r *gin.Engine) {
	controllers.RegisterHealthRoutes(r)
}
