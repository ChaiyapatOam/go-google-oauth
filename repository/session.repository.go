package repository

import (
	"database/sql"

	"github.com/chaiyapatoam/go-google-oauth/domain"

	"github.com/chaiyapatoam/go-google-oauth/db"
)

type DBtest struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func GetSession(id string) (*domain.Session, error) {
	session := domain.Session{}
	err := db.DB.Get(&session, "SELECT * FROM session WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}

	return &session, err
}

func GetUser(email string) (*domain.Session, error) {
	session := domain.Session{}
	err := db.DB.Get(&session, "SELECT * FROM session WHERE email = ? LIMIT 1", email)
	if err != nil {
		return nil, err
	}

	return &session, err
}

func Create(session *domain.Session) (sql.Result, error) {

	result, err := db.DB.NamedExec("INSERT INTO session VALUES (:id, :email, :ip_address, :created_at, :expired_at) ", session)
	if err != nil {
		return nil, err
	}

	return result, nil
}
