package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
)

type basvuruTaniService struct {
	repo repository.BasvuruTaniRepository
}

// NewBasvuruTaniService creates a new instance of BasvuruTaniService
func NewBasvuruTaniService(repo repository.BasvuruTaniRepository) BasvuruTaniService {
	return &basvuruTaniService{repo: repo}
}

// GetByKodu retrieves a diagnosis by its code
func (s *basvuruTaniService) GetByKodu(kodu string) (*models.BasvuruTani, error) {
	if kodu == "" {
		return nil, errors.New("basvuru_tani_kodu is required")
	}

	tani, err := s.repo.FindByKodu(kodu)
	if err != nil {
		return nil, err
	}
	if tani == nil {
		return nil, errors.New("diagnosis not found")
	}

	return tani, nil
}

// GetByHastaKodu retrieves diagnoses by patient code
func (s *basvuruTaniService) GetByHastaKodu(hastaKodu string, page, limit int) ([]models.BasvuruTani, int64, error) {
	if hastaKodu == "" {
		return nil, 0, errors.New("hasta_kodu is required")
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindByHastaKodu(hastaKodu, page, limit)
}

// GetByBasvuruKodu retrieves diagnoses by patient visit code
func (s *basvuruTaniService) GetByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.BasvuruTani, int64, error) {
	if basvuruKodu == "" {
		return nil, 0, errors.New("hasta_basvuru_kodu is required")
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindByBasvuruKodu(basvuruKodu, page, limit)
}
