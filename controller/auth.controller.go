package controller

import (
	"github.com/chaiyapatoam/go-google-oauth/service"
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	path := service.GoogleConfig()
	url := path.AuthCodeURL("state")
	return c.Redirect(url)
}

func CallBack(c *fiber.Ctx) error {
	token, err := service.GoogleConfig().Exchange(c.Context(), c.FormValue("code"))
	if err != nil {
		panic(err)
	}

	email := service.GetEmail(token.AccessToken)
	return c.Status(200).JSON(fiber.Map{"success": true, "email": email, "accessToken": token.AccessToken})
}
