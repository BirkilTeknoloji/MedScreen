package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
)

type hastaTibbiBilgiService struct {
	repo repository.HastaTibbiBilgiRepository
}

// NewHastaTibbiBilgiService creates a new instance of HastaTibbiBilgiService
func NewHastaTibbiBilgiService(repo repository.HastaTibbiBilgiRepository) HastaTibbiBilgiService {
	return &hastaTibbiBilgiService{repo: repo}
}

// GetByKodu retrieves patient medical information by its code
func (s *hastaTibbiBilgiService) GetByKodu(kodu string) (*models.HastaTibbiBilgi, error) {
	if kodu == "" {
		return nil, errors.New("hasta_tibbi_bilgi_kodu is required")
	}

	bilgi, err := s.repo.FindByKodu(kodu)
	if err != nil {
		return nil, err
	}
	if bilgi == nil {
		return nil, errors.New("patient medical information not found")
	}

	return bilgi, nil
}

// GetByHastaKodu retrieves patient medical information by patient code
func (s *hastaTibbiBilgiService) GetByHastaKodu(hastaKodu string, page, limit int) ([]models.HastaTibbiBilgi, int64, error) {
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

// GetByTuru retrieves patient medical information by type code
func (s *hastaTibbiBilgiService) GetByTuru(turuKodu string, page, limit int) ([]models.HastaTibbiBilgi, int64, error) {
	if turuKodu == "" {
		return nil, 0, errors.New("tibbi_bilgi_turu_kodu is required")
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindByTuru(turuKodu, page, limit)
}
