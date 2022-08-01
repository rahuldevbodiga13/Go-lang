package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rahuldevbodiga13/mongoapi/helpers"
	"github.com/rahuldevbodiga13/mongoapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionUrl = "mongodb+srv://rahuldev27:Rahana%4027@cluster0.duaoe.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "netflix"
const collectionName = "watchlist"

//MOST IMP - reference of mongoDb collection
var collection *mongo.Collection

//connect with MongoDB

func init() {
	//client options
	clientOption := options.Client().ApplyURI(connectionUrl)

	//connect to mongoDb
	client, err := mongo.Connect(context.TODO(), clientOption)
	helpers.CheckForError(err)

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(collectionName)

	//collection instance/reference
	fmt.Println("Collection reference is ready")

}

//MongoDB helpers - go into separate file

//insert one record

func insertOneMovie(movie models.Netflix) {

	inserted, err := collection.InsertOne(context.Background(), movie)
	helpers.CheckForError(err)

	fmt.Println("Inserted one movie in DB with Id : ", inserted.InsertedID)

}

// update one record

func updateOneMovie(movieId string) {
	objId, err := primitive.ObjectIDFromHex(movieId)
	helpers.CheckForError(err)

	filter := bson.M{"_id": objId}
	update := bson.M{"$set": bson.M{"watched": true}}
	res, err := collection.UpdateOne(context.Background(), filter, update)
	helpers.CheckForError(err)

	fmt.Println("Updated count is : ", res.ModifiedCount)
}

// delete one movie/record

func deleteOneMovie(movieId string) {
	objId, err := primitive.ObjectIDFromHex(movieId)
	helpers.CheckForError(err)

	filter := bson.M{"_id": objId}

	res, err := collection.DeleteOne(context.Background(), filter)
	helpers.CheckForError(err)
	fmt.Println("Deleted record count is : ", res.DeletedCount)
}

// delete multiple records

func deleteManyMovies() {
	res, err := collection.DeleteMany(context.Background(), bson.D{{}})
	helpers.CheckForError(err)
	fmt.Println("Deleted count is :", res.DeletedCount, nil)
}

// get all movies from mongoDB

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{})
	helpers.CheckForError(err)

	defer cursor.Close(context.Background())

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		helpers.CheckForError(err)
		movies = append(movies, movie)

	}
	return movies
}

// Actual controllers - current file ( DB helpers go into seperate file )

func GetAllMoviesInJson(w http.ResponseWriter, r *http.Request) {
	movies := getAllMovies()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	helpers.CheckForError(err)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode("Update Successful")
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("Deletion successfull")
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	deleteManyMovies()
	json.NewEncoder(w).Encode("All movies deleted")
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Experience our fastest netflix Apis</h1>"))
}
