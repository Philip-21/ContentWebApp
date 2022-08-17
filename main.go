package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Philip-21/proj1/config"
	"github.com/Philip-21/proj1/handlers"
	"github.com/Philip-21/proj1/models"
	"github.com/alexedwards/scs/v2"

	"github.com/Philip-21/proj1/database"
	"github.com/Philip-21/proj1/routes"
)

const portNumber = ":8080"

var session *scs.SessionManager
var app config.AppConfig

func main() {

	gob.Register(models.Content{})
	gob.Register(models.ContentUser{})
	gob.Register(map[string]int{})

	//app.InProduction = false
	app.Session = session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	//session.Cookie.Secure = app.InProduction

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
