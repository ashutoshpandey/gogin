package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/ashutoshpandey/gogin/config"
	"github.com/ashutoshpandey/gogin/controllers"
)

func main() {
	cfg := config.LoadConfig()

	r := gin.Default()

	controllers.RegisterHealthRoutes(r)

	// Start the server using the loaded configuration
	port := cfg.PORT
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting server on port %s\n", port)
	r.Run(":" + port)
}
