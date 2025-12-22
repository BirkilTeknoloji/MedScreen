package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
)

type tabletCihazService struct {
	repo repository.TabletCihazRepository
}

// NewTabletCihazService creates a new instance of TabletCihazService
func NewTabletCihazService(repo repository.TabletCihazRepository) TabletCihazService {
	return &tabletCihazService{repo: repo}
}

// GetByKodu retrieves a tablet device by its code
func (s *tabletCihazService) GetByKodu(kodu string) (*models.TabletCihaz, error) {
	if kodu == "" {
		return nil, errors.New("tablet_cihaz_kodu is required")
	}

	cihaz, err := s.repo.FindByKodu(kodu)
	if err != nil {
		return nil, err
	}
	if cihaz == nil {
		return nil, errors.New("tablet device not found")
	}

	return cihaz, nil
}

// GetByYatakKodu retrieves tablet devices by bed code
func (s *tabletCihazService) GetByYatakKodu(yatakKodu string, page, limit int) ([]models.TabletCihaz, int64, error) {
	if yatakKodu == "" {
		return nil, 0, errors.New("yatak_kodu is required")
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindByYatakKodu(yatakKodu, page, limit)
}

// GetAll retrieves all tablet devices with pagination
func (s *tabletCihazService) GetAll(page, limit int) ([]models.TabletCihaz, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindAll(page, limit)
}
