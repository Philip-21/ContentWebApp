package routes

import (
	"github.com/Philip-21/proj1/database"
	"github.com/Philip-21/proj1/handlers"
	"github.com/Philip-21/proj1/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(app *middleware.TokenServer) *gin.Engine {
	router := gin.Default()

	//con* handlers.Server a variable for Content Repository and User Repository
	api := &handlers.Repository{
		DB: database.GetDB(),
	}

	router.GET("/get-contents", api.GetContent)
	router.GET("/get-content/:id", api.GetContentByID)
	router.POST("/signup", api.CreateUser)
	router.POST("login", api.LoginUser)

	user := router.Group("/user")
	{
		user.Use(middleware.AuthMiddleware(&middleware.PasetoMaker{}))

		user.POST("/signup", api.CreateUser)
		user.POST("/login", api.LoginUser)
		user.POST("/post-content", api.CreateContent)
		user.PUT("/update-content/:id", api.UpdateContent)
		user.DELETE("/delete-content/:id", api.DeleteContent)

	}
	return router

}
