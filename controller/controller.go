package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/samdtech/go-mongo-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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


// insert one document
func insertOne(movie model.Netflix) {
	insertResult, err := collection.InsertOne(context.TODO(), movie)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

// update one movie
func UpdateOne(movieId string ) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	// create filter
	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)

}

func deleteOne(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	deleteCount, err :=	collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie got deleted: ", deleteCount)


}

func deleteAll(){

	resultCount, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

		if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie got deleted: ", resultCount.DeletedCount)


}


func getAllMovies() []primitive.M {
	cursor, err :=	collection.Find(context.Background(), bson.M{})

	if err != nil{
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()){
		var movie bson.M

		err := cursor.Decode(&movie)

		if err !=nil{
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())

	return movies
}


// actual controllers

// get all movies
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movies := getAllMovies()

	fmt.Println("all movies: ", movies)

	json.NewEncoder(w).Encode(movies)
}

// create movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOne(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	// get the params
	params := mux.Vars(r)

	 UpdateOne(params["id"])

	 json.NewEncoder(w).Encode("Updated Successfully")
}

func DeleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	// get the params
	params := mux.Vars(r)

	deleteOne(params["id"])

	 json.NewEncoder(w).Encode("Deleted Successfully")
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")


	deleteAll()

	 json.NewEncoder(w).Encode("Deleted Successfully")
}