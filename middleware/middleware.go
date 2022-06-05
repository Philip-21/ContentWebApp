package middleware

import (
	"fmt"

	"github.com/Philip-21/proj1/config"
	"github.com/Philip-21/proj1/models"
	"github.com/Philip-21/proj1/token"
	"github.com/gin-gonic/gin"
)

//
type Server struct {
	config     config.Envconfig
	store      *models.User
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config config.Envconfig, store *models.User) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker("") //requires a symmetric key string
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err) //%w is used to wrap the original error.
	}
	//paramters for the new server function
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
