package controllers

import (
	"go-backend/models"
	"go-backend/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var input struct {
		Name        string              `json:"Name" validate:"required"`
		Role        string              `json:"Role" validate:"required"`
		CardID      string              `json:"CardID" validate:"required"`
		PatientInfo *models.PatientInfo `json:"PatientInfo,omitempty"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Tüm alanlar gereklidir: " + err.Error()})
	}

	// Eğer rol "patient" ise ve hasta bilgisi (özellikle TC no) gelmediyse hata ver
	if input.Role == "patient" && (input.PatientInfo == nil || input.PatientInfo.TCNumber == "") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Rol 'patient' olduğunda PatientInfo alanı ve TCNumber zorunludur."})
	}

	// Servis katmanına göndermek için User modelini oluştur
	userToCreate := models.User{Name: input.Name, Role: input.Role, CardID: input.CardID}

	// Servis fonksiyonunu doğru parametrelerle çağırıyoruz:
	// Hatalı eski çağrı: services.CreateUser(input.Name, input.Role, input.CardID)
	user, err := services.CreateUser(userToCreate, input.PatientInfo)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func GetUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz ID"})
	}

	user, err := services.GetUserByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Kullanıcı bulunamadı"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func DeleteUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz kullanıcı ID"})
	}

	err = services.DeleteUserByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Kullanıcı başarıyla silindi"})
}
func GetUserByCardID(c *fiber.Ctx) error {
	cardID := c.Params("card_id")

	user, err := services.GetUserByCardID(cardID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func GetAllUsers(c *fiber.Ctx) error {
	users, err := services.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Kullanıcılar alınamadı"})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
