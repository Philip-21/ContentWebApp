package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Philip-21/Content/config"
	"github.com/Philip-21/Content/handlers"
	"github.com/Philip-21/Content/models"
	"github.com/alexedwards/scs/v2"

	"github.com/Philip-21/Content/database"
	"github.com/Philip-21/Content/routes"
)

const portNumber = ":8080"

var session *scs.SessionManager
var app config.AppConfig

func main() {

	//Seesions
	gob.Register(models.Content{})
	gob.Register(models.ContentUser{})
	gob.Register(map[string]int{})

	//app.InProduction = false
	app.Session = session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true                  //persists if the user should still leave page
	session.Cookie.SameSite = http.SameSiteLaxMode //applies to the site you want the cookie to apply to

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
