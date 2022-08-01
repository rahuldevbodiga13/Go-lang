package router

import (
	"github.com/gorilla/mux"
	"github.com/rahuldevbodiga13/mongoapi/controllers"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.ServeHome).Methods("GET")
	router.HandleFunc("/api/movies", controllers.GetAllMoviesInJson).Methods("GET")
	router.HandleFunc("/api/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controllers.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/deletemovies", controllers.DeleteAllMovies).Methods("DELETE")
	return router
}
