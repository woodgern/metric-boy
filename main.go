package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
	mongoclient "github.com/woodgern/metric-boy/mongo"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	r := gin.New()
	if err := mongoclient.Client.Ping(context.TODO(), nil); err != nil {
		log.Println(err.Error())
	}

	r.GET("/ping", GetPing)
	r.POST("/metric", PostMetric)
	r.POST("/query", GetMetrics)
	r.Run(":8080")
}

func GetPing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func PostMetric(c *gin.Context) {
	var requestbody interface{}
	if err := c.ShouldBindJSON(&requestbody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jsonString, err := json.Marshal(requestbody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var doc interface{}
	if err := bson.UnmarshalJSON([]byte(jsonString), &doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := getCollection()
	collection.InsertOne(context.TODO(), &doc)

	c.JSON(http.StatusOK, requestbody)
}

func GetMetrics(c *gin.Context) {
	return
}

func getCollection() *mongo.Collection {
	return mongoclient.Client.Database("metrics").Collection("metrics")
}
