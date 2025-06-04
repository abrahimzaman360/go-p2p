package controllers

import (
	"main/database"
	"main/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}
