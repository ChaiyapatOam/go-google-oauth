package authController

import (
	"fmt"

	"github.com/chaiyapatoam/go-google-oauth/service"
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	path := service.GoogleConfig()
	url := path.AuthCodeURL("state")

	// Return GoogleAuth Url to Frontend
	return c.Status(200).JSON(fiber.Map{"success": true, "url": url})
}

// Check User == email
// if User => get User
// if !User => Create User
// Create Session
// if session => delete and create new
// if !session => create
func CallBack(c *fiber.Ctx) error {
	token := service.GetToken(c)
	profile := service.GetProfile(token.AccessToken)
	fmt.Println(profile.Name, profile.Email)

	ipAddress := c.IP()

	// Return Cookie
	cookie, err := service.CreateCookie(ipAddress, profile.Email)
	if err != nil {
		fmt.Println(err)
	}
	user, err := service.GetUser(profile.Email)
	if err != nil {
		return err
	}
	if user == nil {
		_, err = service.CreateCookie(ipAddress, profile.Email)
		if err != nil {
			panic(err)
		}
	}
	// fmt.Println("Cookie:", cookie)
	c.Cookie(cookie)

	// Redirect to Frontend
	return c.Redirect("http://localhost:5173")
}
