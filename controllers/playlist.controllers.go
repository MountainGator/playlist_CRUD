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
	var playlist *models.Playlist
	var err error
	if err = c.BindJSON(&playlist); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err = pc.PlaylistService.UpdatePlaylist(playlist); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"Success": "playlist updated"})
}

func (pc *PlaylistController) FindPlaylist(c *gin.Context) {
	var (
		playlist_name string
		play_List     []*models.Playlist
		err           error
	)

	playlist_name = c.Param("playlist")

	play_List, err = pc.PlaylistService.FindPlaylist(playlist_name)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error finding playlist": err})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"playlist data": play_List})

}
func (pc *PlaylistController) GetAllPlaylist(c *gin.Context) {

}
func (pc *PlaylistController) GetSongs(c *gin.Context) {
	song_list, err := pc.PlaylistService.GetSongs()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Get Songs Error": err})
	}

	c.JSON(http.StatusAccepted, gin.H{"data": song_list})

}
func (pc *PlaylistController) AddSong(c *gin.Context) {
	var song *models.Song

	c.BindJSON(&song)

	if err := pc.PlaylistService.AddSong(song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": "added song"})

}

func (pc *PlaylistController) DeleteSong(c *gin.Context) {
	song := c.Param("id")

	c.BindJSON(&song)

	if err := pc.PlaylistService.DeleteSong(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.JSON(http.StatusAccepted, gin.H{"sucess": "deleted song"})

}
func (pc *PlaylistController) DeletePlaylist(c *gin.Context) {
	id := c.Param("id")

	if err := pc.PlaylistService.DeletePlaylist(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.JSON(http.StatusAccepted, gin.H{"sucess": "deleted playlist"})

}
