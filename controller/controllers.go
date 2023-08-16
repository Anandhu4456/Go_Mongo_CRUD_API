package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"go-mongo/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://firstUser:mongo5656@cluster0.hrr1qic.mongodb.net/?retryWrites=true&w=majority"
const dbName = "movies"
const colName = "watchlist"

var collection *mongo.Collection

// connect to mongoDB

func init() {

	clientOptions := options.Client().ApplyURI(connectionString)

	// connect
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb connection success")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")
}

// Mongo helpers

// insert 1 record

func insertOneM(movie model.Movies) {
	insert, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("One movie data added to DB with Id : ", insert.InsertedID)
}

// update 1 movie

func updateOneM(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("updated count :", result.ModifiedCount)
}

// delete 1 record

func deleteOneM(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	dltCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got delete with delete count :", dltCount)
}

// delete all movies

func deleteAllM() int64 {
	dltResult, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies deleted with count: ", dltResult.DeletedCount)
	return dltResult.DeletedCount
}

// get all movies

func getAllM() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

// Actual controller file

func GetAllMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllM()
	json.NewEncoder(res).Encode(allMovies)
}

func CreateMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-www-form-urlencode")
	res.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Movies
	json.NewDecoder(req.Body).Decode(&movie)
	insertOneM(movie)
	json.NewEncoder(res).Encode(movie)
}

func MarkAsWatched(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-www-form-urlencode")
	res.Header().Set("Allow-Control-Allow-Method", "PUT")

	params := mux.Vars(req)
	updateOneM(params["id"])
	json.NewEncoder(res).Encode(params["id"])
}

func DltOneMovie(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type","application/x-www-form-urlencode")
	res.Header().Set("Allow-control-Allow-Method","DELETE")

	params:=mux.Vars(req)
	deleteOneM(params["id"])
	json.NewEncoder(res).Encode(params["id"])
}

func DltAllMovie(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type","application/x-www-form-urlencode")
	res.Header().Set("Allow-control-Allow-Method","DELETE")

	count:=deleteAllM()
	json.NewEncoder(res).Encode(count)
}