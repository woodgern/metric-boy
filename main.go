package main

import (
  "net/http"
  "time"
  "gopkg.in/mgo.v2/bson"
  "log"

  "github.com/gin-gonic/gin"
  mongoclient "github.com/woodgern/metric-boy/mongo"
)

type Metric struct {
  Event string `json:"event" binding:"required"`
  Body  string `json:"body" binding:"required"`
}


func main() {
  r := gin.New()
  if err := mongoclient.Client.Ping(ctx, nil); err != nil {
    log.Println(err.Error())
  }

  collection := mongoclient.Client.Database("metrics").Collection("metrics")

  r.GET("/ping", Get_ping)
  r.POST("/test", func (c *gin.Context) {
    var requestbody Metric
    if err := c.ShouldBindJSON(&requestbody); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    var bdoc interface{}
    if err := bson.UnmarshalJSON([]byte(requestbody.Body), &bdoc); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    log.Println(&bdoc)
    collection.InsertOne(ctx, &bdoc)

    c.JSON(http.StatusOK, requestbody)
  })
  r.Run(":8080")
}

func Get_ping(c *gin.Context) {
  c.String(http.StatusOK, "pong")
}
