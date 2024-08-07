package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// Model for our API
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"-"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware - method for course struct
func (c *Course) Isempty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API")
	r := mux.NewRouter()

	//Seeding Database

	courses = append(courses, Course{CourseId: "2", CourseName: "Reactjs", CoursePrice: 299, Author: &Author{Fullname: "Yagnik Talaviya", Website: "yagnik.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN", CoursePrice: 500, Author: &Author{Fullname: "Yagnik Talaviya", Website: "go.dev"}})

	//Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//Listen
	log.Fatal(http.ListenAndServe(":5000", r))

}

//controllers

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Courses</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All the courses")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found with id: " + params["id"])
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode((&course))
	if course.Isempty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	for _, item := range courses {
		if course.CourseName == item.CourseName {
			json.NewEncoder(w).Encode("Course with the same name exists")
			return
		}
	}

	//generate new id -> to string -> append course into courses

	course.CourseId = strconv.Itoa(rand.Intn(1000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for ind, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:ind], courses[ind+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One course")
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for ind, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:ind], courses[ind+1:]...)
			json.NewEncoder(w).Encode("Course with id: " + params["id"] + " deleted!!")
			return
		}
	}
	json.NewEncoder(w).Encode("Course with id " + params["id"] + " Not found!!")
	return
}
