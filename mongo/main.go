package mongo

import (
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func init() {
  var err error
  if clientOptions, err := options.Client().ApplyURI("mongodb://mongo:27017"); err != nil {
    log.Fatal(err)
  }
  Client, err = mongo.Connect(context.TODO(), clientOptions)
}
