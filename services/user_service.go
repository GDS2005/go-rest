package services

import (
	"example/models"
	"example/repositories"
)

// UserService contains business logic for user operations
type UserService struct {
	userRepository *repositories.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

// GetAllUsers retrieves all users from the repository
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepository.GetAll()
}

// GetUserByID retrieves a user by ID from the repository
func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.userRepository.GetByID(id)
}

// CreateUser creates a new user in the repository
func (s *UserService) CreateUser(user *models.User) error {
	// You can add validation logic here if needed
	return s.userRepository.Create(user)
}

// UpdateUser updates an existing user in the repository
func (s *UserService) UpdateUser(id int, user *models.User) error {
	// You can add validation logic here if needed
	return s.userRepository.Update(id, user)
}

// DeleteUser deletes a user from the repository
func (s *UserService) DeleteUser(id int) error {
	return s.userRepository.Delete(id)
}
