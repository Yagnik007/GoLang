package moviecontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	moviemodel "github.com/Yagnik007/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "YOUR_CONNECTION_STRING"

const dbName = "netflix"

const colName = "movielist"

// This is a reference to the mongodb collection
var collection *mongo.Collection

// connect with mongoDB
// We will write an init method, it runs only called once when this application is going to execute for initialization
func init() {
	//client options
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB - PLEASE REFER CONTEXT DOCS
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Success")

	collection = client.Database(dbName).Collection(colName)

	//collecton instance
	fmt.Println("Colletion instance is ready")
}

//MONGO helpers - file

//insert 1 record

func insertOneMovie(movie moviemodel.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie) //The Id of the inserted row is returned back

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a movie in db with id: ", inserted.InsertedID)
}

//update 1 record

// update 1 record
func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatalf("Error converting movieId to ObjectID: %v", err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatalf("Error updating movie: %v", err)
	}

	if res.ModifiedCount == 0 {
		log.Printf("No documents matched the filter: %v", filter)
	}

	fmt.Println("Modified count: ", res.ModifiedCount)
}

//Delete 1 record

func deleteOneMovie(movieId string) int64 {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with delete count", deleteCount.DeletedCount)

	return deleteCount.DeletedCount
}

//Delete all records from mongodb

func deleteAllMovies() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The number of movies deleted: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

//Get all movies from mongodb

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	defer cursor.Close(context.Background())
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
	return movies
}

//Actual Controllers

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie moviemodel.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-encode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
