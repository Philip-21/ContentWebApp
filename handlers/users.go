package handlers

import (
	"net/http"

	"github.com/Philip-21/proj1/database"
	"github.com/Philip-21/proj1/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//---------------initialize a new repository for users----------------

//Creating a User Account
func (r *Repository) CreateUser(c *gin.Context) {

	var req models.SigninUserRequest
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	create := models.ContentUser{
		Email:    req.Email,
		Password: string(passwordHash),
	}

	json := c.BindJSON(&create)
	if json != nil {
		c.JSON(http.StatusInternalServerError, json)
		return
	}
	//putting the post in the database(the Content table )
	if err := r.DB.Create(&create).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, "User Exists")
		return
	}
	c.JSON(http.StatusOK, create)
}

var Repo *Repository

func (r *Repository) Login(c *gin.Context) {
	var req models.SigninUserRequest

	user, err := database.GetUser(r.DB, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Invalid credentials")
		c.JSON(http.StatusInternalServerError, user)

	}
	if user.Password != req.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "incorrect password",
		})
		return
	}

	c.JSON(http.StatusOK, "logged in successfully")
}
