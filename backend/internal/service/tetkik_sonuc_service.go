package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
)

type tetkikSonucService struct {
	repo repository.TetkikSonucRepository
}

// NewTetkikSonucService creates a new instance of TetkikSonucService
func NewTetkikSonucService(repo repository.TetkikSonucRepository) TetkikSonucService {
	return &tetkikSonucService{repo: repo}
}

// GetByKodu retrieves test results by their code
func (s *tetkikSonucService) GetByKodu(kodu string) (*models.TetkikSonuc, error) {
	if kodu == "" {
		return nil, errors.New("tetkik_sonuc_kodu is required")
	}

	sonuc, err := s.repo.FindByKodu(kodu)
	if err != nil {
		return nil, err
	}
	if sonuc == nil {
		return nil, errors.New("test result not found")
	}

	return sonuc, nil
}

// GetByBasvuruKodu retrieves test results by patient visit code
func (s *tetkikSonucService) GetByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.TetkikSonuc, int64, error) {
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
