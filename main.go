package main

import (
	"encoding/gob"
	"fmt"
	"log"

	"github.com/Philip-21/Content/config"
	"github.com/Philip-21/Content/handlers"

	"github.com/Philip-21/Content/database"
	"github.com/Philip-21/Content/routes"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {

	gob.Register(map[string]int{})

	//DATABASE
	config.LoadConfig() //load viper configurations
	log.Println("Connecting to database...")

	database.Initdb(config.Conf) //calling the db function

	log.Println("Connected to Database..")
	log.Println("migration successful")

	//running the application
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	r := routes.Routes(handlers.Repo)
	err := r.Run(portNumber)
	if err != nil {
		log.Fatal(err)

	}
}
