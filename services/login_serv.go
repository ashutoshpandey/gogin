package services

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/ashutoshpandey/gogin/dtos"
	"github.com/ashutoshpandey/gogin/models"
)

// Initialization logic
// --------------------------------------------------------

var JWTSecretKey = []byte("your-secret-key")

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

	// Find the user by email
	if err := loginService.dbService.DB.Where("email = ?", loginDto.Email).First(&user).Error; err != nil {
		return "", errors.New("invalid email or password")
	}

	// Check the password
	if err := bcrypt.CompareHashAndPassword([]byte(loginDto.Password), []byte(user.Password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate a JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(JWTSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
