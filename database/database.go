package database

import (
	"fmt"
	"log"

	"github.com/Philip-21/Content/config"
	"github.com/Philip-21/Content/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initdb(connect *config.Envconfig) error {
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s  password=%s sslmode=%s",
		connect.Host, connect.Port, connect.DBName, connect.User, connect.Password, connect.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Content{}, &models.ContentUser{})
	DB = db
	return nil

}

// GetDB helps gets a connection,to the handlers , routes and database
func GetDB() *gorm.DB {
	return DB
}
