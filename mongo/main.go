package mongo

import (
  "context"
  "log"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func init() {
  var err error
  clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
  if Client, err = mongo.Connect(context.TODO(), clientOptions); err != nil {
    log.Fatal(err)
  }
}
