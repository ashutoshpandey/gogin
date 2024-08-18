package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/ashutoshpandey/gogin/config"
)

var publicUrls = []string{
	"login",
	"register",
}

// AuthMiddleware checks for a valid JWT in the Authorization header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ignore public urls that don't require jwt authentication
		requestPath := c.Request.URL.Path
		for _, url := range publicUrls {
			if strings.HasPrefix(requestPath, url) {
				// If the URL is public, skip authentication
				c.Next()
				return
			}
		}

		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")

		// Check if the header is present and starts with "Bearer "
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing or invalid"})
			c.Abort()
			return
		}

		// Extract the token from the header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		appConfig := config.LoadAppConfig()

		// JWTSecretKey is the secret key used for signing JWTs (should be same as in your LoginService)
		var JWTSecretKey = []byte(appConfig.JWT_SECRET_KEY)

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Make sure that the token's method is the expected signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
			}
			return JWTSecretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// If everything is fine, proceed to the next handler
		c.Next()
	}
}
