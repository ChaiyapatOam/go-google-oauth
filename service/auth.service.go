package service

import (
	"github.com/chaiyapatoam/go-google-oauth/domain"
	"github.com/chaiyapatoam/go-google-oauth/repository"
)

func GetUser(email string) (*domain.Session, error) {
	user, err := repository.GetUser(email)
	if err != nil {
		return nil, err
	}
	
	return user, nil
}
