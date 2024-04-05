package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ramgopalsiddh/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://ramgopalsiddh:<password>@cluster0.6gq3p7r.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"


// most important part

var collection *mongo.Collection

// connect with mongoDB

func init(){
	// create clientOption/ connection
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	// context.TODO this use when you unclear which context to use [ https://pkg.go.dev/context#TODO ]
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connected successful")

	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("Collection instance/reference is ready")
}


// MongoDB helpers 

// functon for update 1 record
func updateOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	// inside MongoDB has no Json it's bson these are same but bson give more things
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	// find out how many record updated
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count: ", result.ModifiedCount)
}


// function for Delete one record
func deleteOneMovie(movieId string){
	id, _ :=  primitive.ObjectIDFromHex(movieId)
	// filter from db
	filter := bson.M{"_id": id}
	// perform operation
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	// message to show delete successfull
	fmt.Println("Movie got delete with delete count: ", deleteCount)
}


// function for Delete all records from MongoDB
func deleteAllMovie() int64{
	// pass direct bson as variable because this is not use tomuch
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil) // empty parathese means delete all parameter

	if err != nil {
		log.Fatal(err)
	}

	// print count of total deleted record
	fmt.Println("Number of movies delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}


// function for get all movies from database
func getAllMovies() []primitive.M {
	curser, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for curser.Next(context.Background()){
		var movie bson.M
		err := curser.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer curser.Close(context.Background())
	return movies
}


// Actual controller 

// Controller for get all movies
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urleancode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

// Create a Movie
func CreateMovie(w http.ResponseWriter, r *http.Request){
	// Header 
	w.Header().Set("Content-Type", "application/x-www-form-urleancode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	// add movie
	var movie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	fmt.Printf("movie: %v \n", movie)
	if err != nil {
		log.Fatal(err)
	}

	inserted, err := collection.InsertOne(context.Background(), movie)
	// return json after operation
	json.NewEncoder(w).Encode(movie)
	// error 
	if err != nil {
		log.Fatal(err)
	}
	// Success message
	fmt.Println("Inserted one movie in db with id:", inserted.InsertedID)
}


// Mark movie as watched
func MarkAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urleancode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	// Use id & mux and update data
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}


// Delete a movie 
func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urleancode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	// Use id & mux and delete movie
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// Delete all movies 
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urleancode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}