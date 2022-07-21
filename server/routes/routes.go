package routes

import (
	"github.com/Philip-21/proj1/database"
	"github.com/Philip-21/proj1/handlers"
	"github.com/Philip-21/proj1/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(app *handlers.Repository) *gin.Engine {
	router := gin.Default()

	//api* handlers.Server a variable for Content Repository and User Repository
	api := &handlers.Repository{
		DB: database.GetDB(),
	}

	router.GET("/get-contents", api.GetContent)
	router.GET("/get-content/:id", api.GetContentByID)
	router.POST("/signup", api.CreateUser)
	router.POST("login", api.Login)

	user := router.Group("/user")
	{
		user.Use(middleware.Auth)
		user.GET("/info", api.UserID)
		user.POST("/signup", api.CreateUser)
		user.POST("/login", api.Login)
		user.POST("/post-content", api.CreateContent)
		user.PUT("/update-content/:id", api.UpdateContent)
		user.DELETE("/delete-content/:id", api.DeleteContent)

	}
	return router

}
