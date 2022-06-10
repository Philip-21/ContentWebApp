package main

import (
	"fmt"
	"testing"

	"github.com/Philip-21/proj1/config"
	"github.com/gin-gonic/gin"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	router := Route(&app)

	switch v := router.(type) {
	case *gin.Engine:
		// do nothing; test passed
	default:
		t.Error(fmt.Sprintf("type is not *gin.engines, type is %T", v))
	}

}
