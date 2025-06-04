package controllers

import (
	"main/database"
	"main/models"

	"github.com/gofiber/fiber/v2"
)

// GetAuth handles user authentication
func PostLogin(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	var dbUser models.User
	if err := database.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&dbUser).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	return c.JSON(dbUser)
}
