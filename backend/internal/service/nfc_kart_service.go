package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
)

type nfcKartService struct {
	repo repository.NFCKartRepository
}

// NewNFCKartService creates a new instance of NFCKartService
func NewNFCKartService(repo repository.NFCKartRepository) NFCKartService {
	return &nfcKartService{repo: repo}
}

// GetByKodu retrieves an NFC card by its code
func (s *nfcKartService) GetByKodu(kodu string) (*models.NFCKart, error) {
	if kodu == "" {
		return nil, errors.New("nfc_kart_kodu is required")
	}

	nfcKart, err := s.repo.FindByKodu(kodu)
	if err != nil {
		return nil, err
	}
	if nfcKart == nil {
		return nil, errors.New("NFC card not found")
	}

	return nfcKart, nil
}

// GetByKartUID retrieves an NFC card by its card UID
func (s *nfcKartService) GetByKartUID(kartUID string) (*models.NFCKart, error) {
	if kartUID == "" {
		return nil, errors.New("kart_uid is required")
	}

	nfcKart, err := s.repo.FindByKartUID(kartUID)
	if err != nil {
		return nil, err
	}
	if nfcKart == nil {
		return nil, errors.New("NFC card not found")
	}

	return nfcKart, nil
}

// GetByPersonelKodu retrieves NFC cards by personnel code
func (s *nfcKartService) GetByPersonelKodu(personelKodu string, page, limit int) ([]models.NFCKart, int64, error) {
	if personelKodu == "" {
		return nil, 0, errors.New("personel_kodu is required")
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindByPersonelKodu(personelKodu, page, limit)
}
