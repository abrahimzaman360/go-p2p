// cmd/main.go
package main

import (
	"main/config"
	"main/database"
	routes "main/routes/rest"
	ws "main/routes/ws"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		ExposeHeaders:    "Content-Length,Authorization",
		AllowCredentials: false,
		MaxAge:           300,
	}))

	config.LoadEnv()
	database.Connect()
	routes.UserRoutes(app)
	ws.WebsocketRoutes(app)

	// Root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("public/index.html")
	})

	// Serve static files
	app.Static("/assets", "./public/assets")

	app.Listen(":3000")
}
