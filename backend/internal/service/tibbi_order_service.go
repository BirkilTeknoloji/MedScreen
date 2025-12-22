package service

import (
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
)

type tibbiOrderService struct {
	repo repository.TibbiOrderRepository
}

// NewTibbiOrderService creates a new instance of TibbiOrderService
func NewTibbiOrderService(repo repository.TibbiOrderRepository) TibbiOrderService {
	return &tibbiOrderService{repo: repo}
}

// GetByKodu retrieves a medical order by its code
func (s *tibbiOrderService) GetByKodu(kodu string) (*models.TibbiOrder, error) {
	if kodu == "" {
		return nil, errors.New("tibbi_order_kodu is required")
	}

	order, err := s.repo.FindByKodu(kodu)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("medical order not found")
	}

	return order, nil
}

// GetByBasvuruKodu retrieves medical orders by patient visit code
func (s *tibbiOrderService) GetByBasvuruKodu(basvuruKodu string, page, limit int) ([]models.TibbiOrder, int64, error) {
	if basvuruKodu == "" {
		return nil, 0, errors.New("hasta_basvuru_kodu is required")
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindByBasvuruKodu(basvuruKodu, page, limit)
}

// GetDetayByOrderKodu retrieves medical order details by order code
func (s *tibbiOrderService) GetDetayByOrderKodu(orderKodu string, page, limit int) ([]models.TibbiOrderDetay, int64, error) {
	if orderKodu == "" {
		return nil, 0, errors.New("tibbi_order_kodu is required")
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindDetayByOrderKodu(orderKodu, page, limit)
}
