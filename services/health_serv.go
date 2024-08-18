package services

// Initialization logic
// --------------------------------------------------------

// HealthService is the interface that defines the methods related to user management.
type HealthService interface {
	GetServerHealth() string
}

// healthServiceImpl is a concrete implementation of the UserService interface.
type healthServiceImpl struct{}

// NewHealthService creates and returns a new HealthServiceImpl instance.
func NewHealthService() HealthService {
	return &healthServiceImpl{}
}

// Business methods
// --------------------------------------------------------

func (s *healthServiceImpl) GetServerHealth() string {
	return "All good, server is running fine"
}
