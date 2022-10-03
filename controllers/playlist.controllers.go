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

	c.JSON(http.StatusAccepted, gin.H{"Success": "Playlist created"})
}

func (pc *PlaylistController) UpdatePlaylist(c *gin.Context) {

}

func (pc *PlaylistController) FindPlaylist(c *gin.Context) {
	var (
		playlist_name string
		fun           string
		play_List     []*models.Playlist
		err           error
	)

	playlist_name = c.Param("playlist")
	fun = c.Param("fun")

	play_List, err = pc.PlaylistService.FindPlaylist(playlist_name, fun)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error finding playlist": err})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"playlist data": play_List})

}

func (pc *PlaylistController) GetSongs(c *gin.Context) {
	song_list, err := pc.PlaylistService.GetSongs()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Get Songs Error": err})
	}

	c.JSON(http.StatusAccepted, gin.H{"data": song_list})

}
func (pc *PlaylistController) AddSong(c *gin.Context) {

}
func (pc *PlaylistController) DeleteSong(c *gin.Context) {

}
func (pc *PlaylistController) DeletePlaylist(c *gin.Context) {

}
