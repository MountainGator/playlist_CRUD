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

type PlaylistController struct {
	PlaylistService services.PlaylistService
}

func NewUserController(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

func NewPlayController(playservice services.PlaylistService) PlaylistController {
	return PlaylistController{
		PlaylistService: playservice,
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

}

func (uc *UserController) UpdateUser(c *gin.Context) {

}

func (uc *UserController) DeleteUser(c *gin.Context) {

}

func (pc *PlaylistController) NewPlaylist(c *gin.Context) {

}

func (pc *PlaylistController) FindPlaylist(c *gin.Context) {

}

func (pc *PlaylistController) AddSong(c *gin.Context) {

}
func (pc *PlaylistController) DeleteSong(c *gin.Context) {

}
func (pc *PlaylistController) DeletePlaylist(c *gin.Context) {

}
