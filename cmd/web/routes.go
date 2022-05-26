package main

import (
	"net/http"

	"github.com/Philip-21/proj1/config"
	"github.com/Philip-21/proj1/handlers"
	"github.com/gin-gonic/gin"
)

func Route(app *config.AppConfig) http.Handler {
	router := gin.Default()
	router.GET("/get-all-contents", handlers.Repo.GetContent)
	router.GET("/get-contents/:id", handlers.Repo.GetContentByID)
	router.POST("/create-content", handlers.Repo.CreateContent)
	router.DELETE("/delete-content/:id", handlers.Repo.DeleteContent)
	router.PUT("/update-content", handlers.Repo.UpdateContent)

	return router
}
