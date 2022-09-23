package controllers

import (
	"net/http"

	"github.com/MountainGator/playlist_CRUD/models"
	"github.com/MountainGator/playlist_CRUD/services"
	"github.com/gin-gonic/gin"
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
