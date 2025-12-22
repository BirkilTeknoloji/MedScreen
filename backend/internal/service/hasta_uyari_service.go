package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
)

type hastaUyariService struct {
	repo repository.HastaUyariRepository
}

// NewHastaUyariService creates a new instance of HastaUyariService
func NewHastaUyariService(repo repository.HastaUyariRepository) HastaUyariService {
	return &hastaUyariService{repo: repo}
}

// GetByKodu retrieves a patient warning by its code
func (s *hastaUyariService) GetByKodu(kodu string) (*models.HastaUyari, error) {
	if kodu == "" {
		return nil, errors.New("hasta_uyari_kodu is required")
	}

	uyari, err := s.repo.FindByKodu(kodu)
	if err != nil {
		return nil, err
	}
	if uyari == nil {
		return nil, errors.New("patient warning not found")
	}

	return uyari, nil
}

// GetByBasvuruKodu retrieves patient warnings by patient visit code
func (s *hastaUyariService) GetByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.HastaUyari, int64, error) {
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

// GetByFilters retrieves patient warnings by various filters
func (s *hastaUyariService) GetByFilters(uyariTuru *string, aktiflik *int, page, limit int) ([]models.HastaUyari, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// If warning type filter is provided
	if uyariTuru != nil && *uyariTuru != "" {
		return s.repo.FindByTuru(*uyariTuru, page, limit)
	}

	// If active status filter is provided
	if aktiflik != nil {
		return s.repo.FindByAktiflik(*aktiflik, page, limit)
	}

	return nil, 0, errors.New("at least one filter (uyari_turu or aktiflik) is required")
}
