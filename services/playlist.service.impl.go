package services

import (
	"context"
	// "errors"
	"fmt"

	"github.com/MountainGator/playlist_CRUD/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlayServiceImpl struct {
	playlistcollection *mongo.Collection
	songcollection     *mongo.Collection
	ctx                context.Context
}

func NewPlaylistService(playlist_collection *mongo.Collection, song_collection *mongo.Collection, ctx context.Context) PlaylistService {
	return &PlayServiceImpl{
		playlistcollection: playlist_collection,
		songcollection:     song_collection,
		ctx:                ctx,
	}
}

func (ps *PlayServiceImpl) NewPlaylist(playlist *models.Playlist) error {
	_, err := ps.playlistcollection.InsertOne(ps.ctx, playlist)
	return err
}

func (ps *PlayServiceImpl) FindPlaylist(name *string) (*models.Playlist, error) {
	var playlist *models.Playlist
	query := bson.D{bson.E{Key: "playlist_name", Value: name}}
	err := ps.playlistcollection.FindOne(ps.ctx, query).Decode(&playlist)
	if err != nil {
		fmt.Println("could not find playlist")
		return nil, err
	}
	return playlist, nil
}

func (ps *PlayServiceImpl) AddSong(song *models.Song, playlist_name *string) error {
	_, err := ps.songcollection.InsertOne(ps.ctx, song)
	if err != nil {
		return err
	}
	er := Add_to_playlist(song, playlist_name, ps.ctx, ps)
	if err != nil {
		return er
	}
	return nil

}
func Add_to_playlist(song *models.Song, playlist_name *string, ctx context.Context, ps *PlayServiceImpl) error {
	var new_song *models.Song
	song_q := bson.D{bson.E{Key: "artist", Value: song.Artist}, bson.E{Key: "title", Value: song.Title}}
	err := ps.songcollection.FindOne(ps.ctx, song_q).Decode(&new_song)
	if err != nil {
		return err
	}

	var playlist *models.Playlist
	query := bson.D{bson.E{Key: "playlist_name", Value: playlist_name}}
	er := ps.playlistcollection.FindOne(ps.ctx, query).Decode(&playlist)
	if er != nil {
		return er
	}
	playlist.Songs = append(playlist.Songs, new_song)
	update := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{Key: "songs", Value: playlist.Songs},
			},
		},
	}

	_, e := ps.playlistcollection.UpdateOne(ctx, query, update)
	return e
}

func (ps *PlayServiceImpl) DeleteSong(name *string, artist *string, playlist_name *string) error {
	var song *models.Song
	song_query := bson.D{bson.E{Key: "title", Value: name}, bson.E{Key: "artist", Value: artist}}
	err := ps.songcollection.FindOne(ps.ctx, song_query).Decode(&song)

	if err != nil {
		return err
	}

	var playlist *models.Playlist
	var n00b []*models.Song
	query := bson.D{bson.E{Key: "playlist_name", Value: name}}
	e := ps.playlistcollection.FindOne(ps.ctx, query).Decode(&playlist)

	if e != nil {
		return e
	}

	for _, each := range playlist.Songs {
		if each.Title != song.Title && each.Artist != song.Artist {
			n00b = append(n00b, each)
		}
	}

	update := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{Key: "songs", Value: n00b},
			},
		},
	}

	_, new_err := ps.playlistcollection.UpdateOne(ps.ctx, query, update)

	if new_err != nil {
		return new_err
	}

	_, er := ps.songcollection.DeleteOne(ps.ctx, song_query)

	if er != nil {
		return er
	}

	return nil

}
func (ps *PlayServiceImpl) DeletePlaylist(name *string) error {
	query := bson.D{bson.E{Key: "playlist_name", Value: name}}
	_, err := ps.playlistcollection.DeleteOne(ps.ctx, query)

	if err != nil {
		return err
	}

	return nil
}