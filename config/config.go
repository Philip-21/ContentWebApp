package config

import "time"

type AppConfig struct {
	//TemplateCache map[string]*template//

}

type PostgresConfig struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

//the configuration for the paseto token
type TokenConfig struct {
	TokenSymmetricKey   string //the size of this symmetric key should be exactly 32 bytes.
	AccessTokenDuration time.Duration
}
