package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
	"time"
)

type hastaVitalFizikiBulguService struct {
	repo repository.HastaVitalFizikiBulguRepository
}

// NewHastaVitalFizikiBulguService creates a new instance of HastaVitalFizikiBulguService
func NewHastaVitalFizikiBulguService(repo repository.HastaVitalFizikiBulguRepository) HastaVitalFizikiBulguService {
	return &hastaVitalFizikiBulguService{repo: repo}
}

// GetByKodu retrieves vital signs by their code
func (s *hastaVitalFizikiBulguService) GetByKodu(kodu string) (*models.HastaVitalFizikiBulgu, error) {
	if kodu == "" {
		return nil, errors.New("hasta_vital_fiziki_bulgu_kodu is required")
	}

	bulgu, err := s.repo.FindByKodu(kodu)
	if err != nil {
		return nil, err
	}
	if bulgu == nil {
		return nil, errors.New("vital signs not found")
	}

	return bulgu, nil
}

// GetByBasvuruKodu retrieves vital signs by patient visit code
func (s *hastaVitalFizikiBulguService) GetByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.HastaVitalFizikiBulgu, int64, error) {
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

// GetByDateRange retrieves vital signs within a date range
func (s *hastaVitalFizikiBulguService) GetByDateRange(startDate, endDate time.Time, page, limit int) ([]models.HastaVitalFizikiBulgu, int64, error) {
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
