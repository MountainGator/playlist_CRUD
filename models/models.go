package models

type User struct {
	Id       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Pwd      string `json:"password" bson:"password"`
}

type Song struct {
	Id     string `json:"id" bson:"_id"`
	Artist string `json:"artist" bson:"artist"`
	Title  string `json:"title" bson:"title"`
	Album  string `json:"album" bson:"album"`
}

type Playlist struct {
	Id            string  `json:"id" bson:"_id"`
	Creator_Id    string  `json:"creatorid" bson:"creator_id"`
	Playlist_name string  `json:"playlist_name" bson:"playlist_name"`
	Songs         []*Song `json:"songs" bson:"songs"`
}

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DeleteSong struct {
	SongId     string `json:"songId"`
	PlaylistId string `json:"playlistId"`
}
