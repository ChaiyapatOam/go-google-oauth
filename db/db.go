package db

import (
	"fmt"

	"github.com/chaiyapatoam/go-google-oauth/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDB() (*sqlx.DB, error) {
	var err error
	DB, err = sqlx.Connect("mysql", config.GetEnv("MYSQL_URI"))
	if err != nil {
		fmt.Println(err)
	}
	return DB, err
}
