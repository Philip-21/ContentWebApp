package routes

import (
	"github.com/Philip-21/proj1/database"
	"github.com/Philip-21/proj1/handlers"
	"github.com/Philip-21/proj1/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(app *handlers.Repository) *gin.Engine {
	router := gin.Default()

	//loads the html file in the directory
	router.LoadHTMLGlob("templates/*.html")

	//reads the images  kept in the static folder
	router.Static("static", "./static")

	//api* handlers.Server a variable for Content Repository and User Repository
	api := &handlers.Repository{
		DB: database.GetDB(),
	}

	router.GET("/", api.Home)
	router.GET("signup", api.ShowSignup)
	router.GET("/login", api.ShowLogin)
	router.GET("/get-contents", api.GetContent)
	router.GET("/get-content/:id", api.GetContentByID)

	router.POST("/signup", api.Signup)
	router.POST("/login", api.Login)
	router.POST("/post-content", api.CreateContent)

	user := router.Group("/user")
	{
		user.Use(middleware.Auth)
		user.GET("/info", api.UserID)
		user.POST("/signup", api.Signup)
		user.POST("/login", api.Login)
		user.POST("/post-content", api.CreateContent)
		user.PUT("/update-content/:id", api.UpdateContent)
		user.DELETE("/delete-content/:id", api.DeleteContent)

	}
	return router

}
