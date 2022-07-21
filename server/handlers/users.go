package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Philip-21/proj1/database"
	"github.com/Philip-21/proj1/middleware"
	"github.com/Philip-21/proj1/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//---------------initialize a new repository for users----------------

//Creating a User Account
func (r *Repository) CreateUser(c *gin.Context) {

	var req models.SigninUserRequest
	//var data map[string]string

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 8)
	if err != nil {
		return
	}
	create := &models.ContentUser{
		Email:    req.Email,
		Password: hashedpassword,
	}
	json := c.BindJSON(&create)
	if json != nil {
		c.JSON(http.StatusInternalServerError, json)
		return
	}
	//putting the post in the database(the Content_users table )
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
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect parameters"})
		return
	}
	user, err := database.GetUser(r.DB, req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("incorrect email %s", req.Email)})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		errors.New("incorrect password")
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("incorrect password %s", req.Password)})
		return
	}

	token, err := middleware.GenerateJwt(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
	c.JSON(http.StatusOK, "logged in successfully")
}

func (r *Repository) UserID(c *gin.Context) {
	id, _, ok := middleware.GetSession(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}
	user, err := database.UserID(r.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, user)
}
