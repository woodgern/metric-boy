package Cassandra

import (
  "github.com/gocql/gocql"
  "fmt"
)

var Session *gocql.Session

func init() {
  var err error

  cluster := gocql.NewCluster("http://cassandra")
  cluster.Keyspace = "metric-boy"
  Session, err = cluster.CreateSession()
  if err != nil {
    panic(err)
  }
  fmt.Println("cassandra init done")
}