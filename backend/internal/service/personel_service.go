package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
)

type personelService struct {
	personelRepo repository.PersonelRepository
	nfcKartRepo  repository.NFCKartRepository
}

// NewPersonelService creates a new instance of PersonelService
func NewPersonelService(personelRepo repository.PersonelRepository, nfcKartRepo repository.NFCKartRepository) PersonelService {
	return &personelService{
		personelRepo: personelRepo,
		nfcKartRepo:  nfcKartRepo,
	}
}

// GetByKodu retrieves a personnel by their code
func (s *personelService) GetByKodu(kodu string) (*models.Personel, error) {
	if kodu == "" {
		return nil, errors.New("personel_kodu is required")
	}

	personel, err := s.personelRepo.FindByKodu(kodu)
	if err != nil {
		return nil, err
	}
	if personel == nil {
		return nil, errors.New("personel not found")
	}

	return personel, nil
}

// GetAll retrieves all personnel with pagination
func (s *personelService) GetAll(page, limit int) ([]models.Personel, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.personelRepo.FindAll(page, limit)
}

// GetByGorevKodu retrieves personnel by their role code
func (s *personelService) GetByGorevKodu(gorevKodu string, page, limit int) ([]models.Personel, int64, error) {
	if gorevKodu == "" {
		return nil, 0, errors.New("personel_gorev_kodu is required")
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.personelRepo.FindByGorevKodu(gorevKodu, page, limit)
}

// AuthenticateByNFC authenticates a personnel by NFC card UID
func (s *personelService) AuthenticateByNFC(kartUID string) (*models.Personel, error) {
	if kartUID == "" {
		return nil, errors.New("kart_uid is required")
	}

	// Find the NFC card by UID
	nfcKart, err := s.nfcKartRepo.FindByKartUID(kartUID)
	if err != nil {
		return nil, err
	}
	if nfcKart == nil {
		return nil, errors.New("NFC card not found")
	}

	// Check if the card is active (aktiflik_bilgisi = 1)
	if nfcKart.AktiflikBilgisi != 1 {
		return nil, errors.New("NFC card is inactive")
	}

	// Get the associated personnel
	personel, err := s.personelRepo.FindByKodu(nfcKart.PersonelKodu)
	if err != nil {
		return nil, err
	}
	if personel == nil {
		return nil, errors.New("associated personnel not found")
	}

	// Check if the personnel is active
	if personel.AktiflikBilgisi != 1 {
		return nil, errors.New("personnel account is inactive")
	}

	return personel, nil
}
