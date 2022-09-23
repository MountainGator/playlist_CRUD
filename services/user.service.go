package services

import (
	"github.com/MountainGator/playlist_CRUD/models"
)

type UserService interface {
	CreateUser(*models.User) error
	Login(*string) (*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}

type PlaylistService interface {
	NewPlaylist(*models.Playlist) error
	AddSong(*models.Song) error
	DeleteSong(*string) error
	DeletePlaylist(*string) error
}
