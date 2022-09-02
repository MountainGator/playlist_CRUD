package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

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
}
