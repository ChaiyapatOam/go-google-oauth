package domain

import (
	"time"
)

type Session struct {
	Id        string    `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	IpAddress string    `json:"ipAddress" db:"ip_address"`
	ExpiredAt time.Time `json:"expiredAt" db:"expired_at"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
