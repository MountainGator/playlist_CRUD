package services

import (
	"github.com/MountainGator/playlist_CRUD/models"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	CreateUser(*models.User, *gin.Context) error
	UserLogin(*string, string, *gin.Context) error
	GetUserDetails(*string) (*models.User, error)
	Logout(*gin.Context) error
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}

type PlaylistService interface {
	NewPlaylist(*models.Playlist) error
	FindPlaylist(string) ([]*models.Playlist, error)
	GetAllPlaylist(*string) ([]*models.Playlist, error)
	GetSongs() ([]*models.Song, error)
	AddSong(*models.Song) error
	UpdatePlaylist(*models.Playlist) error
	DeleteSong(*string) error
	DeletePlaylist(*string) error
}
