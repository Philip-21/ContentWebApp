package handlers

import (
	"fmt"
	"net/http"

	"github.com/Philip-21/proj1/config"
	"github.com/Philip-21/proj1/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//initialize a new repository for users

//Creating a User Account
func (r *Repository) CreateUser(c *gin.Context) {
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

func (server *Repository) LoginUser(c *gin.Context) {
	//the Login  request model
	var req models.LoginUserRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//verifying the email
	user, err := server.store.GetUserEmail(req.Email, &gorm.DB{}) //gotten from the interface in the database
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
