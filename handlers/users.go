package handlers

import (
	"fmt"
	"net/http"

	"github.com/Philip-21/proj1/config"
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
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	create := models.ContentUser{
		Email:          req.Email,
		HashedPassword: string(passwordHash),
	}

	json := c.BindJSON(&create)
	if json != nil {
		c.JSON(http.StatusInternalServerError, json)
		return
	}
	//putting the post in the database(the Content table )
	if err := r.DB.Create(&create).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, create)
}

func (r *Repository) LoginUser(c *gin.Context) {
	user, err := database.GetUser(r.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	//creating  the token
	accessToken, err := r.tokenMaker.CreateToken(
		user.Email,
		r.config.AccessTokenDuration,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	rsp := models.LoginUserResponse{
		AccessToken: accessToken,
		User:        models.NewUserResponse(user), //the password wont be exposed to the client after loggin in
	}
	c.JSON(http.StatusOK, "logged in succesfully")
	c.JSON(http.StatusOK, rsp)

}

func GenerateToken(config config.Envconfig, store *database.AuthUser) (*Repository, error) {
	tokenMaker, err := middleware.NewPasetoMaker("")
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err) //%w is used to wrap the original error.
	}
	server := &Repository{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.router.Routes()
	return server, nil
}

var Repo *Repository
