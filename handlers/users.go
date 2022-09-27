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

var Repo *Repository //used in the main func to run all handlers

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

// a Temporary func
func (r *Repository) Use(c *gin.Context) {
	c.HTML(http.StatusOK, "userprofile.html", &models.TemplateData{})
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
	firstname := c.Request.Form.Get("firstname")
	lastname := c.Request.Form.Get("lastname")
	password := c.Request.Form.Get("password")
	hashedpassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)

	form := forms.New(c.Request.Form)
	form.Required("email", "password", "firstname", "lastname")
	form.ValidEmail("email")
	form.MinLength("password", 8)
	if !form.Valid() {
		c.HTML(http.StatusMethodNotAllowed, "signup.html", &models.TemplateData{
			Form: form,
		})
		return
	}
	create := &models.ContentUser{
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
		Password:  hashedpassword,
	}
	c.ShouldBindJSON(&create)

	session, _ := helpers.GetCookieStore().Get(c.Request, "user")
	session.Values["user-signup"] = &create
	//Save must be called before writing to the response,
	// otherwise the session cookie will not be sent to the client.
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	//putting the post in the database(the Content_users table )
	if err := r.DB.Create(&create).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, "User Exists")
		helpers.SetFlash(c, "error", "User Exists")
		return
	}
	helpers.SetFlash(c, "message", "SignedUp Successfully")
	log.Println("Signed Up")
	c.Redirect(http.StatusSeeOther, "/user/content-home")

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
	user, err := database.Authenticate(r.DB, req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("incorrect email %s", req.Email)})
		c.Writer.Header().Set("Content-Type", "application/json")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		errors.New("incorrect password")
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("incorrect password %s", req.Password)})
		return
	}

	token, err := helpers.GenerateToken(email, true)
	if err != nil {
		helpers.SetFlash(c, "error", "Token not generated")
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println("token generated")
	_, err = json.Marshal(token)
	if err != nil {
		return
	}
	json.NewEncoder(c.Writer).Encode(token)
	c.Request.Header.Set("Token", token)
	c.Writer.Header().Set("Content-Type", "application/json")
	//helpers.SetFlash(c, "message", "logged in successfully")
	session, _ := helpers.GetCookieStore().Get(c.Request, "user")
	session.Values["email"] = req.Email
	session.Values["password"] = req.Password
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Redirect(http.StatusSeeOther, "/user/content-home")
	log.Println("logged in Successfully")

}

// viewing your profile when loged in
func (r *Repository) UserProfile(c *gin.Context) {

	var user models.ContentUser
	profile, err := helpers.GetCookieStore().Get(c.Request, "user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldnt Get session"})
	}
	datasession := make(map[string]interface{})
	datasession["user"] = profile
	DisplayProf := make(map[string]string)
	DisplayProf["email"] = user.Email
	DisplayProf["firstname"] = user.FirstName
	DisplayProf["lastname"] = user.LastName

	c.HTML(http.StatusOK, "user.html", &models.TemplateData{
		SessionData: datasession,
		User:        DisplayProf,
	})
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
