package mongo

import (
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func init() {
  Client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:27017"))
}