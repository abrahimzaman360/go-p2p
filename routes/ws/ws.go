package ws

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func WebsocketRoutes(app *fiber.App) {
	app.Get("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return websocket.New(func(conn *websocket.Conn) {
				for {
					msgType, msg, err := conn.ReadMessage()
					if err != nil {
						break
					}
					if err := conn.WriteMessage(msgType, msg); err != nil {
						break
					}
				}
			})(c)
		}
		return fiber.ErrUpgradeRequired
	})
}
