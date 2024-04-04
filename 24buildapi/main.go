package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for courses -file
type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

// Fake DB
var courses []Course

// Middlewre, helper (Use for add restruction) -file
func (c *Course) IsEmpty() bool {
	// Add restriction for stop submit empty data
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}


// Main function (entery point of program)
func main() {
	// welcome message
	fmt.Println("Welcome in my first API")
}


// controllers 

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome in API home page"))
}

// Controller for get all Courses
func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// Controller for get one course
func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get One course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	// Loop throught courses, find matching id and return the response
	for _, course := range courses{
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

// Controller for create a new course
func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create/POST One course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// if data like - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate unique id, string
	// append course into course

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

// Controller for update a course
func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("update course")
	w.Header().Set("Content-Type", "application/json")

	//first - grab id from req
	params := mux.Vars(r)

	//loop, id, remove, add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO: send a response when id id not found
}

// Controller for delete a course
func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	//first - grab id from req
	params := mux.Vars(r)

	// loop, id, remove(index, index+1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}