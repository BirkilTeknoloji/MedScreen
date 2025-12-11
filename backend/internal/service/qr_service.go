package service

import (
	"context"
	"medscreen/internal/models"
	"medscreen/internal/repository"
	"encoding/base64"
	"errors"
	"time"

	"github.com/google/uuid"
	qrcode "github.com/skip2/go-qrcode"
	"gorm.io/gorm"
)

// QRService defines the interface for QR code operations
type QRService interface {
	GeneratePatientAssignmentQR(ctx context.Context, patientID uint, expiryHours int) (string, string, error)
	GeneratePrescriptionInfoQR(ctx context.Context, patientID uint, deviceID uint, expiryHours int) (string, string, error)
	ValidateToken(tokenStr string) (*models.QRToken, error)
	AssignPatientToDevice(ctx context.Context, tokenStr string, macAddress string) error
}

type qrService struct {
	qrTokenRepo   repository.QRTokenRepository
	patientRepo   repository.PatientRepository
	deviceRepo    repository.DeviceRepository
}

// NewQRService creates a new QR service instance
func NewQRService(
	qrTokenRepo repository.QRTokenRepository,
	patientRepo repository.PatientRepository,
	deviceRepo repository.DeviceRepository,
) QRService {
	return &qrService{
		qrTokenRepo: qrTokenRepo,
		patientRepo: patientRepo,
		deviceRepo:  deviceRepo,
	}
}

// GeneratePatientAssignmentQR generates a QR code for patient assignment
func (s *qrService) GeneratePatientAssignmentQR(ctx context.Context, patientID uint, expiryHours int) (string, string, error) {
	// Verify patient exists
	_, err := s.patientRepo.FindByID(patientID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", "", errors.New("patient not found")
		}
		return "", "", err
	}

	// Generate unique token
	token := uuid.New().String()

	// Create QR token record
	qrToken := &models.QRToken{
		Token:     token,
		Type:      models.QRTypePatientAssignment,
		PatientID: patientID,
		ExpiresAt: time.Now().Add(time.Duration(expiryHours) * time.Hour),
		IsUsed:    false,
	}

	if err := s.qrTokenRepo.Create(ctx, qrToken); err != nil {
		return "", "", err
	}

	// Generate QR code image
	qrImage, err := s.generateQRImage(token)
	if err != nil {
		return "", "", err
	}

	return token, qrImage, nil
}

// GeneratePrescriptionInfoQR generates a QR code for prescription information
func (s *qrService) GeneratePrescriptionInfoQR(ctx context.Context, patientID uint, deviceID uint, expiryHours int) (string, string, error) {
	// Verify patient exists
	_, err := s.patientRepo.FindByID(patientID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", "", errors.New("patient not found")
		}
		return "", "", err
	}

	// Generate unique token
	token := uuid.New().String()

	// Create QR token record
	qrToken := &models.QRToken{
		Token:     token,
		Type:      models.QRTypePrescriptionInfo,
		PatientID: patientID,
		DeviceID:  &deviceID,
		ExpiresAt: time.Now().Add(time.Duration(expiryHours) * time.Hour),
		IsUsed:    false,
	}

	if err := s.qrTokenRepo.Create(ctx, qrToken); err != nil {
		return "", "", err
	}

	// Generate QR code image
	qrImage, err := s.generateQRImage(token)
	if err != nil {
		return "", "", err
	}

	return token, qrImage, nil
}

// ValidateToken validates a QR token and returns the token data
func (s *qrService) ValidateToken(tokenStr string) (*models.QRToken, error) {
	qrToken, err := s.qrTokenRepo.FindByToken(tokenStr)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid token")
		}
		return nil, err
	}

	if !qrToken.CanBeUsed() {
		if qrToken.IsExpired() {
			return nil, errors.New("token has expired")
		}
		return nil, errors.New("token has already been used")
	}

	return qrToken, nil
}

// AssignPatientToDevice assigns a patient to a device using a QR token
func (s *qrService) AssignPatientToDevice(ctx context.Context, tokenStr string, macAddress string) error {
	// Validate token
	qrToken, err := s.ValidateToken(tokenStr)
	if err != nil {
		return err
	}

	// Ensure it's a patient assignment token
	if qrToken.Type != models.QRTypePatientAssignment {
		return errors.New("invalid token type for patient assignment")
	}

	// Find device by MAC address
	device, err := s.deviceRepo.GetByMAC(macAddress)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("device not found")
		}
		return err
	}

	// Assign patient to device
	device.PatientID = &qrToken.PatientID
	if err := s.deviceRepo.Update(ctx, device); err != nil {
		return err
	}

	// Mark token as used
	now := time.Now()
	qrToken.IsUsed = true
	qrToken.UsedAt = &now
	if err := s.qrTokenRepo.Update(ctx, qrToken); err != nil {
		return err
	}

	return nil
}

// generateQRImage generates a base64-encoded PNG QR code image
func (s *qrService) generateQRImage(token string) (string, error) {
	// Generate QR code with medium error correction
	png, err := qrcode.Encode(token, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}

	// Encode to base64
	base64Image := base64.StdEncoding.EncodeToString(png)
	return "data:image/png;base64," + base64Image, nil
}
