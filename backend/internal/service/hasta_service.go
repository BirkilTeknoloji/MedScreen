package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
)

type hastaService struct {
	repo repository.HastaRepository
}

// NewHastaService creates a new instance of HastaService
func NewHastaService(repo repository.HastaRepository) HastaService {
	return &hastaService{repo: repo}
}

// GetByKodu retrieves a patient by their code
func (s *hastaService) GetByKodu(kodu string) (*models.Hasta, error) {
	if kodu == "" {
		return nil, errors.New("hasta_kodu is required")
	}

	hasta, err := s.repo.FindByKodu(kodu)
	if err != nil {
		return nil, err
	}
	if hasta == nil {
		return nil, errors.New("patient not found")
	}

	return hasta, nil
}

// GetByTCKimlik retrieves a patient by their Turkish ID number
func (s *hastaService) GetByTCKimlik(tcKimlik string) (*models.Hasta, error) {
	if tcKimlik == "" {
		return nil, errors.New("tc_kimlik_numarasi is required")
	}

	// Validate TC number format (11 digits)
	if err := validateTCKimlik(tcKimlik); err != nil {
		return nil, err
	}

	hasta, err := s.repo.FindByTCKimlik(tcKimlik)
	if err != nil {
		return nil, err
	}
	if hasta == nil {
		return nil, errors.New("patient not found")
	}

	return hasta, nil
}

// GetAll retrieves all patients with pagination
func (s *hastaService) GetAll(page, limit int) ([]models.Hasta, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindAll(page, limit)
}

// SearchByName searches for patients by name
func (s *hastaService) SearchByName(name string, page, limit int) ([]models.Hasta, int64, error) {
	if name == "" {
		return nil, 0, errors.New("name is required for search")
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.SearchByName(name, page, limit)
}

// validateTCKimlik validates that the TC number is exactly 11 digits
func validateTCKimlik(tcKimlik string) error {
	if len(tcKimlik) != 11 {
		return errors.New("tc_kimlik_numarasi must be exactly 11 digits")
	}

	for _, char := range tcKimlik {
		if char < '0' || char > '9' {
			return errors.New("tc_kimlik_numarasi must contain only digits")
		}
	}

	return nil
}
