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

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	us     services.UserService
	ps     services.PlaylistService
	uc     controllers.UserController
	pc     controllers.PlaylistController
	coll   *mongo.Collection
	client *mongo.Client
	err    error
)

func init() {
	if err := godotenv.Load(); err != nil {
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

	coll = client.Database("playlist_db").Collection("playlist")
	us = services.NewUserSerice(coll, context.TODO())
	uc = controllers.NewUserController(us)
	pc = controllers.NewPlayController(ps)
}

func main() {
	r := gin.Default()

	if err != nil {
		panic(err)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{
			"msg": "wazzup",
		})
	})

	r.POST("/login", uc.UserLogin)
	r.POST("/create-playlist", pc.NewPlaylist)
	r.GET("/find-playlist", pc.FindPlaylist)

	r.Use(cors.Default())
	r.Run()
}
