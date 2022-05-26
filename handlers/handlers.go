package handlers

import (
	"fmt"
	"net/http"

	"github.com/Philip-21/proj1/config"
	"github.com/Philip-21/proj1/database"
	"github.com/Philip-21/proj1/middleware"
	"github.com/Philip-21/proj1/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//repository for the databse and the application Contents
type Repository struct {
	App *config.AppConfig
	DB  *gorm.DB
}

//initializing a new repository
var Repo *Repository

func (r *Repository) CreateContent(c *gin.Context) {

	post := models.Content{}
	//BindJSON passes the 400 status code to the context then returns a pointer or an error
	err := c.BindJSON(&post)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//putting the post in the database(the Content table )
	if err := r.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, post)

}

func (r *Repository) GetContent(c *gin.Context) {
	con, err := database.GetContents(r.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, con)
}

func (r *Repository) GetContentByID(c *gin.Context) {
	id := c.Params.ByName("id")
	content, exist, err := database.GetContentByID(id, r.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	if !exist {
		c.JSON(http.StatusNotFound, "there is no Content in database")
		return
	}

	c.JSON(http.StatusOK, content)
}

func (r *Repository) DeleteContent(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := database.GetContentByID(id, r.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "record not exists")
		return
	}

	err = database.DeleteContent(id, r.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "record deleted successfully")
}

func (r *Repository) UpdateContent(c *gin.Context) {
	id := c.Params.ByName("id")
	//getting the content to update
	_, exists, err := database.GetContentByID(id, r.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "record not exists")
		return
	}
	//updating the content according to the model format
	updatedContent := models.Content{}
	//BindJSON passes the 400 status code to the context then returns a pointer or an error
	err = c.BindJSON(&updatedContent)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//putting the updated content in the database
	if err := database.UpdateContent(r.DB, &updatedContent); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	r.GetContent(c)
}

////-----users------///

func (server *middleware.Server) CreateUser(ctx *gin.Context) {

	user, err := server.store.CreateUser(ctx, arg)

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)

}
