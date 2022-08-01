package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	Greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", ServeHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))

}

func Greeter() {
	fmt.Println("Hello go mod users ")
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series!</h1>"))

}
