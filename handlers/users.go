package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Philip-21/proj1/database"
	"github.com/Philip-21/proj1/forms"
	"github.com/Philip-21/proj1/middleware"
	"github.com/Philip-21/proj1/models"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"golang.org/x/crypto/bcrypt"
)

//---------------initialize a new repository for users----------------

var Repo *Repository

func (r *Repository) ShowSignup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", &models.TemplateData{
		Form: forms.New(nil), //creating an empty form

	})
}

func (r *Repository) ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", &models.TemplateData{
		IsAuthenticated: 1,
	})
}

//Creating a User Account
func (r *Repository) Signup(c *gin.Context) {
	//preventing session fixation attack by renewing the token
	//_ = r.App.Session.RenewToken(c.Request.Context())
	err := c.Request.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	email := c.Request.Form.Get("email")
	password := c.Request.Form.Get("password")
	hashedpassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	form := forms.New(c.Request.Form)
	form.Required("email", "password")
	form.ValidEmail("email")
	form.MinLength("password", 8)
	if !form.Valid() {
		c.HTML(http.StatusMethodNotAllowed, "signup.html", &models.TemplateData{
			Form: form,
		})
		return
	}
	create := &models.ContentUser{
		Email:    email,
		Password: hashedpassword,
	}
	// err = c.ShouldBindJSON(&create)
	// if err != nil {
	// 	//c.JSON(http.StatusInternalServerError, create)
	// 	return
	// }
	//putting the post in the database(the Content_users table )
	if err := r.DB.Create(&create).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, "User Exists")
		return
	}
	r.App.Session.Put(c.Request.Context(), "email", create)
	r.App.Session.Put(c.Request.Context(), "flash", "signed in successfuly")
	c.Redirect(http.StatusSeeOther, "/")
	c.JSON(http.StatusOK, create)

}

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
	c.String(200, csrf.GetToken(c))
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
