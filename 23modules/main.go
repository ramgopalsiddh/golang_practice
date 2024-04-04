package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Welcome message
	fmt.Println("Welcome in module section")

	// Mux syntex for serve data
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	// Create server in go lang
	log.Fatal(http.ListenAndServe(":4000", r))
}

// This is serve data in web
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome in Mode test</h1>"))

}