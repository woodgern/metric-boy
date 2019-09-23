package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/woodgern/metric-boy/mongo"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)


func main() {
  r := gin.New()
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  err = Client.Ping(ctx, readpref.Primary())

  client.Database("testing").Collection("numbers")

  r.GET("/ping", Get_ping)
  r.Run(":8080")
}

func Get_ping(c *gin.Context) {
  c.String(http.StatusOK, "pong")
}
