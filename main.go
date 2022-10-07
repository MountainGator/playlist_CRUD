package main

import (
	"context"
	"log"
	"net/http"

	"github.com/MountainGator/playlist_CRUD/controllers"
	"github.com/MountainGator/playlist_CRUD/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	us        services.UserService
	ps        services.PlaylistService
	uc        controllers.UserController
	pc        controllers.PlaylistController
	play_coll *mongo.Collection
	user_coll *mongo.Collection
	song_coll *mongo.Collection
	client    *mongo.Client
	err       error
	key       []byte
	store     *sessions.CookieStore
)

func init() {
	if err = godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("error connecting to mongo", err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal("error pinging mongo", err)
	}

	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)

	store.Options.HttpOnly = false
	store.Options.Secure = false

	play_coll = client.Database("playlist_db").Collection("playlist")
	user_coll = client.Database("playlist_db").Collection("users")
	song_coll = client.Database("playlist_db").Collection("songs")
	us = services.NewUserService(user_coll, store, context.TODO())
	ps = services.NewPlaylistService(user_coll, song_coll, context.TODO())
	uc = controllers.NewUserController(us)
	pc = controllers.NewPlayController(ps)
}

func Auth(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	_, ok := session.Values["user"]

	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"Error": "Not logged in"})
		c.Abort()
		return
	}
	c.Next()
}

func GetCreds(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	_, ok := session.Values["user"]

	if !ok {
		c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{"error": "Not logged in"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"user": ok})
}

func main() {
	r := gin.Default()
	user_router := r.Group("/user", Auth)
	play_router := r.Group("/playlist", Auth)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{
			"msg": "wazzup",
		})
	})
	r.GET("/check-login", GetCreds)
	r.POST("/login", uc.UserLogin)
	r.POST("/new-user", uc.CreateUser)
	r.GET("/find-playlist/:playlist", pc.FindPlaylist)
	r.GET("/songs", pc.GetSongs)
	user_router.PATCH("/update", uc.UpdateUser)
	user_router.DELETE("/delete", uc.DeleteUser)
	play_router.POST("/create", pc.NewPlaylist)
	play_router.GET("/get-all/:username", pc.GetAllPlaylist)
	play_router.POST("/add-song", pc.AddSong)
	play_router.PATCH("/update", pc.UpdatePlaylist)
	play_router.DELETE("/delete-song/:id", pc.DeleteSong)
	play_router.DELETE("/delete/:id", pc.DeletePlaylist)

	r.Use(cors.Default())
	r.Run()
}
