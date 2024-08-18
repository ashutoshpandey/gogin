package services

import (
	"time"

	"github.com/ulule/deepcopier"

	"github.com/ashutoshpandey/gogin/dtos"
	"github.com/ashutoshpandey/gogin/models"
)

// Initialization logic
// --------------------------------------------------------

// UserService provides user-related operations
type UserService struct {
	dbService *DbService
}

// NewUserService creates a new UserService
func NewUserService(dbService *DbService) *UserService {
	return &UserService{
		dbService: dbService,
	}
}

// Business methods
// --------------------------------------------------------

// CreateUser maps the DTO to the User model and saves it in the database
func (userService *UserService) CreateUser(userDto dtos.CreateUserDTO) (*models.User, error) {
	var user models.User
	deepcopier.Copy(userDto).To(&user)

	// Convert the DateOfBirth string to time.Time
	dateOfBirth, _ := time.Parse("2006-01-02", userDto.DateOfBirth)
	user.DateOfBirth = dateOfBirth

	err := userService.dbService.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUsers returns a list of users with pagination
func (userService *UserService) GetUsers(pageNumber, pageSize int) ([]models.User, int64, error) {
	var total int64
	var users []models.User

	// Calculate the offset for pagination
	offset := (pageNumber - 1) * pageSize

	// Query to get the total count of users
	err := userService.dbService.DB.Model(&models.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Query the database with limit and offset
	err = userService.dbService.DB.Limit(pageSize).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
