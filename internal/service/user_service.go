package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
)

type userService struct {
	repo repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// CreateUser creates a new user with validation
func (s *userService) CreateUser(user *models.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	// Validate required fields
	if user.FirstName == "" {
		return errors.New("first_name is required")
	}
	if user.LastName == "" {
		return errors.New("last_name is required")
	}
	if user.Phone == "" {
		return errors.New("phone is required")
	}

	// Validate role enum
	if err := validateUserRole(user.Role); err != nil {
		return err
	}

	return s.repo.Create(user)
}

// GetUser retrieves a user by ID
func (s *userService) GetUser(id uint) (*models.User, error) {
	if id == 0 {
		return nil, errors.New("invalid user id")
	}

	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// GetUsers retrieves all users with optional role filter and pagination
func (s *userService) GetUsers(page, limit int, role *models.UserRole) ([]models.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	if role != nil {
		// Validate role enum if provided
		if err := validateUserRole(*role); err != nil {
			return nil, 0, err
		}
		return s.repo.FindByRole(*role, page, limit)
	}

	return s.repo.FindAll(page, limit)
}

// UpdateUser updates an existing user with validation
func (s *userService) UpdateUser(id uint, user *models.User) error {
	if id == 0 {
		return errors.New("invalid user id")
	}
	if user == nil {
		return errors.New("user cannot be nil")
	}

	// Check if user exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("user not found")
	}

	// Validate role enum if being updated
	if user.Role != "" {
		if err := validateUserRole(user.Role); err != nil {
			return err
		}
	}

	// Set the ID to ensure we're updating the correct record
	user.ID = id

	return s.repo.Update(user)
}

// DeleteUser soft deletes a user
func (s *userService) DeleteUser(id uint) error {
	if id == 0 {
		return errors.New("invalid user id")
	}

	// Check if user exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("user not found")
	}

	return s.repo.Delete(id)
}

// AuthenticateByNFC authenticates a user by NFC card ID
func (s *userService) AuthenticateByNFC(nfcCardID string) (*models.User, error) {
	if nfcCardID == "" {
		return nil, errors.New("nfc_card_id is required")
	}

	user, err := s.repo.FindByNFCCardID(nfcCardID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found with provided nfc_card_id")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, errors.New("user account is inactive")
	}

	return user, nil
}

// validateUserRole validates that the role is one of the allowed values
func validateUserRole(role models.UserRole) error {
	switch role {
	case models.RoleDoctor, models.RoleNurse, models.RoleReceptionist, models.RoleAdmin:
		return nil
	default:
		return errors.New("invalid role: must be one of doctor, nurse, receptionist, admin")
	}
}
