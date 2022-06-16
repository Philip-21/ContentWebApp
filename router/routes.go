package router

import (
	"github.com/Philip-21/proj1/database"
	"github.com/Philip-21/proj1/handlers"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()
	//api*handlers.Sever a variable for User Repository
	api := &handlers.Server{
		DB: database.GetDB(),
	}
	//con* handlers.Server a variable for Content Repository
	con := &handlers.Repository{
		DB: database.GetDB(),
	}

	router.GET("/get-contents", con.GetContent)
	router.GET("/get-content/:id", con.GetContentByID)

	router.POST("/post-content", con.CreateContent)
	router.PUT("/update-content/:id", con.UpdateContent)
	router.DELETE("/delete-content/:id", con.DeleteContent)

	router.POST("/signup", api.CreateUser)
	router.POST("/login", api.LoginUser)

	//user := router.Group("/user")
	//user.Use(middleware.Auth())

	{
		// user.POST("/signup", handlers.UserRepo.CreateUser)
		// user.POST("/login", handlers.UserRepo.LoginUser)
		// user.POST("/create-content", handlers.Repo.CreateContent)
		// user.DELETE("/delete-content/:id", handlers.Repo.DeleteContent)
		// user.PUT("/update-content", handlers.Repo.UpdateContent)
	}

	return router

}
