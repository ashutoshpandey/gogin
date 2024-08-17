package services

// UserService is the interface that defines the methods related to user management.
type HealthService interface {
	GetServerHealth() string
}

// userServiceImpl is a concrete implementation of the UserService interface.
type healthServiceImpl struct{}

// NewUserService creates and returns a new UserService instance.
func NewUserService() HealthService {
	return &healthServiceImpl{}
}

// GetAllUsers returns a list of users (simulated business logic).
func (s *healthServiceImpl) GetServerHealth() string {
	// Simulated business logic for retrieving users
	return "All good, server is running fine"
}
