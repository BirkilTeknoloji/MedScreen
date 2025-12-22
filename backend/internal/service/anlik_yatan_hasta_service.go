package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
)

type anlikYatanHastaService struct {
	repo repository.AnlikYatanHastaRepository
}

// NewAnlikYatanHastaService creates a new instance of AnlikYatanHastaService
func NewAnlikYatanHastaService(repo repository.AnlikYatanHastaRepository) AnlikYatanHastaService {
	return &anlikYatanHastaService{repo: repo}
}

// GetByKodu retrieves a current inpatient by their code
func (s *anlikYatanHastaService) GetByKodu(kodu string) (*models.AnlikYatanHasta, error) {
	if kodu == "" {
		return nil, errors.New("anlik_yatan_hasta_kodu is required")
	}

	inpatient, err := s.repo.FindByKodu(kodu)
	if err != nil {
		return nil, err
	}
	if inpatient == nil {
		return nil, errors.New("current inpatient not found")
	}

	return inpatient, nil
}

// GetByYatakKodu retrieves current inpatients by bed code
func (s *anlikYatanHastaService) GetByYatakKodu(yatakKodu string, page, limit int) ([]models.AnlikYatanHasta, int64, error) {
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

// GetByHastaKodu retrieves current inpatients by patient code
func (s *anlikYatanHastaService) GetByHastaKodu(hastaKodu string, page, limit int) ([]models.AnlikYatanHasta, int64, error) {
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

// GetByBirimKodu retrieves current inpatients by unit code
func (s *anlikYatanHastaService) GetByBirimKodu(birimKodu string, page, limit int) ([]models.AnlikYatanHasta, int64, error) {
	if birimKodu == "" {
		return nil, 0, errors.New("birim_kodu is required")
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindByBirimKodu(birimKodu, page, limit)
}
