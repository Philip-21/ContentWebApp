package handlers

import (
	"fmt"
	"net/http"

	"github.com/Philip-21/proj1/config"
	"github.com/Philip-21/proj1/database"
	"github.com/Philip-21/proj1/models"
	"github.com/Philip-21/proj1/token"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//the configuration for authentication and repository for User handlers
type Server struct {
	App        *config.AppConfig
	DB         *gorm.DB
	config     config.Envconfig
	store      *database.Userctx
	tokenMaker token.Maker
	router     *gin.Engine //initializing the router for the authentication
}

func NewServer(config config.Envconfig, store *database.Userctx) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker("") //requires a symmetric key string
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
func (r *Server) CreateUser(c *gin.Context) (models.ContentUser, error) {
	create := models.ContentUser{}
	err := c.BindJSON(&create)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return create, nil
	}
	//putting the post in the database(the Content table )
	if err := r.DB.Create(&create).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return create, nil
	}
	c.JSON(http.StatusOK, create)
	return create, nil
}

// type CreateUserinterface interface {
// 	CreateUser(c *gin.Context) (models.User, error)
// }

func (server *Server) loginUser(c *gin.Context) {
	//var rt database.Userctx
	var req models.LoginUserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//getting the email
	user, err := server.store.GetUserID(req.Email, &gorm.DB{}) //gotten from the interface in the database
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	new, err := config.HashPassword(user.HashedPassword)

}
