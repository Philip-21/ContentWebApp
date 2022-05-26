package database

import (
	"fmt"

	"github.com/Philip-21/proj1/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initdb(connect *config.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password%s dbname=%s ",
		connect.Host, connect.Port, connect.User, connect.Password, connect.DBName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil

}
