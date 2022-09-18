package routes

import (
	"html/template"
	"log"
	"time"

	"github.com/Philip-21/Content/database"
	"github.com/Philip-21/Content/handlers"
	"github.com/Philip-21/Content/middleware"
	"github.com/Philip-21/Content/models"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func Routes(app *handlers.Repository) *gin.Engine {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, //AllowCredentials indicates whether the request can include user credentials like cookies
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour, //MaxAge indicates how long (with second-precision) the results of a preflight request can be cached
	}))

	//loads the html file in the directory
	router.LoadHTMLGlob("templates/*.html")
	html := template.Must(template.ParseGlob("templates/*.html"))
	err := html.Execute(gin.DefaultWriter, &models.TemplateData{})
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
	router.GET("/content-home", api.ContentHome)
	router.GET("/signup", api.ShowSignup)
	router.GET("/login", api.ShowLogin)
	router.GET("/get-contents", api.GetContents)

	router.POST("/signup", api.Signup)
	router.POST("/login", api.Login)

	user := router.Group("/user")
	{
		user.Use(middleware.Auth())

		//user.GET("/info", api.UserID)

		user.POST("/post-content", api.PostCreateContent)
		user.PUT("/update-content/:id", api.UpdateContent)
		user.DELETE("/delete-content/:id", api.DeleteContent)
		router.GET("/get-content/:id", api.GetContentByID)

	}
	return router

}
