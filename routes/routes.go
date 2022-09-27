package routes

import (
	"bytes"
	"encoding/gob"
	"html/template"
	"log"

	"github.com/Philip-21/Content/database"
	"github.com/Philip-21/Content/handlers"
	"github.com/Philip-21/Content/middleware"
	"github.com/Philip-21/Content/models"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

var td *models.TemplateData

func AddData(td *models.TemplateData, c *gin.Context) *models.TemplateData {

	return td
}

func Routes(app *handlers.Repository) *gin.Engine {
	gob.Register(models.ContentUser{})
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https//*", "http://*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Access-Control-Request-Method", "Access-Control-Request-Headers", "Accept", "Authorization", " Accept-Encoding",
			"Content-Type", "Connection", " Host", "Origin", "User-Agent", "Referer", "Cache-Control", "X-header", "Token", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link", "Content-Length"},
		AllowCredentials: true, //AllowCredentials indicates whether the request can include user credentials like cookies

		MaxAge: 300, //MaxAge indicates how long (with second-precision) the results of a preflight request can be cached
	}))

	//loads the html file in the directory
	router.LoadHTMLGlob("templates/*.html")
	html := template.Must(template.ParseGlob("templates/*.html"))
	buf := new(bytes.Buffer)
	var w gin.ResponseWriter
	buf.WriteTo(w)
	td = AddData(td, &gin.Context{})
	err := html.Execute(buf, td)
	if err != nil {
		log.Println(err)
		log.Println("Cannot execute template")
	}
	//
	// _, err = buf.WriteTo(w)
	// if err != nil {
	// 	fmt.Println("Error writing template to browser", err)
	// }
	router.SetHTMLTemplate(html)
	//reads the images  kept in the static folder
	router.Static("static", "./static")

	//api* handlers.Repository a variable for Content Repository and User Repository
	api := &handlers.Repository{
		DB: database.GetDB(),
	}

	router.GET("/userprofile", api.Use)

	router.GET("/", api.Home)

	router.GET("/signup", api.ShowSignup)
	router.GET("/post/login", api.ShowLogin)
	router.GET("/get-contents", api.GetContents)

	router.POST("/signup", api.Signup)
	router.POST("/post/login", api.Login)

	user := router.Group("/user")
	{
		user.Use(middleware.Auth())

		//user.GET("/info", api.UserID)
		user.GET("/userprofile", api.Use)

		user.GET("/content-home", api.ContentHome)

		user.POST("/post-content", api.PostCreateContent)
		user.PUT("/update-content/:id", api.UpdateContent)
		user.DELETE("/delete-content/:id", api.DeleteContent)
		router.GET("/get-content/:id", api.GetContentByID)

	}
	return router

}
