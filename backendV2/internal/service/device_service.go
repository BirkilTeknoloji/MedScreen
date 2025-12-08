package service

import (
	"context"
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
)

// DeviceUpdateRequest is a DTO for partial device updates
type DeviceUpdateRequest struct {
	RoomNumber  *string `json:"room_number"`
	Description *string `json:"description"`
	IsActive    *bool   `json:"is_active"`
}

type deviceService struct {
	deviceRepo  repository.DeviceRepository
	patientRepo repository.PatientRepository
}

// NewDeviceService creates a new device service
func NewDeviceService(deviceRepo repository.DeviceRepository, patientRepo repository.PatientRepository) DeviceService {
	return &deviceService{
		deviceRepo:  deviceRepo,
		patientRepo: patientRepo,
	}
}

func (s *deviceService) RegisterDevice(ctx context.Context, device *models.Device) error {
	// Check if device already exists
	existingDevice, _ := s.deviceRepo.GetByMAC(device.MACAddress)
	if existingDevice != nil {
		return errors.New("device with this mac address already exists")
	}
	return s.deviceRepo.Create(ctx, device)
}

func (s *deviceService) GetDeviceByMAC(mac string) (*models.Device, error) {
	return s.deviceRepo.GetByMAC(mac)
}

func (s *deviceService) GetDeviceByID(id uint) (*models.Device, error) {
	return s.deviceRepo.GetByID(id)
}

func (s *deviceService) GetAllDevices(page, limit int) ([]models.Device, int64, error) {
	return s.deviceRepo.FindAll(page, limit)
}

func (s *deviceService) AssignPatient(ctx context.Context, mac string, patientID uint) error {
	device, err := s.deviceRepo.GetByMAC(mac)
	if err != nil {
		return err
	}
	if device == nil {
		return errors.New("device not found")
	}

	// Validate patient exists before assignment
	patient, err := s.patientRepo.FindByID(patientID)
	if err != nil {
		return errors.New("patient not found")
	}
	if patient == nil {
		return errors.New("patient not found")
	}

	device.PatientID = &patientID
	device.Patient = nil // Clear loaded association to avoid GORM issues
	return s.deviceRepo.Update(ctx, device)
}

func (s *deviceService) UnassignPatient(ctx context.Context, mac string) error {
	device, err := s.deviceRepo.GetByMAC(mac)
	if err != nil {
		return err
	}
	if device == nil {
		return errors.New("device not found")
	}

	device.PatientID = nil
	device.Patient = nil // Clear loaded association to avoid GORM issues
	return s.deviceRepo.Update(ctx, device)
}

func (s *deviceService) UpdateDevice(ctx context.Context, mac string, updates *DeviceUpdateRequest) (*models.Device, error) {
	device, err := s.deviceRepo.GetByMAC(mac)
	if err != nil {
		return nil, err
	}
	if device == nil {
		return nil, errors.New("device not found")
	}

	// Apply non-nil updates
	if updates.RoomNumber != nil {
		device.RoomNumber = updates.RoomNumber
	}
	if updates.Description != nil {
		device.Description = updates.Description
	}
	if updates.IsActive != nil {
		device.IsActive = *updates.IsActive
	}

	if err := s.deviceRepo.Update(ctx, device); err != nil {
		return nil, err
	}

	return device, nil
}

func (s *deviceService) DeleteDevice(ctx context.Context, mac string) error {
	return s.deviceRepo.DeleteByMAC(ctx, mac)
}

func (s *deviceService) GetDevicesByFilters(roomNumber *string, patientID *uint, page, limit int) ([]models.Device, int64, error) {
	return s.deviceRepo.FindByFilters(roomNumber, patientID, page, limit)
}
