package main

import (
	"context"
	"log"
	"net/http"

	"github.com/MountainGator/playlist_CRUD/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	r := gin.Default()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	coll := client.Database("playlist_db").Collection("playlist")
	var result bson.M

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{
			"msg": "wazzup",
		})
	})


	r.POST("/login", controllers.UserLogin)
	r.POST("/create-playlist", controllers.NewPlaylist)
	r.GET("/find-playlist", controllers.FindPlaylist)

	r.Use(cors.Default())
	r.Run()
}
