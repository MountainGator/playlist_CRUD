package controllers

import (
	"net/http"

	"github.com/MountainGator/playlist_CRUD/models"
	"github.com/MountainGator/playlist_CRUD/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Pwd), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Pwd = string(bytes)

	err = uc.UserService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) UserLogin(c *gin.Context) {

	var user *models.LoginInfo

	c.BindJSON(&user)

	err := uc.UserService.UserLogin(&user.Username, user.Password, c)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"invalid password": err.Error()})
	}

	c.JSON(http.StatusAccepted, gin.H{"success": "logged in"})

}

func (uc *UserController) Logout(c *gin.Context) {
}

func (uc *UserController) UpdateUser(c *gin.Context) {

}

func (uc *UserController) DeleteUser(c *gin.Context) {

}
