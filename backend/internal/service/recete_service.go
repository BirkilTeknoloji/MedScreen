package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
)

type receteService struct {
	repo repository.ReceteRepository
}

// NewReceteService creates a new instance of ReceteService
func NewReceteService(repo repository.ReceteRepository) ReceteService {
	return &receteService{repo: repo}
}

// GetByKodu retrieves a prescription by its code
func (s *receteService) GetByKodu(kodu string) (*models.Recete, error) {
	if kodu == "" {
		return nil, errors.New("recete_kodu is required")
	}

	recete, err := s.repo.FindByKodu(kodu)
	if err != nil {
		return nil, err
	}
	if recete == nil {
		return nil, errors.New("prescription not found")
	}

	return recete, nil
}

// GetByBasvuruKodu retrieves prescriptions by patient visit code
func (s *receteService) GetByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.Recete, int64, error) {
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

// GetByHekimKodu retrieves prescriptions by physician code
func (s *receteService) GetByHekimKodu(hekimKodu string, page, limit int) ([]models.Recete, int64, error) {
	if hekimKodu == "" {
		return nil, 0, errors.New("hekim_kodu is required")
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindByHekimKodu(hekimKodu, page, limit)
}

// GetIlaclar retrieves prescription medications by prescription code
func (s *receteService) GetIlaclar(receteKodu string, page, limit int) ([]models.ReceteIlac, int64, error) {
	if receteKodu == "" {
		return nil, 0, errors.New("recete_kodu is required")
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindIlacByReceteKodu(receteKodu, page, limit)
}
