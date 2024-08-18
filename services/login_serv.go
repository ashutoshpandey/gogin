package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/ashutoshpandey/gogin/config"
	"github.com/ashutoshpandey/gogin/dtos"
	"github.com/ashutoshpandey/gogin/models"
)

// Initialization logic
// --------------------------------------------------------

// LoginService is the interface that defines the methods related to user management.
type LoginService struct {
	dbService *DbService
}

// NewLoginService creates a new LoginService with direct injection of DbService
func NewLoginService(dbService *DbService) *LoginService {
	return &LoginService{
		dbService: dbService,
	}
}

// Business methods
// --------------------------------------------------------

func (loginService *LoginService) DoLogin(loginDto dtos.LoginDTO) (string, error) {
	var user models.User

	fmt.Printf("\nEmail = %s", loginDto.Email)
	fmt.Printf("\nPassword = %s", loginDto.Password)

	// Find the user by email
	if err := loginService.dbService.DB.Where("LOWER(email) = ?", loginDto.Email).First(&user).Error; err != nil {
		fmt.Println(err)
		return "", errors.New("invalid email")
	}

	// Check the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
		return "", errors.New("invalid password")
	}

	// Generate a JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	})

	appConfig := config.LoadAppConfig()

	// JWTSecretKey is the secret key used for signing JWTs (should be same as in your LoginService)
	var JWTSecretKey = []byte(appConfig.JWT_SECRET_KEY)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(JWTSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
