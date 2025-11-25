package repository

import (
	"medscreen/internal/models"

	"gorm.io/gorm"
)

type deviceRepository struct {
	db *gorm.DB
}

// NewDeviceRepository creates a new device repository
func NewDeviceRepository(db *gorm.DB) DeviceRepository {
	return &deviceRepository{db: db}
}

func (r *deviceRepository) Create(device *models.Device) error {
	return r.db.Create(device).Error
}

func (r *deviceRepository) GetByMAC(mac string) (*models.Device, error) {
	var device models.Device
	err := r.db.Preload("Patient").Where("mac_address = ?", mac).First(&device).Error
	if err != nil {
		return nil, err
	}
	return &device, nil
}

func (r *deviceRepository) GetByID(id uint) (*models.Device, error) {
	var device models.Device
	err := r.db.Preload("Patient").First(&device, id).Error
	if err != nil {
		return nil, err
	}
	return &device, nil
}

func (r *deviceRepository) Update(device *models.Device) error {
	return r.db.Save(device).Error
}

func (r *deviceRepository) Delete(id uint) error {
	return r.db.Delete(&models.Device{}, id).Error
}

func (r *deviceRepository) FindAll(page, limit int) ([]models.Device, int64, error) {
	var devices []models.Device
	var total int64
	offset := (page - 1) * limit

	if err := r.db.Model(&models.Device{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Preload("Patient").Limit(limit).Offset(offset).Find(&devices).Error
	if err != nil {
		return nil, 0, err
	}

	return devices, total, nil
}

func (r *deviceRepository) DeleteByMAC(mac string) error {
	var device models.Device
	err := r.db.Where("mac_address = ?", mac).First(&device).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&device).Error
}

func (r *deviceRepository) FindByFilters(roomNumber *string, patientID *uint, page, limit int) ([]models.Device, int64, error) {
	var devices []models.Device
	var total int64
	offset := (page - 1) * limit

	query := r.db.Model(&models.Device{})

	if roomNumber != nil {
		query = query.Where("room_number = ?", *roomNumber)
	}
	if patientID != nil {
		query = query.Where("patient_id = ?", *patientID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload("Patient").Limit(limit).Offset(offset).Find(&devices).Error
	if err != nil {
		return nil, 0, err
	}

	return devices, total, nil
}
