package service

import (
	"time"

	"github.com/chaiyapatoam/go-google-oauth/domain"
	"github.com/chaiyapatoam/go-google-oauth/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateCookie(ipAddress string, email string) (*fiber.Cookie, error) {
	id := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)
	createdAt := time.Now()

	// Create In DB
	_, err := repository.Create(&domain.Session{
		Id:        id,
		Email:     email,
		CreatedAt: createdAt,
		IpAddress: ipAddress,
		ExpiredAt: expiresAt,
	})
	if err != nil {
		return nil, err
	}

	cookie := &fiber.Cookie{
		Name:     "ssid",
		Value:    id,
		HTTPOnly: true,
		Secure:   true,
		Expires:  expiresAt,
	}
	return cookie, nil
}
