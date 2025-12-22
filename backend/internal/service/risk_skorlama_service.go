package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
)

type riskSkorlamaService struct {
	repo repository.RiskSkorlamaRepository
}

// NewRiskSkorlamaService creates a new instance of RiskSkorlamaService
func NewRiskSkorlamaService(repo repository.RiskSkorlamaRepository) RiskSkorlamaService {
	return &riskSkorlamaService{repo: repo}
}

// GetByKodu retrieves a risk score by its code
func (s *riskSkorlamaService) GetByKodu(kodu string) (*models.RiskSkorlama, error) {
	if kodu == "" {
		return nil, errors.New("risk_skorlama_kodu is required")
	}

	skor, err := s.repo.FindByKodu(kodu)
	if err != nil {
		return nil, err
	}
	if skor == nil {
		return nil, errors.New("risk score not found")
	}

	return skor, nil
}

// GetByBasvuruKodu retrieves risk scores by patient visit code
func (s *riskSkorlamaService) GetByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.RiskSkorlama, int64, error) {
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

// GetByTuru retrieves risk scores by score type
func (s *riskSkorlamaService) GetByTuru(turu string, page, limit int) ([]models.RiskSkorlama, int64, error) {
	if turu == "" {
		return nil, 0, errors.New("risk_skorlama_turu is required")
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindByTuru(turu, page, limit)
}
