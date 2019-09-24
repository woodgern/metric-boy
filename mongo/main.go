package mongo

import (
  "context"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func init() {
  clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
  Client, err = mongo.Connect(context.TODO(), clientOptions)
}
