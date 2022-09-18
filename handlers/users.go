package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Philip-21/Content/database"
	"github.com/Philip-21/Content/forms"
	"github.com/Philip-21/Content/helpers"
	"github.com/Philip-21/Content/models"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"golang.org/x/crypto/bcrypt"
)

//---------------initialize a new repository for users----------------

var Repo *Repository

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (r *Repository) ShowSignup(c *gin.Context) {
	data := make(map[string]interface{})
	data["messages"] = helpers.GetFlash(c, "message")

	res := make(map[string]interface{})
	res["errors"] = helpers.GetFlash(c, "error")

	c.HTML(http.StatusOK, "signup.html", &models.TemplateData{
		Form:    forms.New(nil), //creating an empty form
		Message: data,
		Error:   res,
	})

}

func (r *Repository) ShowLogin(c *gin.Context) {
	data := make(map[string]interface{})
	data["messages"] = helpers.GetFlash(c, "message")

	res := make(map[string]interface{})
	res["errors"] = helpers.GetFlash(c, "error")

	c.HTML(http.StatusOK, "login.html", &models.TemplateData{
		Form:    forms.New(nil),
		Message: data,
		Error:   res,
	})

}

var Secret = os.Getenv("SESSION_KEY")

// Creating a User Account
func (r *Repository) Signup(c *gin.Context) {

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
	c.ShouldBindJSON(&create)
	session, _ := helpers.GetCookieStore().Get(c.Request, "session-cookie")
	session.Values["user"] = &create
	session.Save(c.Request, c.Writer)
	//putting the post in the database(the Content_users table )
	if err := r.DB.Create(&create).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, "User Exists")
		helpers.SetFlash(c, "error", "User Exists")
		return
	}
	helpers.SetFlash(c, "message", "SignedUp Successfully")
	log.Println("Signed Up")
	c.Redirect(http.StatusSeeOther, "/")

}

func (r *Repository) Login(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		log.Println((err))
		return
	}
	email := c.Request.Form.Get("email")
	password := c.Request.Form.Get("password")
	form := forms.New(c.Request.Form)
	form.Required("email", "password")
	form.ValidEmail("email")
	if !form.Valid() {
		c.HTML(http.StatusMethodNotAllowed, "login.html", &models.TemplateData{
			Form: form,
		})
		return
	}
	req := &models.SignInResponse{
		Email:    email,
		Password: password,
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

	token, _, err := helpers.GenerateToken(user.Email)
	if err != nil {
		helpers.SetFlash(c, "error", "Token not generated")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = json.Marshal(token)
	if err != nil {
		return
	}
	log.Println("token generated")
	helpers.SetFlash(c, "message", "logged in successfully")
	log.Println("logged in Successfully")
	c.Writer.Header()
	c.Redirect(http.StatusSeeOther, "/content-home")

}
func (r *Repository) LogOut() {}

// func (r *Repository) UserID(c *gin.Context) {
// 	id, _, ok := middleware.GetSession(c)
// 	if !ok {
// 		c.JSON(http.StatusUnauthorized, gin.H{})
// 		return
// 	}
// 	user, err := database.UserID(r.DB, id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{})
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// }
