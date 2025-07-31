package controllers

import (
	"go-backend/models"
	"go-backend/services"
	"go-backend/utils"

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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "All fields are required: " + err.Error()})
	}

	// If the role is "patient" and the patient information (especially the TR ID number) is not received, give an error.
	if input.Role == "patient" && (input.PatientInfo == nil || input.PatientInfo.TCNumber == "") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "When the role is 'patient', PatientInfo field and TCNumber are required."})
	}

	// Create the User model to send to the service layer
	userToCreate := models.User{Name: input.Name, Role: input.Role, CardID: input.CardID}

	// Calling the service function with the correct parameters:
	// Incorrect old call: services.CreateUser(input.Name, input.Role, input.CardID)
	user, err := services.CreateUser(userToCreate, input.PatientInfo)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func GetUserByID(c *fiber.Ctx) error {
	userID, err := utils.ParseUserID(c, "id")
	if err != nil {
		return nil
	}
	user, err := services.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func DeleteUserByID(c *fiber.Ctx) error {
	userID, err := utils.ParseUserID(c, "id")
	if err != nil {
		return nil
	}
	err = services.DeleteUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User successfully deleted"})
}

// GetUserQRCode, returns the QR code generated from a specific user's CardID as a PNG image.
func GetUserQRCode(c *fiber.Ctx) error {
	userID, err := utils.ParseUserID(c, "id")
	if err != nil {
		return nil
	}
	pngData, err := services.GenerateQRCodeForUser(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	c.Set("Content-Type", "image/png")
	return c.Send(pngData)
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Users could not be retrieved"})
	}
	return c.Status(fiber.StatusOK).JSON(users)
}
