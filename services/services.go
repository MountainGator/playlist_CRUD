package services

import (
	"github.com/MountainGator/playlist_CRUD/models"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	CreateUser(*models.User) error
	UserLogin(*string, string, *gin.Context) error
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}

type PlaylistService interface {
	NewPlaylist(*models.Playlist) error
	FindPlaylist(*string) (*models.Playlist, error)
	AddSong(*models.Song, *string) error
	DeleteSong(*string, *string, *string) error
	DeletePlaylist(*string) error
}