package database

import (
	"fmt"

	"github.com/Philip-21/proj1/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initdb(connect *config.Envconfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s  password=%s sslmode=%s",
		connect.Host, connect.Port, connect.DBName, connect.User, connect.Password, connect.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil

}
