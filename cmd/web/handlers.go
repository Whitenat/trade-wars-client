package main

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
	// "strconv"
)

//Home handler function which serves the home page
func home(w http.ResponseWriter, r *http.Request) {
	//Check if URL path matches exactly
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
	}
	//If it's a POST request, display callsign
	if r.Method == http.MethodPost {

	}
	//If it's a GET request, display "Enter username"
	if r.Method == http.MethodGet {

	}
	//Otherwise, throw a 405 error
	if r.Method != http.MethodPost && r.method != http.MethodGet {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
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

//Navigation handler function which serves the navigation page
func showNavigation(w http.ResponseWriter, r *http.Request) {
	//Check if URL path matches exactly
	if r.URL.Path != "/navigation" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/navigation.page.tmpl",
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

//Navigation handler function which serves the trade page
func showNavigation(w http.ResponseWriter, r *http.Request) {
	//Check if URL path matches exactly
	if r.URL.Path != "/trade" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/trade.page.tmpl",
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
