package routes

import (
	"html/template"
	"log"
	"os"

	"github.com/Philip-21/Content/config"
	"github.com/Philip-21/Content/database"
	"github.com/Philip-21/Content/handlers"
	"github.com/Philip-21/Content/middleware"
	"github.com/Philip-21/Content/render"
	"github.com/gin-gonic/gin"
)

func Routes(app *config.AppConfig) *gin.Engine {
	router := gin.Default()

	//loads the html file in the directory
	//router.LoadHTMLGlob("templates/*.html")
	html := template.Must(template.ParseGlob("templates/*.html"))
	err := html.Execute(os.Stdout, render.AddData)
	if err != nil {
		log.Println(err)
		log.Println("Cannot execute template")
	}
	router.SetHTMLTemplate(html)
	//reads the images  kept in the static folder
	router.Static("static", "./static")

	//api* handlers.Repository a variable for Content Repository and User Repository
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
		user.Use(middleware.Auth())
		//user.GET("/info", api.UserID)
		user.POST("/post-content", api.CreateContent)
		user.PUT("/update-content/:id", api.UpdateContent)
		user.DELETE("/delete-content/:id", api.DeleteContent)

	}
	return router

}
