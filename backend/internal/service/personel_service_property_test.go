package service

import (
	"errors"
	"medscreen/internal/models"
	"testing"
	"time"

	"pgregory.net/rapid"
)

// Feature: vem-database-migration, Property 6: Active Status Filter for Authentication
// Property 6: Active Status Filter for Authentication
// *For any* NFC card authentication attempt, the system SHALL only return a valid
// personnel record if the NFC card's aktiflik_bilgisi equals 1 (active).

// mockPersonelRepository is a mock implementation of PersonelRepository for testing
type mockPersonelRepository struct {
	personelMap map[string]*models.Personel
}

func newMockPersonelRepository() *mockPersonelRepository {
	return &mockPersonelRepository{
		personelMap: make(map[string]*models.Personel),
	}
}

func (m *mockPersonelRepository) FindByKodu(kodu string) (*models.Personel, error) {
	if personel, ok := m.personelMap[kodu]; ok {
		return personel, nil
	}
	return nil, nil
}

func (m *mockPersonelRepository) FindAll(page, limit int) ([]models.Personel, int64, error) {
	var result []models.Personel
	for _, p := range m.personelMap {
		result = append(result, *p)
	}
	return result, int64(len(result)), nil
}

func (m *mockPersonelRepository) FindByGorevKodu(gorevKodu string, page, limit int) ([]models.Personel, int64, error) {
	var result []models.Personel
	for _, p := range m.personelMap {
		if p.PersonelGorevKodu == gorevKodu {
			result = append(result, *p)
		}
	}
	return result, int64(len(result)), nil
}

func (m *mockPersonelRepository) addPersonel(p *models.Personel) {
	m.personelMap[p.PersonelKodu] = p
}

// mockNFCKartRepository is a mock implementation of NFCKartRepository for testing
type mockNFCKartRepository struct {
	kartMap map[string]*models.NFCKart
}

func newMockNFCKartRepository() *mockNFCKartRepository {
	return &mockNFCKartRepository{
		kartMap: make(map[string]*models.NFCKart),
	}
}

func (m *mockNFCKartRepository) FindByKodu(kodu string) (*models.NFCKart, error) {
	for _, k := range m.kartMap {
		if k.NFCKartKodu == kodu {
			return k, nil
		}
	}
	return nil, nil
}

func (m *mockNFCKartRepository) FindByKartUID(kartUID string) (*models.NFCKart, error) {
	if kart, ok := m.kartMap[kartUID]; ok {
		return kart, nil
	}
	return nil, nil
}

func (m *mockNFCKartRepository) FindByPersonelKodu(personelKodu string, page, limit int) ([]models.NFCKart, int64, error) {
	var result []models.NFCKart
	for _, k := range m.kartMap {
		if k.PersonelKodu == personelKodu {
			result = append(result, *k)
		}
	}
	return result, int64(len(result)), nil
}

func (m *mockNFCKartRepository) FindAll(page, limit int) ([]models.NFCKart, int64, error) {
	var result []models.NFCKart
	for _, k := range m.kartMap {
		result = append(result, *k)
	}
	return result, int64(len(result)), nil
}

func (m *mockNFCKartRepository) addKart(k *models.NFCKart) {
	m.kartMap[k.KartUID] = k
}

// generatePersonelKodu generates a random personnel code
func generatePersonelKodu(t *rapid.T) string {
	return "P" + rapid.StringMatching(`[0-9]{6}`).Draw(t, "personel_kodu")
}

// generateKartUID generates a random NFC card UID
func generateKartUID(t *rapid.T) string {
	return rapid.StringMatching(`[A-F0-9]{8}`).Draw(t, "kart_uid")
}

// generateNFCKartKodu generates a random NFC card code
func generateNFCKartKodu(t *rapid.T) string {
	return "NFC" + rapid.StringMatching(`[0-9]{6}`).Draw(t, "nfc_kart_kodu")
}

// TestProperty_ActiveStatusFilterForAuthentication tests Property 6
// For any NFC card authentication attempt, the system SHALL only return a valid
// personnel record if the NFC card's aktiflik_bilgisi equals 1 (active).
func TestProperty_ActiveStatusFilterForAuthentication(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Setup mock repositories
		personelRepo := newMockPersonelRepository()
		nfcKartRepo := newMockNFCKartRepository()

		// Generate random personnel
		personelKodu := generatePersonelKodu(t)
		personelAktif := rapid.IntRange(0, 1).Draw(t, "personel_aktif")
		personel := &models.Personel{
			PersonelKodu:         personelKodu,
			Ad:                   "Test",
			Soyadi:               "User",
			PersonelGorevKodu:    "DOKTOR",
			AktiflikBilgisi:      personelAktif,
			KayitZamani:          time.Now(),
			EkleyenKullaniciKodu: "SYSTEM",
		}
		personelRepo.addPersonel(personel)

		// Generate random NFC card
		kartUID := generateKartUID(t)
		kartAktif := rapid.IntRange(0, 1).Draw(t, "kart_aktif")
		nfcKart := &models.NFCKart{
			NFCKartKodu:          generateNFCKartKodu(t),
			PersonelKodu:         personelKodu,
			KartUID:              kartUID,
			AktiflikBilgisi:      kartAktif,
			KayitZamani:          time.Now(),
			EkleyenKullaniciKodu: "SYSTEM",
		}
		nfcKartRepo.addKart(nfcKart)

		// Create service
		svc := NewPersonelService(personelRepo, nfcKartRepo)

		// Attempt authentication
		result, err := svc.AuthenticateByNFC(kartUID)

		// Property verification:
		// Authentication should only succeed if BOTH card AND personnel are active
		shouldSucceed := kartAktif == 1 && personelAktif == 1

		if shouldSucceed {
			// Should return valid personnel
			if err != nil {
				t.Errorf("Expected successful authentication but got error: %v", err)
			}
			if result == nil {
				t.Error("Expected personnel result but got nil")
			}
			if result != nil && result.PersonelKodu != personelKodu {
				t.Errorf("Expected personnel code %s but got %s", personelKodu, result.PersonelKodu)
			}
		} else {
			// Should fail with error
			if err == nil {
				t.Error("Expected authentication to fail but it succeeded")
			}
			if result != nil {
				t.Error("Expected nil result for failed authentication")
			}
		}
	})
}

// TestProperty_InactiveCardRejectsAuthentication tests that inactive cards always fail
func TestProperty_InactiveCardRejectsAuthentication(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Setup mock repositories
		personelRepo := newMockPersonelRepository()
		nfcKartRepo := newMockNFCKartRepository()

		// Generate active personnel
		personelKodu := generatePersonelKodu(t)
		personel := &models.Personel{
			PersonelKodu:         personelKodu,
			Ad:                   "Test",
			Soyadi:               "User",
			PersonelGorevKodu:    "DOKTOR",
			AktiflikBilgisi:      1, // Active personnel
			KayitZamani:          time.Now(),
			EkleyenKullaniciKodu: "SYSTEM",
		}
		personelRepo.addPersonel(personel)

		// Generate INACTIVE NFC card
		kartUID := generateKartUID(t)
		nfcKart := &models.NFCKart{
			NFCKartKodu:          generateNFCKartKodu(t),
			PersonelKodu:         personelKodu,
			KartUID:              kartUID,
			AktiflikBilgisi:      0, // Inactive card
			KayitZamani:          time.Now(),
			EkleyenKullaniciKodu: "SYSTEM",
		}
		nfcKartRepo.addKart(nfcKart)

		// Create service
		svc := NewPersonelService(personelRepo, nfcKartRepo)

		// Attempt authentication
		result, err := svc.AuthenticateByNFC(kartUID)

		// Property: Inactive card should ALWAYS fail authentication
		if err == nil {
			t.Error("Expected authentication to fail for inactive card but it succeeded")
		}
		if result != nil {
			t.Error("Expected nil result for inactive card authentication")
		}
	})
}

// TestProperty_InactivePersonnelRejectsAuthentication tests that inactive personnel always fail
func TestProperty_InactivePersonnelRejectsAuthentication(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Setup mock repositories
		personelRepo := newMockPersonelRepository()
		nfcKartRepo := newMockNFCKartRepository()

		// Generate INACTIVE personnel
		personelKodu := generatePersonelKodu(t)
		personel := &models.Personel{
			PersonelKodu:         personelKodu,
			Ad:                   "Test",
			Soyadi:               "User",
			PersonelGorevKodu:    "DOKTOR",
			AktiflikBilgisi:      0, // Inactive personnel
			KayitZamani:          time.Now(),
			EkleyenKullaniciKodu: "SYSTEM",
		}
		personelRepo.addPersonel(personel)

		// Generate active NFC card
		kartUID := generateKartUID(t)
		nfcKart := &models.NFCKart{
			NFCKartKodu:          generateNFCKartKodu(t),
			PersonelKodu:         personelKodu,
			KartUID:              kartUID,
			AktiflikBilgisi:      1, // Active card
			KayitZamani:          time.Now(),
			EkleyenKullaniciKodu: "SYSTEM",
		}
		nfcKartRepo.addKart(nfcKart)

		// Create service
		svc := NewPersonelService(personelRepo, nfcKartRepo)

		// Attempt authentication
		result, err := svc.AuthenticateByNFC(kartUID)

		// Property: Inactive personnel should ALWAYS fail authentication
		if err == nil {
			t.Error("Expected authentication to fail for inactive personnel but it succeeded")
		}
		if result != nil {
			t.Error("Expected nil result for inactive personnel authentication")
		}
	})
}

// TestProperty_NonExistentCardRejectsAuthentication tests that non-existent cards fail
func TestProperty_NonExistentCardRejectsAuthentication(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Setup mock repositories with no data
		personelRepo := newMockPersonelRepository()
		nfcKartRepo := newMockNFCKartRepository()

		// Create service
		svc := NewPersonelService(personelRepo, nfcKartRepo)

		// Generate random card UID that doesn't exist
		kartUID := generateKartUID(t)

		// Attempt authentication
		result, err := svc.AuthenticateByNFC(kartUID)

		// Property: Non-existent card should ALWAYS fail authentication
		if err == nil {
			t.Error("Expected authentication to fail for non-existent card but it succeeded")
		}
		if result != nil {
			t.Error("Expected nil result for non-existent card authentication")
		}
	})
}

// TestProperty_EmptyKartUIDRejectsAuthentication tests that empty card UID fails
func TestProperty_EmptyKartUIDRejectsAuthentication(t *testing.T) {
	// Setup mock repositories
	personelRepo := newMockPersonelRepository()
	nfcKartRepo := newMockNFCKartRepository()

	// Create service
	svc := NewPersonelService(personelRepo, nfcKartRepo)

	// Attempt authentication with empty UID
	result, err := svc.AuthenticateByNFC("")

	// Property: Empty card UID should ALWAYS fail authentication
	if err == nil {
		t.Error("Expected authentication to fail for empty card UID but it succeeded")
	}
	if result != nil {
		t.Error("Expected nil result for empty card UID authentication")
	}
	expectedErr := errors.New("kart_uid is required")
	if err.Error() != expectedErr.Error() {
		t.Errorf("Expected error '%s' but got '%s'", expectedErr.Error(), err.Error())
	}
}
