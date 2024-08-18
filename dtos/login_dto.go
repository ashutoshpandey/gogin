package dtos

// LoginDTO represents the data structure for a login request
type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
