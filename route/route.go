package route

import (
	"github.com/chaiyapatoam/go-google-oauth/controller"
	"github.com/gofiber/fiber/v2"
)

func MainRoute(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/login", controller.Auth)
	api.Get("/auth/google/callback", controller.CallBack)
}
