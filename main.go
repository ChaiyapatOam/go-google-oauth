package main

import (
	"log"

	"github.com/chaiyapatoam/go-google-oauth/db"
	"github.com/chaiyapatoam/go-google-oauth/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db.ConnectDB()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"message": "Hello World"})
	})

	route.MainRoute(app)
	log.Fatal(app.Listen(":5000"))
}
