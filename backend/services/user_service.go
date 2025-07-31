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
	ErrScannerNotFound         = errors.New("The user who scanned the card could not be found.")
	ErrPatientNotFound         = errors.New("The requested patient could not be found.")
	ErrRequestedUserNotPatient = errors.New("The requested user is not a patient.")
	ErrPermissionDenied        = errors.New("You do not have permission to view this patient's information.")
	ErrPatientSelfViewOnly     = errors.New("Patients can only view their own information.")
	ErrScannerNotAuthorized    = errors.New("You are not authorized to perform this action (doctors only).")
	ErrPatientFromQRNotFound   = errors.New("No patient found matching the QR code.")
	ErrUserHasNoCardID         = errors.New("This user does not have a card ID.")
)

// MedicationResponse, it contains the medication information that will be returned when the QR code is scanned.
type MedicationResponse struct {
	PatientName   string   `json:"patient_name"`
	Prescriptions []string `json:"prescriptions"`
}

// GetPatientDataForView, contains the authorization and data fetching logic for the bedside device.
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

// GetMedicationByQRCode, it verifies and retrieves the patient's medication information using the QR code data and the reader's card.
func GetMedicationByQRCode(scannerCardID, qrData string) (MedicationResponse, error) {
	var scanner models.User
	if err := config.DB.Where("card_id = ?", scannerCardID).First(&scanner).Error; err != nil {
		return MedicationResponse{}, ErrScannerNotFound
	}

	if scanner.Role != "doctor" {
		return MedicationResponse{}, ErrScannerNotAuthorized
	}

	var patient models.User
	if err := config.DB.Preload("PatientInfo").Where("card_id = ?", qrData).First(&patient).Error; err != nil {
		return MedicationResponse{}, ErrPatientFromQRNotFound
	}

	if patient.Role != "patient" {
		return MedicationResponse{}, ErrPatientFromQRNotFound
	}

	var prescriptionList []string
	for _, prescription := range patient.PatientInfo.Prescriptions {
		prescriptionList = append(prescriptionList, fmt.Sprintf("%v", prescription))
	}
	response := MedicationResponse{
		PatientName:   patient.Name,
		Prescriptions: prescriptionList,
	}

	return response, nil
}

// CreateUser creates a new user and, if role is patient, creates PatientInfo as well.
func CreateUser(user models.User, patientInfo *models.PatientInfo) (models.User, error) {
	role := strings.ToLower(user.Role)
	if role != "personnel" && role != "patient" {
		return models.User{}, errors.New("Invalid role: must be 'personnel' or 'patient'")
	}

	newUser := models.User{
		Name:   user.Name,
		Role:   role,
		CardID: user.CardID,
	}

	tx := config.DB.Begin()
	if tx.Error != nil {
		return models.User{}, fmt.Errorf("Transaction could not be started: %w", tx.Error)
	}

	if err := tx.Create(&newUser).Error; err != nil {
		tx.Rollback()
		if strings.Contains(err.Error(), "Duplicate key value violates unique constraint") {
			return models.User{}, errors.New("This CardID or another unique field is already in use")
		}
		return models.User{}, fmt.Errorf("Could not create user: %w", err)
	}

	if role == "patient" && patientInfo != nil {
		patientInfo.UserID = newUser.ID
		if err := tx.Create(patientInfo).Error; err != nil {
			tx.Rollback()
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				return models.User{}, errors.New("No matching TC ID number or UserID found")
			}
			return models.User{}, fmt.Errorf("Patient information could not be created: %w", err)
		}
		newUser.PatientInfo = *patientInfo
	}

	return newUser, tx.Commit().Error
}

// GenerateQRCodeForUser, creates a QR code (PNG) from a user's CardID based on that user's ID.
func GenerateQRCodeForUser(userID uint) ([]byte, error) {
	user, err := GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	if user.CardID == "" {
		return nil, ErrUserHasNoCardID
	}

	png, err := qrcode.Encode(user.CardID, qrcode.Medium, 256)
	return png, err
}

// GetUserByID returns user by ID, including PatientInfo.
func GetUserByID(id uint) (models.User, error) {
	var user models.User
	if err := config.DB.Preload("PatientInfo").First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// DeleteUserByID deletes user by ID.
func DeleteUserByID(id uint) error {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return errors.New("User not found")
	}
	if err := config.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByCardID returns user by CardID, including PatientInfo.
func GetUserByCardID(cardID string) (models.User, error) {
	var user models.User
	if err := config.DB.Preload("PatientInfo").Where("card_id = ?", cardID).First(&user).Error; err != nil {
		return models.User{}, errors.New("No user found matching the card ID")
	}
	return user, nil
}

// GetAllUsers returns all users, including PatientInfo.
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Preload("PatientInfo").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
