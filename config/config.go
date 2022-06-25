package config

import (
	"log"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type AppConfig struct {
	//TemplateCache map[string]*template//
	InfoLog  *log.Logger //a standard library that writes information to log files
	ErrorLog *log.Logger //this allows us to write logs and handle errors efficiently
	Session  *scs.SessionManager
	router   *gin.Engine
}

type Envconfig struct {
	Host                string        `mapstructure:"DB_HOST"`
	Port                string        `mapstructure:"DB_PORT"`
	Password            string        `mapstructure:"DB_PASSWORD"`
	User                string        `mapstructure:"DB_USER"`
	DBName              string        `mapstructure:"DB_NAME"`
	SSLMode             string        `mapstructure:"DB_SSLMODE"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"` //the size of this symmetric key should be exactly 32 bytes.
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

//conf variable will be accesed by other files and packages within the app
var Conf *Envconfig

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigFile("app.env")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		log.Fatal(err)
	}
}
