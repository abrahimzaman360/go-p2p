package routes

import (
	"main/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/users", controllers.GetUsers)
}
