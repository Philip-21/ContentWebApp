package middleware

import (
	"fmt"

	"github.com/Philip-21/proj1/config"
	"github.com/Philip-21/proj1/token"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//
type Server struct {
	config     config.TokenConfig
	store      *gorm.DB
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config config.TokenConfig, store *gorm.DB) (*Server, error) {
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
