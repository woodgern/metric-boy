package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/woodgern/metric-boy/Cassandra"
)

func main() {
  r := gin.New()
  CassandraSession := Cassandra.Session
  defer CassandraSession.Close()

  r.GET("/ping", Get_ping)
  r.Run(":8080")
}

func Get_ping(c *gin.Context) {
  c.String(http.StatusOK, "pong")
}
