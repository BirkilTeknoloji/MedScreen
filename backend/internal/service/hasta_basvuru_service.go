package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
	"time"
)

type hastaBasvuruService struct {
	repo repository.HastaBasvuruRepository
}

// NewHastaBasvuruService creates a new instance of HastaBasvuruService
func NewHastaBasvuruService(repo repository.HastaBasvuruRepository) HastaBasvuruService {
	return &hastaBasvuruService{repo: repo}
}

// GetByKodu retrieves a patient visit by its code
func (s *hastaBasvuruService) GetByKodu(kodu string) (*models.HastaBasvuru, error) {
	if kodu == "" {
		return nil, errors.New("hasta_basvuru_kodu is required")
	}

	basvuru, err := s.repo.FindByKodu(kodu)
	if err != nil {
		return nil, err
	}
	if basvuru == nil {
		return nil, errors.New("patient visit not found")
	}

	return basvuru, nil
}

// GetByHastaKodu retrieves patient visits by patient code
func (s *hastaBasvuruService) GetByHastaKodu(hastaKodu string, page, limit int) ([]models.HastaBasvuru, int64, error) {
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

// GetByHekimKodu retrieves patient visits by physician code
func (s *hastaBasvuruService) GetByHekimKodu(hekimKodu string, page, limit int) ([]models.HastaBasvuru, int64, error) {
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

// GetByFilters retrieves patient visits by various filters
func (s *hastaBasvuruService) GetByFilters(durum *string, startDate, endDate *time.Time, page, limit int) ([]models.HastaBasvuru, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// If status filter is provided
	if durum != nil && *durum != "" {
		return s.repo.FindByDurum(*durum, page, limit)
	}

	// If date range filter is provided
	if startDate != nil && endDate != nil {
		if startDate.After(*endDate) {
			return nil, 0, errors.New("start_date must be before end_date")
		}
		return s.repo.FindByDateRange(*startDate, *endDate, page, limit)
	}

	return nil, 0, errors.New("at least one filter (durum or date range) is required")
}
