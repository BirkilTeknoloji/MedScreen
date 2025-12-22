package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
)

type yatakService struct {
	repo repository.YatakRepository
}

// NewYatakService creates a new instance of YatakService
func NewYatakService(repo repository.YatakRepository) YatakService {
	return &yatakService{repo: repo}
}

// GetByKodu retrieves a bed by its code
func (s *yatakService) GetByKodu(kodu string) (*models.Yatak, error) {
	if kodu == "" {
		return nil, errors.New("yatak_kodu is required")
	}

	yatak, err := s.repo.FindByKodu(kodu)
	if err != nil {
		return nil, err
	}
	if yatak == nil {
		return nil, errors.New("bed not found")
	}

	return yatak, nil
}

// GetByBirimAndOda retrieves beds by unit and room codes
func (s *yatakService) GetByBirimAndOda(birimKodu, odaKodu string, page, limit int) ([]models.Yatak, int64, error) {
	if birimKodu == "" {
		return nil, 0, errors.New("birim_kodu is required")
	}
	if odaKodu == "" {
		return nil, 0, errors.New("oda_kodu is required")
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindByBirimAndOda(birimKodu, odaKodu, page, limit)
}

// GetAll retrieves all beds with pagination
func (s *yatakService) GetAll(page, limit int) ([]models.Yatak, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindAll(page, limit)
}
