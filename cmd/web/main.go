package main

import (
	"log"
	"net/http"

	"github.com/Philip-21/proj1/config"
	"github.com/Philip-21/proj1/database"
	"github.com/Philip-21/proj1/models"

	"gorm.io/gorm"
)

const portNumber = ":8080"

var app config.AppConfig

type Repository struct {
	DB *gorm.DB
}

//var db *gorm.DB

func main() {

	config.LoadConfig()
	db, err := database.Initdb(config.Conf)
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
