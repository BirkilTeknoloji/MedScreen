package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
)

type basvuruYemekService struct {
	repo repository.BasvuruYemekRepository
}

// NewBasvuruYemekService creates a new instance of BasvuruYemekService
func NewBasvuruYemekService(repo repository.BasvuruYemekRepository) BasvuruYemekService {
	return &basvuruYemekService{repo: repo}
}

// GetByKodu retrieves meal information by its code
func (s *basvuruYemekService) GetByKodu(kodu string) (*models.BasvuruYemek, error) {
	if kodu == "" {
		return nil, errors.New("basvuru_yemek_kodu is required")
	}

	yemek, err := s.repo.FindByKodu(kodu)
	if err != nil {
		return nil, err
	}
	if yemek == nil {
		return nil, errors.New("meal information not found")
	}

	return yemek, nil
}

// GetByBasvuruKodu retrieves meal information by patient visit code
func (s *basvuruYemekService) GetByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.BasvuruYemek, int64, error) {
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

// GetByTuru retrieves meal information by meal type
func (s *basvuruYemekService) GetByTuru(yemekTuru string, page, limit int) ([]models.BasvuruYemek, int64, error) {
	if yemekTuru == "" {
		return nil, 0, errors.New("yemek_turu is required")
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindByTuru(yemekTuru, page, limit)
}
