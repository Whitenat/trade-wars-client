package main

import (
	// "fmt"
	"html/template"
	"log"
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

files := []string{
	"./ui/html/home.page.tmpl",
	}

// Use the template.ParseFiles() function to read the template file into a
// template set. If there's an error, we log the detailed error message and use
// the http.Error() function to send a generic 500 Internal Server Error
// response to the user.
    ts, err := template.ParseFiles(files...)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
 	}

// We then use the Execute() method on the template set to write the template
// content as the response body. The last parameter to Execute() represents any
// dynamic data that we want to pass in, which for now we'll leave as nil.
    err = ts.Execute(w, nil)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
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
