package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
	"time"
)

type randevuService struct {
	repo repository.RandevuRepository
}

// NewRandevuService creates a new instance of RandevuService
func NewRandevuService(repo repository.RandevuRepository) RandevuService {
	return &randevuService{repo: repo}
}

// GetByKodu retrieves an appointment by its code
func (s *randevuService) GetByKodu(kodu string) (*models.Randevu, error) {
	if kodu == "" {
		return nil, errors.New("randevu_kodu is required")
	}

	randevu, err := s.repo.FindByKodu(kodu)
	if err != nil {
		return nil, err
	}
	if randevu == nil {
		return nil, errors.New("appointment not found")
	}

	return randevu, nil
}

// GetByHastaKodu retrieves appointments by patient code
func (s *randevuService) GetByHastaKodu(hastaKodu string, page, limit int) ([]models.Randevu, int64, error) {
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

// GetByBasvuruKodu retrieves appointments by visit code
func (s *randevuService) GetByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.Randevu, int64, error) {
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

// GetByHekimKodu retrieves appointments by physician code
func (s *randevuService) GetByHekimKodu(hekimKodu string, page, limit int) ([]models.Randevu, int64, error) {
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

// GetByTuru retrieves appointments by type
func (s *randevuService) GetByTuru(randevuTuru string, page, limit int) ([]models.Randevu, int64, error) {
	if randevuTuru == "" {
		return nil, 0, errors.New("randevu_turu is required")
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindByTuru(randevuTuru, page, limit)
}

// GetByDateRange retrieves appointments within a date range
func (s *randevuService) GetByDateRange(startDate, endDate time.Time, page, limit int) ([]models.Randevu, int64, error) {
	if startDate.After(endDate) {
		return nil, 0, errors.New("start_date must be before end_date")
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindByDateRange(startDate, endDate, page, limit)
}
