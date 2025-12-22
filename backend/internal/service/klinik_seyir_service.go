package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
	"time"
)

type klinikSeyirService struct {
	repo repository.KlinikSeyirRepository
}

// NewKlinikSeyirService creates a new instance of KlinikSeyirService
func NewKlinikSeyirService(repo repository.KlinikSeyirRepository) KlinikSeyirService {
	return &klinikSeyirService{repo: repo}
}

// GetByKodu retrieves clinical progress notes by their code
func (s *klinikSeyirService) GetByKodu(kodu string) (*models.KlinikSeyir, error) {
	if kodu == "" {
		return nil, errors.New("klinik_seyir_kodu is required")
	}

	seyir, err := s.repo.FindByKodu(kodu)
	if err != nil {
		return nil, err
	}
	if seyir == nil {
		return nil, errors.New("clinical progress note not found")
	}

	return seyir, nil
}

// GetByBasvuruKodu retrieves clinical progress notes by patient visit code
func (s *klinikSeyirService) GetByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.KlinikSeyir, int64, error) {
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

// GetByFilters retrieves clinical progress notes by various filters
func (s *klinikSeyirService) GetByFilters(seyirTipi *string, startDate, endDate *time.Time, page, limit int) ([]models.KlinikSeyir, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// If note type filter is provided
	if seyirTipi != nil && *seyirTipi != "" {
		return s.repo.FindBySeyirTipi(*seyirTipi, page, limit)
	}

	// If date range filter is provided
	if startDate != nil && endDate != nil {
		if startDate.After(*endDate) {
			return nil, 0, errors.New("start_date must be before end_date")
		}
		return s.repo.FindByDateRange(*startDate, *endDate, page, limit)
	}

	return nil, 0, errors.New("at least one filter (seyir_tipi or date range) is required")
}
