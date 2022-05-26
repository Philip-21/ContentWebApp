package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Philip-21/proj1/config"
	"github.com/Philip-21/proj1/database"
	"github.com/Philip-21/proj1/models"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

const portNumber = ":8080"

var app config.AppConfig

type Repository struct {
	DB *gorm.DB
}

//var db *gorm.DB

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	connect := &config.PostgresConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := database.Initdb(connect)

	if err != nil {
		log.Fatal("could not load the database")
	}
	err = models.MigrateContent(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	srv := &http.Server{
		Addr:    portNumber,
		Handler: Route(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}
