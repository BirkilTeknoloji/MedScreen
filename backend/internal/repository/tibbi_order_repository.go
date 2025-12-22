package repository

import (
	"medscreen/internal/models"

	"gorm.io/gorm"
)

// tibbiOrderRepository implements TibbiOrderRepository interface
type tibbiOrderRepository struct {
	db *gorm.DB
}

// NewTibbiOrderRepository creates a new TibbiOrderRepository instance
func NewTibbiOrderRepository(db *gorm.DB) TibbiOrderRepository {
	return &tibbiOrderRepository{db: db}
}

// FindByKodu retrieves a medical order by its code
func (r *tibbiOrderRepository) FindByKodu(kodu string) (*models.TibbiOrder, error) {
	var order models.TibbiOrder
	if err := r.db.Preload("HastaBasvuru").Preload("Hekim").Preload("Detaylar").
		Where("tibbi_order_kodu = ?", kodu).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

// FindByBasvuruKodu retrieves medical orders by visit code with pagination
func (r *tibbiOrderRepository) FindByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.TibbiOrder, int64, error) {
	var orders []models.TibbiOrder
	var total int64

	if err := r.db.Model(&models.TibbiOrder{}).Where("hasta_basvuru_kodu = ?", basvuruKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("HastaBasvuru").Preload("Hekim").Preload("Detaylar").
		Where("hasta_basvuru_kodu = ?", basvuruKodu).
		Order("order_zamani DESC").
		Offset(offset).Limit(limit).Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

// FindDetayByOrderKodu retrieves order details by order code with pagination
func (r *tibbiOrderRepository) FindDetayByOrderKodu(orderKodu string, page, limit int) ([]models.TibbiOrderDetay, int64, error) {
	var detaylar []models.TibbiOrderDetay
	var total int64

	if err := r.db.Model(&models.TibbiOrderDetay{}).Where("tibbi_order_kodu = ?", orderKodu).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("UygulayanPersonel").
		Where("tibbi_order_kodu = ?", orderKodu).
		Order("planlanan_uygulama_zamani ASC").
		Offset(offset).Limit(limit).Find(&detaylar).Error; err != nil {
		return nil, 0, err
	}

	return detaylar, total, nil
}
