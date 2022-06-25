package middleware

import (
	"fmt"
	"time"

	"github.com/Philip-21/proj1/config"
	"github.com/Philip-21/proj1/database"
	"github.com/gin-gonic/gin"
)

// a token Maker interface to manage the creation and verification of the Paseto tokens
type Maker interface {
	CreateToken(email string, duration time.Duration) (string, error) //create and sign a new token for a specific username and valid duration.
	VerifyToken(token string) (*Payload, error)                       //checks if the input token is valid or not.
}

//the repo for configuring the token generation
type TokenServer struct {
	config     config.Envconfig
	store      *database.Userctx
	inter      database.AuthUser
	tokenMaker Maker
	router     *gin.Engine
}

//TokenRepo will be used in the main function when parsing it as the  parameters for  Routes
var TokenRepo *TokenServer

func NewServer(config config.Envconfig, store *database.Userctx) (*TokenServer, error) {
	tokenMaker, err := NewPasetoMaker("") //requires a symmetric key string
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err) //%w is used to wrap the original error.
	}
	//paramters for the new server function
	server := &TokenServer{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.router.Routes() //new routes initialized
	return server, nil
}
