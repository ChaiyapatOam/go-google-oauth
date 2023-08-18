package route

import (
	"github.com/chaiyapatoam/go-google-oauth/controller"
	"github.com/gofiber/fiber/v2"
)

func MainRoute(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/login", authController.Auth)
	api.Get("/auth/google/callback", authController.CallBack)

	auth := app.Group("/auth") //Need Auth
	auth.Get("/", func(c *fiber.Ctx) error {
		ssid := c.Cookies("UUID")
		return c.JSON(&fiber.Map{"message": "Hello Auth", "cookie": ssid})
	})
}
