package services

import (
	"github.com/MountainGator/playlist_CRUD/models"
)

type UserService interface {
	CreateUser(*models.User) error
	UserLogin(*string) (*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}

type PlaylistService interface {
	NewPlaylist(*models.Playlist) error
	FindPlaylist(*string) error
	AddSong(*string, *models.Song) error
	DeleteSong(*string, *string) error
	DeletePlaylist(*string) error
}
