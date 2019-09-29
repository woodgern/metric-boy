package main

import (
  "context"
  "encoding/json"
  "log"
  "net/http"

  "gopkg.in/mgo.v2/bson"
  "github.com/gin-gonic/gin"
  mongoclient "github.com/woodgern/metric-boy/mongo"
)

func main() {
  r := gin.New()
  if err := mongoclient.Client.Ping(context.TODO(), nil); err != nil {
    log.Println(err.Error())
  }

  collection := mongoclient.Client.Database("metrics").Collection("metrics")

  r.GET("/ping", Get_ping)
  r.POST("/metric", Post_metric)
  r.Run(":8080")
}

func Get_ping(c *gin.Context) {
  c.String(http.StatusOK, "pong")
}

func Post_metric(c *gin.Context) {
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

  collection.InsertOne(context.TODO(), &doc)

  c.JSON(http.StatusOK, requestbody)
}

func Get_metrics(c *gin.Context) {
  start := c.Query("start")
  end := c.Query("end")
}
