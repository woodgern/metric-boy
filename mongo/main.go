package Cassandra

import (
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func init() {
  Client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:27017"))
}