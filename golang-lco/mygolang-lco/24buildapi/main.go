package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Model for courses - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB
var courses []Course

// middleware, helpers - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""}
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to building apis in golang")
	router := mux.NewRouter()

	//Seeding the data

	courses = append(courses, Course{CourseId: "27", CourseName: "World", CoursePrice: 2727, Author: &Author{FullName: "Rahana Dev", Website: "mr.com"}})
	courses = append(courses, Course{CourseId: "13", CourseName: "OURS", CoursePrice: 272, Author: &Author{FullName: "Rahana Dev", Website: "mr.com"}})

	router.HandleFunc("/", ServeHome).Methods("GET")
	router.HandleFunc("/courses", GetAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", GetOneCourse).Methods("GET")
	router.HandleFunc("/course", CreateOneCourse).Methods("POST")
	router.HandleFunc("/course/{id}", UpdateOneCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", DeleteOneCourse).Methods("DELETE")

	//listening to a port
	log.Fatal(http.ListenAndServe(":4000", router))

}

// Controllers go into another file

//Serve home route

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to our fastest api service</h1>"))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through courses and find matching id , return response

	for _, value := range courses {
		if value.CourseId == params["id"] {
			json.NewEncoder(w).Encode(value)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given id " + params["id"])
	return

}

func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating one course...")
	w.Header().Set("Content-Type", "application/json")

	// what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//generate unique id -> convert to string
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func UpdateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course...")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	// loop through, find id, remove the object,  and add it with ID (update it)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var c Course
			_ = json.NewDecoder(r.Body).Decode(&c)
			c.CourseId = params["id"]
			courses = append(courses, c)
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	// Cases like if course with required id is not found should also be handled
}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Deletion successful")
			return
		}
	}
}
