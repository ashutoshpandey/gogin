package dtos

// CreateUserDTO represents the data structure for creating a new user
type CreateUserDTO struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required,datetime=2000-01-22"`
	// datetime=2000-01-22 is reference for datetime format, yyyy-mm-dd
}
