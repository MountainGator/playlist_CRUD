package controllers

import (
	"net/http"

	"github.com/MountainGator/playlist_CRUD/models"
	"github.com/MountainGator/playlist_CRUD/services"
	"github.com/gin-gonic/gin"
	// "golang.org/x/crypto/bcrypt"
)

type PlaylistController struct {
	PlaylistService services.PlaylistService
}

func NewPlayController(playservice services.PlaylistService) PlaylistController {
	return PlaylistController{
		PlaylistService: playservice,
	}
}

func (pc *PlaylistController) NewPlaylist(c *gin.Context) {
	var play *models.Playlist

	c.BindJSON(&play)

	if err := pc.PlaylistService.NewPlaylist(play); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (pc *PlaylistController) FindPlaylist(c *gin.Context) {
}

func (pc *PlaylistController) AddSong(c *gin.Context) {

}
func (pc *PlaylistController) DeleteSong(c *gin.Context) {

}
func (pc *PlaylistController) DeletePlaylist(c *gin.Context) {

}
