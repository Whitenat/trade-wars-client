package main

import (
	// "fmt"
	"net/http"
	// "strconv"
)

//Home handler function which writes a byte slice
//containing "Welcome to the thunder dome"
func home(w http.ResponseWriter, r *http.Request) {
//Check if URL path matches exactly
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Welcome to the thunder dome"))
}

//Add a showNavigation handler function
func showNavigation(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// //Add a createSnippet handler function
// func createSnippet(w http.ResponseWriter, r *http.Request) {
// //use r.Method to check if request is POST
// 	if r.Method != http.MethodPost {
// //If not, send 405 error and let them know what is allowed
// 	w.Header().Set("Allow", http.MethodPost)
// 	w.WriteHeader(405)
// 	w.Write([]byte("Method Not Allowed"))
// 	return
// 	}

// 	w.Write([]byte("Create a new snippet..."))
// }