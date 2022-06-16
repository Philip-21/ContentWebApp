package main

import (
	"fmt"
	"log"

	"github.com/Philip-21/proj1/config"

	"github.com/Philip-21/proj1/database"
	"github.com/Philip-21/proj1/router"
)

const portNumber = ":8080"

func main() {

	//DATABASE
	config.LoadConfig() //load viper configurations
	log.Println("Connecting to database...")

	database.Initdb(config.Conf) //calling the db function

	log.Println("Connected to Database..")
	log.Println("migration successful")

	//running the application
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	r := router.Routes()
	err := r.Run(portNumber)
	if err != nil {
		log.Fatal(err)

	}
}
