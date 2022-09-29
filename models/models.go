package models

type User struct {
	Name     string `json:"name" bson:"name"`
	Age      int    `json:"age" bson:"age"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Pwd      string `json:"password" bson:"password"`
}

type Song struct {
	Artist string `json:"artist" bson:"artist"`
	Title  string `json:"title" bson:"title"`
	Album  string `json:"album" bson:"album"`
}

type Playlist struct {
	Creator       string  `json:"creator" bson:"creator"`
	Playlist_name string  `json:"playlist_name" bson:"playlist_name"`
	Songs         []*Song `json:"songs" bson:"songs"`
}
