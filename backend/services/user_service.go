package services

import (
	"errors"
	"fmt"
	"go-backend/config"
	"go-backend/models"
	"strings"

	"github.com/skip2/go-qrcode"
)

// Custom errors for device view logic
var (
	ErrScannerNotFound         = errors.New("kartı okutan kullanıcı bulunamadı")
	ErrPatientNotFound         = errors.New("bilgisi istenen hasta bulunamadı")
	ErrRequestedUserNotPatient = errors.New("bilgisi istenen kullanıcı bir hasta değil")
	ErrPermissionDenied        = errors.New("bu hasta bilgilerini görme yetkiniz yok")
	// ErrPatientSelfViewOnly, bir hastanın başka bir hastanın verisini görmeye çalışması durumunda kullanılır.
	ErrPatientSelfViewOnly   = errors.New("hastalar sadece kendi bilgilerini görebilir")
	ErrScannerNotAuthorized  = errors.New("bu işlemi yapma yetkiniz yok (sadece doktorlar)")
	ErrPatientFromQRNotFound = errors.New("QR kod ile eşleşen hasta bulunamadı")
	ErrUserHasNoCardID       = errors.New("bu kullanıcının bir kart ID'si bulunmuyor")
)

// MedicationResponse, QR kod okutulduğunda dönecek olan ilaç bilgilerini içerir.
type MedicationResponse struct {
	PatientName   string   `json:"patient_name"`
	Prescriptions []string `json:"prescriptions"`
}

// GetPatientDataForView, yatak başı cihazı için yetkilendirme ve veri getirme mantığını içerir.
func GetPatientDataForView(scannerCardID string, patientUserID uint) (models.User, error) {
	var scanner models.User
	if err := config.DB.Where("card_id = ?", scannerCardID).First(&scanner).Error; err != nil {
		return models.User{}, ErrScannerNotFound
	}

	var patient models.User
	if err := config.DB.Preload("PatientInfo").First(&patient, patientUserID).Error; err != nil {
		return models.User{}, ErrPatientNotFound
	}

	if patient.Role != "patient" {
		return models.User{}, ErrRequestedUserNotPatient
	}

	switch scanner.Role {
	case "doctor":
		return patient, nil
	case "patient":
		if scanner.ID == patient.ID {
			return patient, nil
		}
		return models.User{}, ErrPatientSelfViewOnly
	default:
		return models.User{}, ErrPermissionDenied
	}
}

// GetMedicationByQRCode, QR kod verisini ve okutan kişinin kartını kullanarak hastanın ilaç bilgilerini doğrular ve getirir.
func GetMedicationByQRCode(scannerCardID, qrData string) (MedicationResponse, error) {
	var scanner models.User
	// 1. Kartı okutan personeli bul
	if err := config.DB.Where("card_id = ?", scannerCardID).First(&scanner).Error; err != nil {
		return MedicationResponse{}, ErrScannerNotFound
	}

	// 2. Personelin yetkisini kontrol et (sadece doktorlar)
	if scanner.Role != "doctor" {
		return MedicationResponse{}, ErrScannerNotAuthorized
	}

	// 3. QR kod verisi (hastanın CardID'si) ile hastayı bul
	var patient models.User
	if err := config.DB.Preload("PatientInfo").Where("card_id = ?", qrData).First(&patient).Error; err != nil {
		return MedicationResponse{}, ErrPatientFromQRNotFound
	}

	// 4. Bulunan kullanıcının rolünün "patient" olduğunu doğrula
	if patient.Role != "patient" {
		return MedicationResponse{}, ErrPatientFromQRNotFound // QR kod bir doktora aitse de hata verelim
	}

	// 5. İlaç bilgilerini (models.Prescriptions) string dilimine ([]string) çevir ve döndür.
	// `models.Prescriptions` muhtemelen bir struct dilimidir (örn: []Prescription).
	// Bu nedenle her bir eleman üzerinden geçip istediğimiz string temsilini oluşturmalıyız.
	var prescriptionList []string
	for _, prescription := range patient.PatientInfo.Prescriptions {
		// Burada `prescription`'ın nasıl bir string'e dönüştürüleceğini belirtmelisiniz.
		// Örneğin, eğer bir struct ise ve 'Name' alanı varsa: prescription.Name
		// Model tanımını bilmediğimiz için `fmt.Sprintf` ile genel bir çevrim yapıyoruz.
		prescriptionList = append(prescriptionList, fmt.Sprintf("%v", prescription))
	}
	response := MedicationResponse{
		PatientName:   patient.Name,
		Prescriptions: prescriptionList,
	}

	return response, nil
}

func CreateUser(user models.User, patientInfo *models.PatientInfo) (models.User, error) {
	role := strings.ToLower(user.Role)
	if role != "doctor" && role != "patient" {
		return models.User{}, errors.New("Geçersiz rol: sadece 'doctor' veya 'patient' olabilir")
	}

	// Gelen verileri ata
	newUser := models.User{
		Name:   user.Name,
		Role:   role,
		CardID: user.CardID,
	}

	// Transaction başlat
	tx := config.DB.Begin()
	if tx.Error != nil {
		return models.User{}, fmt.Errorf("transaction başlatılamadı: %w", tx.Error)
	}

	// 1. Kullanıcıyı oluştur
	if err := tx.Create(&newUser).Error; err != nil {
		tx.Rollback()
		// Hata mesajını daha anlaşılır hale getirelim
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return models.User{}, errors.New("bu CardID veya başka bir benzersiz alan zaten kullanılıyor")
		}
		return models.User{}, fmt.Errorf("kullanıcı oluşturulamadı: %w", err)
	}

	// Eğer kullanıcı hasta ise ve hasta bilgisi varsa PatientInfo oluştur
	if role == "patient" && patientInfo != nil {
		patientInfo.UserID = newUser.ID // Oluşturulan User'ın ID'sini ata
		if err := tx.Create(patientInfo).Error; err != nil {
			tx.Rollback()
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				return models.User{}, errors.New("bu TC Kimlik Numarası veya UserID zaten kullanılıyor")
			}
			return models.User{}, fmt.Errorf("hasta bilgisi oluşturulamadı: %w", err)
		}
		newUser.PatientInfo = *patientInfo
	}

	// Her şey yolundaysa transaction'ı onayla
	return newUser, tx.Commit().Error
}

// GenerateQRCodeForUser, bir kullanıcı ID'sine göre o kullanıcının CardID'sinden bir QR kod (PNG) oluşturur.
func GenerateQRCodeForUser(userID uint) ([]byte, error) {
	// 1. Kullanıcıyı veritabanından bul
	user, err := GetUserByID(userID)
	if err != nil {
		// GetUserByID zaten "not found" hatası dönecektir.
		return nil, err
	}

	// 2. Kullanıcının bir CardID'si olduğundan emin ol
	if user.CardID == "" {
		return nil, ErrUserHasNoCardID
	}

	// 3. CardID'yi kullanarak 256x256 boyutunda bir QR kod PNG'si oluştur
	// qrcode.Encode(data string, level RecoveryLevel, size int) ([]byte, error)
	png, err := qrcode.Encode(user.CardID, qrcode.Medium, 256)
	return png, err
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	// Preload ile ilişkili PatientInfo verisini de çekiyoruz
	if err := config.DB.Preload("PatientInfo").First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// Kullanıcıyı ID ile sil
func DeleteUserByID(id uint) error {
	var user models.User

	// Kullanıcıyı önce bul
	if err := config.DB.First(&user, id).Error; err != nil {
		return errors.New("Kullanıcı bulunamadı")
	}

	// Sil (soft delete)
	// Not: İlişkili PatientInfo'yu da silmek isterseniz, burada transaction içinde ek bir silme işlemi gerekir.
	// Şimdilik sadece User siliniyor.
	if err := config.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

// CardID ile kullanıcı getir
func GetUserByCardID(cardID string) (models.User, error) {
	var user models.User
	if err := config.DB.Preload("PatientInfo").Where("card_id = ?", cardID).First(&user).Error; err != nil {
		return models.User{}, errors.New("Kart ile eşleşen kullanıcı bulunamadı")
	}
	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Preload("PatientInfo").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
