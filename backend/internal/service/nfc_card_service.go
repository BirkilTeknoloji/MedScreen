package service

import (
	"context"
	"errors"
	"medscreen/internal/models"
	"medscreen/internal/repository"
	"time"
)

type nfcCardService struct {
	repo repository.NFCCardRepository
}

// NewNFCCardService creates a new instance of NFCCardService
func NewNFCCardService(repo repository.NFCCardRepository) NFCCardService {
	return &nfcCardService{repo: repo}
}

// CreateCard creates a new NFC card with validation
func (s *nfcCardService) CreateCard(ctx context.Context, card *models.NFCCard) error {
	if card == nil {
		return errors.New("card cannot be nil")
	}

	// Validate required fields
	if card.CardUID == "" {
		return errors.New("card_uid is required")
	}

	// Validate card_uid uniqueness
	existing, err := s.repo.FindByCardUID(card.CardUID)
	if err == nil && existing != nil {
		return errors.New("card_uid already exists")
	}

	// Set default values
	if card.IssuedAt.IsZero() {
		card.IssuedAt = time.Now()
	}

	return s.repo.Create(ctx, card)
}

// GetCard retrieves an NFC card by ID
func (s *nfcCardService) GetCard(id uint) (*models.NFCCard, error) {
	if id == 0 {
		return nil, errors.New("invalid card id")
	}

	card, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if card == nil {
		return nil, errors.New("card not found")
	}

	return card, nil
}

// GetCards retrieves all NFC cards with pagination
func (s *nfcCardService) GetCards(page, limit int) ([]models.NFCCard, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindAll(page, limit)
}

// UpdateCard updates an existing NFC card with validation
func (s *nfcCardService) UpdateCard(ctx context.Context, id uint, card *models.NFCCard) error {
	if id == 0 {
		return errors.New("invalid card id")
	}
	if card == nil {
		return errors.New("card cannot be nil")
	}

	// Check if card exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("card not found")
	}

	// Validate card_uid uniqueness if being updated
	if card.CardUID != "" && card.CardUID != existing.CardUID {
		existingByUID, err := s.repo.FindByCardUID(card.CardUID)
		if err == nil && existingByUID != nil {
			return errors.New("card_uid already exists")
		}
	}

	// Set the ID to ensure we're updating the correct record
	card.ID = id

	return s.repo.Update(ctx, card)
}

// DeleteCard soft deletes an NFC card
func (s *nfcCardService) DeleteCard(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid card id")
	}

	// Check if card exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("card not found")
	}

	return s.repo.Delete(ctx, id)
}

// AssignCardToUser assigns an NFC card to a user
func (s *nfcCardService) AssignCardToUser(ctx context.Context, cardID, userID uint) error {
	if cardID == 0 {
		return errors.New("invalid card id")
	}
	if userID == 0 {
		return errors.New("invalid user id")
	}

	// Check if card exists
	card, err := s.repo.FindByID(cardID)
	if err != nil {
		return err
	}
	if card == nil {
		return errors.New("card not found")
	}

	// Update card with user assignment
	card.AssignedUserID = &userID
	card.AssignedUser = nil // Clear loaded association to avoid GORM issues
	card.IsActive = true

	return s.repo.Update(ctx, card)
}

// DeactivateCard deactivates an NFC card
func (s *nfcCardService) DeactivateCard(ctx context.Context, cardID uint) error {
	if cardID == 0 {
		return errors.New("invalid card id")
	}

	// Check if card exists
	card, err := s.repo.FindByID(cardID)
	if err != nil {
		return err
	}
	if card == nil {
		return errors.New("card not found")
	}

	// Deactivate the card
	card.IsActive = false

	return s.repo.Update(ctx, card)
}

// GetCardByUID retrieves an NFC card by card UID
func (s *nfcCardService) GetCardByUID(cardUID string) (*models.NFCCard, error) {
	if cardUID == "" {
		return nil, errors.New("card_uid is required")
	}

	card, err := s.repo.FindByCardUID(cardUID)
	if err != nil {
		return nil, err
	}
	if card == nil {
		return nil, errors.New("card not found")
	}

	return card, nil
}
