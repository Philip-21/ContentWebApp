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

//the configuration for authentication and repository for User handlers
type Server struct {
	App        *config.AppConfig
	DB         *gorm.DB
	config     config.Envconfig
	store      *database.Userctx
	tokenMaker middleware.Maker
	router     *gin.Engine //initializing the router for the authentication
}

//initialize a new repository for users
var UserRepo *Server

func NewServer(config config.Envconfig, store *database.Userctx) (*Server, error) {
	tokenMaker, err := middleware.NewPasetoMaker("") //requires a symmetric key string
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err) //%w is used to wrap the original error.
	}
	//paramters for the new server function
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.router.Routes() //new routes initialized
	return server, nil
}

//Creating a User Account
func (r *Server) CreateUser(c *gin.Context) {
	create := models.ContentUser{
		Email:          c.PostForm("email"),
		HashedPassword: c.PostForm("pasword"),
	}
	err := c.BindJSON(&create)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//putting the post in the database(the Content table )

	if err := r.DB.Create(&create).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, create)
}

func (server *Server) LoginUser(c *gin.Context) {
	//the Login  request model
	var req models.LoginUserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//verifying the email
	user, err := server.store.GetUserID(req.Email, &gorm.DB{}) //gotten from the interface in the database
	if err != nil {
		fmt.Println("Invalid Credentials", user)
		c.JSON(http.StatusInternalServerError, user)
		return
	}
	//verifying the password
	passerr := config.CheckPassword(req.Password, user.HashedPassword) //matching the req password to the main ContentUser Password
	if passerr != nil {
		fmt.Println("Invalid Credentials", passerr)
		c.JSON(http.StatusNonAuthoritativeInfo, err)
		return
	}
	//creating  the token
	accessToken, err := server.tokenMaker.CreateToken(
		user.Email,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	rsp := models.LoginUserResponse{
		AccessToken: accessToken,
		User:        models.NewUserResponse(user), //the password wont be exposed to the client after loggin in
	}
	c.JSON(http.StatusOK, rsp)

}
