package controller

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


const connectionString = "mongodb://localhost:27017"



// db name: netflix
var DbName = "netflix"
// collection name: movies

var collection *mongo.Collection

// connect to the database
func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI(
		connectionString)

	// connect to the database
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	// check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}

	collection = client.Database(DbName).Collection("movies")

	fmt.Println("Connected to MongoDB!")
}


