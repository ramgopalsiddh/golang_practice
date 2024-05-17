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

// Model for courses - file
type Course struct {
	CourseId   string  `json:"courseid"`
	CourseName string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author     *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB
var courses []Course

// Middlewre, helper (Use for add restruction) -file
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

// Main function (entery point of program)
func main() {
	// welcome message
	fmt.Println("Welcome in my first API")

	// Create new router
	r := mux.NewRouter()

	// Seeding
	rand.Seed(time.Now().UnixNano())
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Ram Gopal", Website: "ramgopal.dev"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Ram Gopal", Website: "ramgopal.dev"}})

	// Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")


	// Listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("<h1>Welcome in API home page</h1>")); err != nil {
		log.Println("Error writing response:", err)
	}
}

// controllers 

// Controller for get all Courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(courses); err != nil {
		log.Println("Error encoding JSON:", err)
	}
}

// Controller for get one course
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	// Loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			if err := json.NewEncoder(w).Encode(course); err != nil {
				log.Println("Error encoding JSON:", err)
			}
			return
		}
	}
	if err := json.NewEncoder(w).Encode("No Course found with given id"); err != nil {
		log.Println("Error encoding JSON:", err)
	}
}

// Controller for create a new course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create/POST One course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		if _, err := w.Write([]byte("Please send some data")); err != nil {
			log.Println("Error writing response:", err)
		}
		return
	}

	// if data like - {}
	var course Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		log.Println("Error decoding JSON:", err)
		if _, err := w.Write([]byte("No data inside JSON")); err != nil {
			log.Println("Error writing response:", err)
		}
		return
	}

	//TODO: check only if title is duplicate
	// loop, title matches with course.coursename, JSON

	// generate unique id, string
	// append course into courses

	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	if err := json.NewEncoder(w).Encode(course); err != nil {
		log.Println("Error encoding JSON:", err)
	}
}

// Controller for update a course
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update course")
	w.Header().Set("Content-Type", "application/json")

	//first - grab id from req
	params := mux.Vars(r)

	//loop, id, remove, add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
				log.Println("Error decoding JSON:", err)
				return
			}
			course.CourseId = params["id"]
			courses = append(courses, course)
			if err := json.NewEncoder(w).Encode(course); err != nil {
				log.Println("Error encoding JSON:", err)
			}
			return
		}
	}
	// TODO: send a response when id is not found
}

// Controller for delete a course
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	//first - grab id from req
	params := mux.Vars(r)

	// loop, id, remove(index, index+1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			//TODO: send confirm or deny response
			break
		}
	}
}
